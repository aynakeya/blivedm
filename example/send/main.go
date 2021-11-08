package main

import (
	"bufio"
	"fmt"
	"github.com/aynakeya/blivedm"
	"os"
	"time"
)

func main() {
	cl := blivedm.BLiveWsClient{ShortId: 544957,
		Account: blivedm.DanmuAccount{
			UID:         114514,
			SessionData: "123",
			BilibiliJCT: "3333",
		},
		HearbeatInterval: 25 * time.Second}
	fmt.Println(cl.GetRoomInfo(), cl.GetDanmuInfo())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		message, err := cl.SendMessage(blivedm.DanmakuSendForm{
			Bubble:   0,
			Message:  input,
			Color:    "16777215",
			Fontsize: 25,
			Rnd:      int(time.Now().Unix()),
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(message)
	}
}
