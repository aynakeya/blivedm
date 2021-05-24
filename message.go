package blivedm

type DanmakuMessage struct {
	Mode      int64
	FontSize  int64
	Color     int64
	Timestamp int64
	Rnd       int64
	UID_CRC32 string
	MsgType   int64
	Bubble    int64

	Msg string

	Uid          int64
	Uname        string
	Admin        bool
	Vip          bool
	Svip         bool
	Urank        int64
	MobileVerify bool
	UnameColor   string

	MedalLevel   int64
	MedalName    string
	MedalUpName  string
	MedalRoomId  int64
	MedalColor   int64
	SpecialMedal string

	UserLevel      int64
	UserLevelColor int64
	UserLevelRank  string

	OldTitle string
	Title    string

	PrivilegeType int64
}

type GiftMessage struct {
	Action           string      `json:"action"`
	BatchComboID     string      `json:"batch_combo_id"`
	BatchComboSend   interface{} `json:"batch_combo_send"`
	BeatID           string      `json:"beatId"`
	BizSource        string      `json:"biz_source"`
	BlindGift        interface{} `json:"blind_gift"`
	BroadcastID      int         `json:"broadcast_id"`
	CoinType         string      `json:"coin_type"`
	ComboResourcesID int         `json:"combo_resources_id"`
	ComboSend        interface{} `json:"combo_send"`
	ComboStayTime    int         `json:"combo_stay_time"`
	ComboTotalCoin   int         `json:"combo_total_coin"`
	CritProb         int         `json:"crit_prob"`
	Demarcation      int         `json:"demarcation"`
	Dmscore          int         `json:"dmscore"`
	Draw             int         `json:"draw"`
	Effect           int         `json:"effect"`
	EffectBlock      int         `json:"effect_block"`
	Face             string      `json:"face"`
	GiftID           int         `json:"giftId"`
	GiftName         string      `json:"giftName"`
	GiftType         int         `json:"giftType"`
	Gold             int         `json:"gold"`
	GuardLevel       int         `json:"guard_level"`
	IsFirst          bool        `json:"is_first"`
	IsSpecialBatch   int         `json:"is_special_batch"`
	Magnification    int         `json:"magnification"`
	MedalInfo        struct {
		AnchorRoomid     int    `json:"anchor_roomid"`
		AnchorUname      string `json:"anchor_uname"`
		GuardLevel       int    `json:"guard_level"`
		IconID           int    `json:"icon_id"`
		IsLighted        int    `json:"is_lighted"`
		MedalColor       int    `json:"medal_color"`
		MedalColorBorder int    `json:"medal_color_border"`
		MedalColorEnd    int    `json:"medal_color_end"`
		MedalColorStart  int    `json:"medal_color_start"`
		MedalLevel       int    `json:"medal_level"`
		MedalName        string `json:"medal_name"`
		Special          string `json:"special"`
		TargetID         int    `json:"target_id"`
	} `json:"medal_info"`
	NameColor         string      `json:"name_color"`
	Num               int         `json:"num"`
	OriginalGiftName  string      `json:"original_gift_name"`
	Price             int         `json:"price"`
	Rcost             int         `json:"rcost"`
	Remain            int         `json:"remain"`
	Rnd               string      `json:"rnd"`
	SendMaster        interface{} `json:"send_master"`
	Silver            int         `json:"silver"`
	Super             int         `json:"super"`
	SuperBatchGiftNum int         `json:"super_batch_gift_num"`
	SuperGiftNum      int         `json:"super_gift_num"`
	SvgaBlock         int         `json:"svga_block"`
	TagImage          string      `json:"tag_image"`
	Tid               string      `json:"tid"`
	Timestamp         int         `json:"timestamp"`
	TopList           interface{} `json:"top_list"`
	TotalCoin         int         `json:"total_coin"`
	UID               int         `json:"uid"`
	Uname             string      `json:"uname"`
}
