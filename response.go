package blivedm

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
}

type RoomInfoData struct {
	AllSpecialTypes []interface{} `json:"all_special_types"`
	Encrypted       bool          `json:"encrypted"`
	HiddenTill      int           `json:"hidden_till"`
	IsHidden        bool          `json:"is_hidden"`
	IsLocked        bool          `json:"is_locked"`
	IsPortrait      bool          `json:"is_portrait"`
	LiveStatus      int           `json:"live_status"`
	LiveTime        int           `json:"live_time"`
	LockTill        int           `json:"lock_till"`
	PlayurlInfo     interface{}   `json:"playurl_info"`
	PwdVerified     bool          `json:"pwd_verified"`
	RoomID          int           `json:"room_id"`
	RoomShield      int           `json:"room_shield"`
	ShortID         int           `json:"short_id"`
	UID             int           `json:"uid"`
}

type RoomInfoResponse struct {
	BaseResponse
	Data RoomInfoData `json:"data"`
}

type DanmuInfoData struct {
	BusinessID int    `json:"business_id"`
	Group      string `json:"group"`
	HostList   []struct {
		Host    string `json:"host"`
		Port    int    `json:"port"`
		WsPort  int    `json:"ws_port"`
		WssPort int    `json:"wss_port"`
	} `json:"host_list"`
	MaxDelay         int     `json:"max_delay"`
	RefreshRate      int     `json:"refresh_rate"`
	RefreshRowFactor float64 `json:"refresh_row_factor"`
	Token            string  `json:"token"`
}

type DanmuInfoResponse struct {
	BaseResponse
	Data DanmuInfoData `json:"data"`
}

type DanmakuSendResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
}
