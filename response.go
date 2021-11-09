package blivedm

type BaseV1Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
}

type RoomInfoV2Data struct {
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

type RoomInfoV2Response struct {
	BaseResponse
	Data RoomInfoV2Data `json:"data"`
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
	BaseV1Response
}

type RoomInfoData struct {
	Uid              int      `json:"uid"`
	RoomId           int      `json:"room_id"`
	ShortId          int      `json:"short_id"`
	Attention        int      `json:"attention"`
	Online           int      `json:"online"`
	IsPortrait       bool     `json:"is_portrait"`
	Description      string   `json:"description"`
	LiveStatus       int      `json:"live_status"`
	AreaId           int      `json:"area_id"`
	ParentAreaId     int      `json:"parent_area_id"`
	ParentAreaName   string   `json:"parent_area_name"`
	OldAreaId        int      `json:"old_area_id"`
	Background       string   `json:"background"`
	Title            string   `json:"title"`
	UserCover        string   `json:"user_cover"`
	Keyframe         string   `json:"keyframe"`
	IsStrictRoom     bool     `json:"is_strict_room"`
	LiveTime         string   `json:"live_time"`
	Tags             string   `json:"tags"`
	IsAnchor         int      `json:"is_anchor"`
	RoomSilentType   string   `json:"room_silent_type"`
	RoomSilentLevel  int      `json:"room_silent_level"`
	RoomSilentSecond int      `json:"room_silent_second"`
	AreaName         string   `json:"area_name"`
	Pendants         string   `json:"pendants"`
	AreaPendants     string   `json:"area_pendants"`
	HotWords         []string `json:"hot_words"`
	HotWordsStatus   int      `json:"hot_words_status"`
	Verify           string   `json:"verify"`
	NewPendants      struct {
		Frame struct {
			Name       string `json:"name"`
			Value      string `json:"value"`
			Position   int    `json:"position"`
			Desc       string `json:"desc"`
			Area       int    `json:"area"`
			AreaOld    int    `json:"area_old"`
			BgColor    string `json:"bg_color"`
			BgPic      string `json:"bg_pic"`
			UseOldArea bool   `json:"use_old_area"`
		} `json:"frame"`
		Badge       interface{} `json:"badge"`
		MobileFrame struct {
			Name       string `json:"name"`
			Value      string `json:"value"`
			Position   int    `json:"position"`
			Desc       string `json:"desc"`
			Area       int    `json:"area"`
			AreaOld    int    `json:"area_old"`
			BgColor    string `json:"bg_color"`
			BgPic      string `json:"bg_pic"`
			UseOldArea bool   `json:"use_old_area"`
		} `json:"mobile_frame"`
		MobileBadge interface{} `json:"mobile_badge"`
	} `json:"new_pendants"`
	UpSession            string `json:"up_session"`
	PkStatus             int    `json:"pk_status"`
	PkId                 int    `json:"pk_id"`
	BattleId             int    `json:"battle_id"`
	AllowChangeAreaTime  int    `json:"allow_change_area_time"`
	AllowUploadCoverTime int    `json:"allow_upload_cover_time"`
	StudioInfo           struct {
		Status     int           `json:"status"`
		MasterList []interface{} `json:"master_list"`
	} `json:"studio_info"`
}

type RoomInfoResponse struct {
	BaseV1Response
	Data RoomInfoData `json:"data"`
}

type UpInfoData struct {
	Mid       int    `json:"mid"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Face      string `json:"face"`
	Sign      string `json:"sign"`
	Rank      int    `json:"rank"`
	Level     int    `json:"level"`
	Jointime  int    `json:"jointime"`
	Moral     int    `json:"moral"`
	Silence   int    `json:"silence"`
	Coins     int    `json:"coins"`
	FansBadge bool   `json:"fans_badge"`
	FansMedal struct {
		Show  bool        `json:"show"`
		Wear  bool        `json:"wear"`
		Medal interface{} `json:"medal"`
	} `json:"fans_medal"`
	Official struct {
		Role  int    `json:"role"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
		Type  int    `json:"type"`
	} `json:"official"`
	Vip struct {
		Type       int   `json:"type"`
		Status     int   `json:"status"`
		DueDate    int64 `json:"due_date"`
		VipPayType int   `json:"vip_pay_type"`
		ThemeType  int   `json:"theme_type"`
		Label      struct {
			Path        string `json:"path"`
			Text        string `json:"text"`
			LabelTheme  string `json:"label_theme"`
			TextColor   string `json:"text_color"`
			BgStyle     int    `json:"bg_style"`
			BgColor     string `json:"bg_color"`
			BorderColor string `json:"border_color"`
		} `json:"label"`
		AvatarSubscript    int    `json:"avatar_subscript"`
		NicknameColor      string `json:"nickname_color"`
		Role               int    `json:"role"`
		AvatarSubscriptUrl string `json:"avatar_subscript_url"`
	} `json:"vip"`
	Pendant struct {
		Pid               int    `json:"pid"`
		Name              string `json:"name"`
		Image             string `json:"image"`
		Expire            int    `json:"expire"`
		ImageEnhance      string `json:"image_enhance"`
		ImageEnhanceFrame string `json:"image_enhance_frame"`
	} `json:"pendant"`
	Nameplate struct {
		Nid        int    `json:"nid"`
		Name       string `json:"name"`
		Image      string `json:"image"`
		ImageSmall string `json:"image_small"`
		Level      string `json:"level"`
		Condition  string `json:"condition"`
	} `json:"nameplate"`
	UserHonourInfo struct {
		Mid    int           `json:"mid"`
		Colour interface{}   `json:"colour"`
		Tags   []interface{} `json:"tags"`
	} `json:"user_honour_info"`
	IsFollowed bool   `json:"is_followed"`
	TopPhoto   string `json:"top_photo"`
	Theme      struct {
	} `json:"theme"`
	SysNotice struct {
	} `json:"sys_notice"`
	LiveRoom struct {
		RoomStatus    int    `json:"roomStatus"`
		LiveStatus    int    `json:"liveStatus"`
		Url           string `json:"url"`
		Title         string `json:"title"`
		Cover         string `json:"cover"`
		Online        int    `json:"online"`
		Roomid        int    `json:"roomid"`
		RoundStatus   int    `json:"roundStatus"`
		BroadcastType int    `json:"broadcast_type"`
	} `json:"live_room"`
	Birthday string `json:"birthday"`
	School   struct {
		Name string `json:"name"`
	} `json:"school"`
	Profession struct {
		Name string `json:"name"`
	} `json:"profession"`
	Tags   interface{} `json:"tags"`
	Series struct {
		UserUpgradeStatus int  `json:"user_upgrade_status"`
		ShowUpgradeWindow bool `json:"show_upgrade_window"`
	} `json:"series"`
}

type UpInfoResponse struct {
	BaseResponse
	Data UpInfoData `json:"data"`
}
