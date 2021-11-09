package main

import (
	"fmt"
	"github.com/aynakeya/blivedm"
	"time"
)

func main() {
	cl := blivedm.BLiveWsClient{ShortId: 23028000, Account: blivedm.DanmuAccount{UID: 0}, HearbeatInterval: 25 * time.Second}
	fmt.Println(cl.GetRoomInfo(), cl.GetDanmuInfo())
	cl.ConnectDanmuServer()
	cl.RegHandler(blivedm.CmdDanmaku, func(context *blivedm.Context) {
		msg, _ := context.ToDanmakuMessage()
		fmt.Println(msg.Uname, msg.Msg)
	})
	cl.RegHandler(blivedm.CmdSendGift, func(context *blivedm.Context) {
		if msg, ok := context.ToGiftMessage(); ok {
			fmt.Println(msg.Uname, msg.Action, msg.GiftName)
		}
	})
	go func() {
		time.Sleep(time.Second * 100000)
		cl.Stop()
	}()
	cl.Start()
}
