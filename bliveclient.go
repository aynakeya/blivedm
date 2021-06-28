package blivedm

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type BLiveWsClient struct {
	RoomId    int
	ShortId   int
	Uid       int
	RoomInfo  RoomInfoData
	DanmuInfo DanmuInfoData

	HearbeatInterval time.Duration

	WsConn  *websocket.Conn
	Running bool

	Chnl chan int

	Handlers map[string][]HandlerFunc
}

func New_BLiveWsClient(uid int) BLiveWsClient {
	return BLiveWsClient{ShortId: uid, HearbeatInterval: 25 * time.Second}
}

func (self *BLiveWsClient) GetRoomInfo() bool {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("room_id", cast.ToString(self.ShortId)).
		Get(ROOM_INFO_API)
	if err != nil {
		self.RoomId = self.ShortId
		return false
	}
	var sResp RoomInfoResponse
	if err := json.Unmarshal(resp.Body(), &sResp); err != nil {
		self.RoomId = self.ShortId
		return false
	}
	// 虽然但是，如果用的是shortid,返回的code不是0，但是data是正确的，所以我也不明白了
	//if sResp.Code != 0 {
	//	self.RoomId = self.ShortId
	//	return false
	//}
	if sResp.Data.RoomID == 0 {
		self.RoomId = self.ShortId
		return false
	}
	self.RoomInfo = sResp.Data
	self.RoomId = sResp.Data.RoomID
	return true
}

func (self *BLiveWsClient) GetDanmuInfo() bool {
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("id", cast.ToString(self.RoomId)).
		SetQueryParam("type", "0").
		Get(DANMAKU_INFO_API)
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
	self.DanmuInfo = sResp.Data
	return true
}

func (self *BLiveWsClient) InitRoom() bool {
	return self.GetRoomInfo() && self.GetDanmuInfo()
}

func (self *BLiveWsClient) Start() {
	if !self.Running {
		self.ConnectDanmuServer()
	}
	self.Chnl = make(chan int)
	_ = <-self.Chnl
}

func (self *BLiveWsClient) Stop() {
	self.Running = false
	self.Chnl <- 1
}

func (self *BLiveWsClient) ConnectDanmuServer() bool {
	danmuServer := self.DanmuInfo.HostList[0]
	danmuHost := fmt.Sprintf("%s:%d", danmuServer.Host, danmuServer.WssPort)
	uri := url.URL{Scheme: "wss", Host: danmuHost, Path: "/sub"}
	conn, _, err := websocket.DefaultDialer.Dial(uri.String(), nil)
	if err != nil {
		return false
	}
	self.WsConn = conn
	self.sendAuth()

	self.Running = true

	go func() {
		for self.Running {
			self.ReadMessage()
		}
	}()

	go func() {
		for self.Running {
			self.sendHeartBeat()
			time.Sleep(self.HearbeatInterval)
		}
	}()
	return true
}

func (self *BLiveWsClient) Disconnect() {
	self.Running = false
	err := self.WsConn.Close()
	if err != nil {
		return
	}
}

func (self *BLiveWsClient) ReadMessage() {
	messageType, message, err := self.WsConn.ReadMessage()
	if err != nil {
		return
	}
	if messageType != websocket.BinaryMessage {
		return
	}
	header, data, ok := ResolveWSPacket(message)
	if !ok {
		return
	}
	if header.ProtocolVersion == 2 {
		if datas, err := ZlibDeCompress(data); err == nil {
			offset := 0
			for offset < len(datas) {
				h, d, k := ResolveWSPacket(datas[offset:])
				if !k {
					break
				}
				self.handleMessage(h, d)
				offset += int(h.PacketLength)
			}
		}
	} else {
		self.handleMessage(header, data)
	}
}

func (self *BLiveWsClient) handleMessage(header WsHeader, data []byte) {
	switch int32(header.Operation) {
	case OpSendMsg, OpSendMsgReply:
		self.handleCommand(data)
	default:
		return
	}
}

func (self *BLiveWsClient) handleCommand(data []byte) {
	jData := gjson.Get(string(data), "@this")
	if cmd := jData.Get("cmd"); cmd.Exists() {
		self.CallHandler(cmd.String(), &Context{
			Cmd:      cmd.String(),
			RawData:  string(data),
			JsonData: jData,
		})
	} else {
		return
	}
}

func (self *BLiveWsClient) RegHandler(cmd string, handlerFunc HandlerFunc) {
	if self.Handlers == nil {
		self.Handlers = map[string][]HandlerFunc{}
	}
	//if _, ok := self.Handlers[cmd]; !ok {
	//	self.Handlers[cmd] = make([]HandlerFunc, 0)
	//}
	self.Handlers[cmd] = append(self.Handlers[cmd], handlerFunc)
}

func (self *BLiveWsClient) CallHandler(cmd string, context *Context) {
	if handlers, ok := self.Handlers[cmd]; ok {
		for _, handler := range handlers {
			handler(context)
		}
	}
}

func (self *BLiveWsClient) sendHeartBeat() {
	err := self.WsConn.WriteMessage(websocket.BinaryMessage, MakeWSPacket(OpHeartbeat, []byte{}))
	if err != nil {
		return
	}
}

func (self *BLiveWsClient) sendAuth() {
	data := map[string]interface{}{
		"uid":    self.Uid,
		"roomid": self.RoomId,
		// protover  = 3 unknown encryption
		"protover": 2,
		"platform": "web",
		"type":     2,
		"key":      self.DanmuInfo.Token,
	}
	dataBytes, _ := json.Marshal(data)
	err := self.WsConn.WriteMessage(websocket.BinaryMessage, MakeWSPacket(OpAuth, dataBytes))
	if err != nil {
		return
	}
}
