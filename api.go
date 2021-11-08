package blivedm

import (
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"net/http"
)

const (
	ROOM_INFO_API         = "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo"
	DANMAKU_INFO_API      = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo"
	SEND_MSG_API          = "https://api.live.bilibili.com/msg/send"
	ROOM_DANMU_CONFIG_API = "https://api.live.bilibili.com/xlive/web-room/v1/dM/GetDMConfigByGroup"
)

func ApiGetRoomInfo(shortId int) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetQueryParam("room_id", cast.ToString(shortId)).
		Get(ROOM_INFO_API)
}

func ApiGetDanmuInfo(roomId int, infoType int) (*resty.Response, error) {
	client := resty.New()
	return client.R().
		SetQueryParam("id", cast.ToString(roomId)).
		SetQueryParam("type", cast.ToString(infoType)).
		Get(DANMAKU_INFO_API)
}

func ApiGetRoomDanmuConfig(roomId int) (*resty.Response, error) {
	client := resty.New()
	return client.R().
		SetQueryParam("room_id", cast.ToString(roomId)).
		Get(DANMAKU_INFO_API)
}

func ApiSendDanmu(sender DanmuAccount, msg DanmakuSendForm) (*resty.Response, error) {
	client := resty.New()
	return client.R().
		SetCookies([]*http.Cookie{
			{
				Name:  "SESSDATA",
				Value: sender.SessionData,
			},
			{
				Name:  "bili_jct",
				Value: sender.BilibiliJCT,
			},
		}).
		SetFormData(map[string]string{
			"bubble":     cast.ToString(msg.Bubble),
			"msg":        msg.Message,
			"color":      msg.Color,
			"mode":       cast.ToString(msg.mode),
			"fontsize":   cast.ToString(msg.Fontsize),
			"rnd":        cast.ToString(msg.Rnd),
			"roomid":     cast.ToString(msg.RoomId),
			"csrf":       cast.ToString(msg.CSRF),
			"csrf_token": cast.ToString(msg.CSRF),
		}).Post(SEND_MSG_API)
}
