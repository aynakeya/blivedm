package blivedm

type DanmakuSendForm struct {
	Bubble   int
	Message  string
	Color    string
	mode     int
	Fontsize int
	Rnd      int
	RoomId   int
	CSRF     string
}
