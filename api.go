package blivedm

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
	"net/http"
)

const (
	ROOM_INFO_API         = "https://api.live.bilibili.com/room/v1/Room/get_info"
	ROOM_INFO_API_V2      = "https://api.live.bilibili.com/xlive/web-room/v2/index/getRoomPlayInfo"
	DANMAKU_INFO_API      = "https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo"
	SEND_MSG_API          = "https://api.live.bilibili.com/msg/send"
	SET_MSG_CONFIG_API    = "https://api.live.bilibili.com/xlive/web-room/v1/dM/AjaxSetConfig"
	ROOM_DANMU_CONFIG_API = "https://api.live.bilibili.com/xlive/web-room/v1/dM/GetDMConfigByGroup"
	UP_INFO_API           = "https://api.bilibili.com/x/space/acc/info"
)

func ApiGetRoomInfo(sender DanmuAccount, shortId int) (*resty.Response, error) {
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
		SetQueryParam("id", cast.ToString(shortId)).
		Get(ROOM_INFO_API)
}

func ApiGetRoomInfoV2(sender DanmuAccount, shortId int) (*resty.Response, error) {
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
		SetQueryParam("room_id", cast.ToString(shortId)).
		Get(ROOM_INFO_API)
}

func ApiGetDanmuInfo(sender DanmuAccount, roomId int, infoType int) (*resty.Response, error) {
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
		SetQueryParam("id", cast.ToString(roomId)).
		SetQueryParam("type", cast.ToString(infoType)).
		Get(DANMAKU_INFO_API)
}

func ApiGetRoomDanmuConfig(sender DanmuAccount, roomId int) (*DanmuConfigResponse, error) {
	resp, err := resty.New().R().
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
		SetQueryParam("room_id", cast.ToString(roomId)).
		Get(ROOM_DANMU_CONFIG_API)
	if err != nil {
		return nil, err
	}
	var dmcfgResp DanmuConfigResponse
	if err := json.Unmarshal(resp.Body(), &dmcfgResp); err != nil {
		return nil, err
	}
	return &dmcfgResp, nil
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
			"mode":       cast.ToString(msg.Mode),
			"fontsize":   cast.ToString(msg.Fontsize),
			"rnd":        cast.ToString(msg.Rnd),
			"roomid":     cast.ToString(msg.RoomId),
			"csrf":       cast.ToString(msg.CSRF),
			"csrf_token": cast.ToString(msg.CSRF),
		}).Post(SEND_MSG_API)
}

func ApiSetDanmuConfig(sender DanmuAccount, roomId int, config string, value string) (*SetDanmuConfigResponse, error) {
	resp, err := resty.New().R().
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
			"room_id":    cast.ToString(roomId),
			config:       value,
			"csrf":       cast.ToString(sender.BilibiliJCT),
			"csrf_token": cast.ToString(sender.BilibiliJCT),
			"visit_id":   "",
		}).Post(SET_MSG_CONFIG_API)
	if err != nil {
		return nil, err
	}
	var tmpResp SetDanmuConfigResponse
	if err := json.Unmarshal(resp.Body(), &tmpResp); err != nil {
		return nil, err
	}
	return &tmpResp, nil
}

func ApiGetUpInfo(mid int) (*UpInfoResponse, error) {
	resp, err := resty.New().R().
		SetQueryParam("mid", cast.ToString(mid)).
		SetQueryParam("jsonp", "jsonp").
		Get(UP_INFO_API)
	if err != nil {
		return nil, err
	}
	var upinfoResp UpInfoResponse
	if err := json.Unmarshal(resp.Body(), &upinfoResp); err != nil {
		return nil, err
	}
	return &upinfoResp, nil
}
