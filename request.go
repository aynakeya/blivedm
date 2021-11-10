package blivedm

type DanmakuSendForm struct {
	Bubble   int
	Message  string
	Color    string
	Mode     int
	Fontsize int
	Rnd      int
	RoomId   int
	CSRF     string
}
