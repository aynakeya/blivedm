package blivedm

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type BLiveWsClient struct {
	RoomId    int
	ShortId   int
	RoomInfo  RoomInfoData
	DanmuInfo DanmuInfoData

	Account DanmuAccount

	HearbeatInterval time.Duration

	WsConn  *websocket.Conn
	Running bool

	Chnl chan int

	Handlers map[string][]HandlerFunc

	OnDisconnect DisconnectCallback
}

func New_BLiveWsClient(shorId int) BLiveWsClient {
	return BLiveWsClient{ShortId: shorId, Account: DanmuAccount{UID: 0}, HearbeatInterval: 25 * time.Second}
}

func (c *BLiveWsClient) GetRoomInfo() bool {
	resp, err := ApiGetRoomInfo(c.Account, c.ShortId)
	if err != nil {
		c.RoomId = c.ShortId
		return false
	}
	var sResp RoomInfoResponse
	if err := json.Unmarshal(resp.Body(), &sResp); err != nil {
		c.RoomId = c.ShortId
		return false
	}
	if sResp.Code != 0 {
		c.RoomId = c.ShortId
		return false
	}
	if sResp.Data.RoomId == 0 {
		c.RoomId = c.ShortId
		return true
	}
	c.RoomInfo = sResp.Data
	c.RoomId = sResp.Data.RoomId
	return true
}

func (c *BLiveWsClient) GetDanmuInfo() bool {
	resp, err := ApiGetDanmuInfo(c.Account, c.RoomId, 0)
	if err != nil {
		return false
	}
	var sResp DanmuInfoResponse
	if err := json.Unmarshal(resp.Body(), &sResp); err != nil {
		return false
	}
	if sResp.Code != 0 {
		return false
	}
	c.DanmuInfo = sResp.Data
	return true
}

func (c *BLiveWsClient) InitRoom() bool {
	return c.GetRoomInfo() && c.GetDanmuInfo()
}

func (c *BLiveWsClient) Start() {
	if !c.Running {
		c.GetRoomInfo()
		c.GetDanmuInfo()
		c.ConnectDanmuServer()
	}
	c.Chnl = make(chan int)
	defer func() { _ = <-c.Chnl }()
}

func (c *BLiveWsClient) Stop() {
	c.Running = false
	c.Chnl <- 1
}

func (c *BLiveWsClient) ConnectDanmuServer() bool {
	if c.DanmuInfo.HostList == nil {
		return false
	}
	danmuServer := c.DanmuInfo.HostList[0]
	danmuHost := fmt.Sprintf("%s:%d", danmuServer.Host, danmuServer.WssPort)
	uri := url.URL{Scheme: "wss", Host: danmuHost, Path: "/sub"}
	conn, _, err := websocket.DefaultDialer.Dial(uri.String(), nil)
	if err != nil {
		return false
	}
	c.WsConn = conn
	c.Running = c.sendAuth()

	go func() {
		for c.Running {
			r := c.ReadMessage()
			if !r && c.OnDisconnect != nil {
				c.OnDisconnect(c)
			}
			c.Running = r
		}
	}()

	go func() {
		for c.Running {
			c.sendHeartBeat()
			time.Sleep(c.HearbeatInterval)
		}
	}()
	return c.Running
}

func (c *BLiveWsClient) Disconnect() {
	c.Running = false
	err := c.WsConn.Close()
	if err != nil {
		return
	}
}

func (c *BLiveWsClient) ReadMessage() bool {
	messageType, message, err := c.WsConn.ReadMessage()
	if err != nil {
		return false
	}
	if messageType != websocket.BinaryMessage {
		return false
	}
	header, data, ok := ResolveWSPacket(message)
	if !ok {
		return false
	}
	if header.ProtocolVersion == 2 {
		if datas, err := ZlibDeCompress(data); err == nil {
			offset := 0
			for offset < len(datas) {
				h, d, k := ResolveWSPacket(datas[offset:])
				if !k {
					break
				}
				c.handleMessage(h, d)
				offset += int(h.PacketLength)
			}
		}
	} else {
		c.handleMessage(header, data)
	}
	return true
}

func (c *BLiveWsClient) handleMessage(header WsHeader, data []byte) {
	switch int(header.Operation) {
	case OpSendMsg, OpSendMsgReply:
		c.handleCommand(data)
	default:
		return
	}
}

func (c *BLiveWsClient) handleCommand(data []byte) {
	jData := gjson.Get(string(data), "@this")
	if cmd := jData.Get("cmd"); cmd.Exists() {
		// fix DANMU_MSG:4:0:2:2:2:0
		cmds := strings.Split(cmd.String(), ":")
		c.CallHandler(cmds[0], &Context{
			Cmd:      cmds[0],
			RawData:  string(data),
			JsonData: jData,
		})
	} else {
		return
	}
}

func (c *BLiveWsClient) RegHandler(cmd string, handlerFunc HandlerFunc) {
	if c.Handlers == nil {
		c.Handlers = map[string][]HandlerFunc{}
	}
	//if _, ok := c.Handlers[cmd]; !ok {
	//	c.Handlers[cmd] = make([]HandlerFunc, 0)
	//}
	c.Handlers[cmd] = append(c.Handlers[cmd], handlerFunc)
}

func (c *BLiveWsClient) CallHandler(cmd string, context *Context) {
	if handlers, ok := c.Handlers[cmd]; ok {
		for _, handler := range handlers {
			handler(context)
		}
	}
}

func (c *BLiveWsClient) sendHeartBeat() error {
	return c.WsConn.WriteMessage(websocket.BinaryMessage, MakeWSPacket(OpHeartbeat, []byte{}))
}

func (c *BLiveWsClient) sendAuth() bool {
	data := map[string]interface{}{
		"uid":    c.Account.UID,
		"roomid": c.RoomId,
		// protover  = 3 unknown encryption
		"protover": 2,
		"platform": "web",
		"type":     2,
		"key":      c.DanmuInfo.Token,
	}
	dataBytes, _ := json.Marshal(data)
	err := c.WsConn.WriteMessage(websocket.BinaryMessage, MakeWSPacket(OpAuth, dataBytes))
	if err != nil {
		return false
	}
	return true
}

func (c *BLiveWsClient) SendMessage(msg DanmakuSendForm) (DanmakuSendResponse, error) {
	msg.CSRF = c.Account.BilibiliJCT
	msg.RoomId = c.RoomId
	resp, err := ApiSendDanmu(c.Account, msg)
	if err != nil {
		return DanmakuSendResponse{BaseV1Response{Code: -1}}, err
	}
	var sendResp DanmakuSendResponse
	if err := json.Unmarshal(resp.Body(), &sendResp); err != nil {
		return DanmakuSendResponse{BaseV1Response{Code: -1}}, err
	}
	return sendResp, err
}
