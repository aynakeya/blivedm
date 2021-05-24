package blivedm

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type Context struct {
	Cmd      string
	RawData  string
	JsonData gjson.Result
}

func (self *Context) ToGiftMessage() (GiftMessage, bool) {
	if self.Cmd != CmdSendGift {
		return GiftMessage{}, false
	}
	var gm GiftMessage
	if err := json.Unmarshal([]byte(self.JsonData.Get("data").Raw), &gm); err != nil {
		return GiftMessage{}, false
	}
	return gm, true
}

func (self *Context) ToDanmakuMessage() (DanmakuMessage, bool) {
	if self.Cmd != CmdDanmaku {
		return DanmakuMessage{}, false
	}
	return DanmakuMessage{
		Mode:           self.JsonData.Get("info.0.1").Int(),
		FontSize:       self.JsonData.Get("info.0.2").Int(),
		Color:          self.JsonData.Get("info.0.3").Int(),
		Timestamp:      self.JsonData.Get("info.0.4").Int(),
		Rnd:            self.JsonData.Get("info.0.5").Int(),
		UID_CRC32:      self.JsonData.Get("info.0.7").String(),
		MsgType:        self.JsonData.Get("info.0.9").Int(),
		Bubble:         self.JsonData.Get("info.0.10").Int(),
		Msg:            self.JsonData.Get("info.1").String(),
		Uid:            self.JsonData.Get("info.2.0").Int(),
		Uname:          self.JsonData.Get("info.2.1").String(),
		Admin:          self.JsonData.Get("info.2.2").Bool(),
		Vip:            self.JsonData.Get("info.2.3").Bool(),
		Svip:           self.JsonData.Get("info.2.4").Bool(),
		Urank:          self.JsonData.Get("info.2.5").Int(),
		MobileVerify:   self.JsonData.Get("info.2.6").Bool(),
		UnameColor:     self.JsonData.Get("info.2.7").String(),
		MedalLevel:     self.JsonData.Get("info.3.0").Int(),
		MedalName:      self.JsonData.Get("info.3.1").String(),
		MedalUpName:    self.JsonData.Get("info.3.2").String(),
		MedalRoomId:    self.JsonData.Get("info.3.3").Int(),
		MedalColor:     self.JsonData.Get("info.3.4").Int(),
		SpecialMedal:   self.JsonData.Get("info.3.5").String(),
		UserLevel:      self.JsonData.Get("info.4.0").Int(),
		UserLevelColor: self.JsonData.Get("info.4.2").Int(),
		UserLevelRank:  self.JsonData.Get("info.4.3").String(),
		OldTitle:       self.JsonData.Get("info.5.0").String(),
		Title:          self.JsonData.Get("info.5.1").String(),
		PrivilegeType:  self.JsonData.Get("info.5").Int(),
	}, true
}
