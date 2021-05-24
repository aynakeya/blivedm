# blivedm

使用golang 获取bilibili直播弹幕，使用websocket协议
 
go version of [xfgryujk/blivedm](https://github.com/xfgryujk/blivedm) 

## quick start

```
import (
	"github.com/LXG-Shadow/blivedm"
	"fmt"
	"time"
)

func main() {
	cl := blivedm.BLiveWsClient{ShortId: 3044248, Uid: 0, HearbeatInterval: 25 * time.Second}
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
	time.Sleep(100 * time.Second)
}
```