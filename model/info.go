package model

type Info struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    Data   `json:"data"`
}
type Medal struct {
	UID              int    `json:"uid"`
	TargetID         int    `json:"target_id"`
	MedalID          int    `json:"medal_id"`
	Level            int    `json:"level"`
	MedalName        string `json:"medal_name"`
	MedalColor       int    `json:"medal_color"`
	Intimacy         int    `json:"intimacy"`
	NextIntimacy     int    `json:"next_intimacy"`
	DayLimit         int    `json:"day_limit"`
	MedalColorStart  int    `json:"medal_color_start"`
	MedalColorEnd    int    `json:"medal_color_end"`
	MedalColorBorder int    `json:"medal_color_border"`
	IsLighted        int    `json:"is_lighted"`
	GuardLevel       int    `json:"guard_level"`
	LightStatus      int    `json:"light_status"`
	WearingStatus    int    `json:"wearing_status"`
	Score            int    `json:"score"`
}
type FansMedal struct {
	Show  bool  `json:"show"`
	Wear  bool  `json:"wear"`
	Medal Medal `json:"medal"`
}
type Official struct {
	Role  int    `json:"role"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Type  int    `json:"type"`
}
type Label struct {
	Path                  string `json:"path"`
	Text                  string `json:"text"`
	LabelTheme            string `json:"label_theme"`
	TextColor             string `json:"text_color"`
	BgStyle               int    `json:"bg_style"`
	BgColor               string `json:"bg_color"`
	BorderColor           string `json:"border_color"`
	UseImgLabel           bool   `json:"use_img_label"`
	ImgLabelURIHans       string `json:"img_label_uri_hans"`
	ImgLabelURIHant       string `json:"img_label_uri_hant"`
	ImgLabelURIHansStatic string `json:"img_label_uri_hans_static"`
	ImgLabelURIHantStatic string `json:"img_label_uri_hant_static"`
}
type Vip struct {
	Type               int    `json:"type"`
	Status             int    `json:"status"`
	DueDate            int64  `json:"due_date"`
	VipPayType         int    `json:"vip_pay_type"`
	ThemeType          int    `json:"theme_type"`
	Label              Label  `json:"label"`
	AvatarSubscript    int    `json:"avatar_subscript"`
	NicknameColor      string `json:"nickname_color"`
	Role               int    `json:"role"`
	AvatarSubscriptURL string `json:"avatar_subscript_url"`
	TvVipStatus        int    `json:"tv_vip_status"`
	TvVipPayType       int    `json:"tv_vip_pay_type"`
}
type Pendant struct {
	Pid               int    `json:"pid"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	Expire            int    `json:"expire"`
	ImageEnhance      string `json:"image_enhance"`
	ImageEnhanceFrame string `json:"image_enhance_frame"`
}
type Nameplate struct {
	Nid        int    `json:"nid"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	ImageSmall string `json:"image_small"`
	Level      string `json:"level"`
	Condition  string `json:"condition"`
}
type UserHonourInfo struct {
	Mid    int           `json:"mid"`
	Colour interface{}   `json:"colour"`
	Tags   []interface{} `json:"tags"`
}
type Theme struct {
}
type SysNotice struct {
}
type WatchedShow struct {
	Switch       bool   `json:"switch"`
	Num          int    `json:"num"`
	TextSmall    string `json:"text_small"`
	TextLarge    string `json:"text_large"`
	Icon         string `json:"icon"`
	IconLocation string `json:"icon_location"`
	IconWeb      string `json:"icon_web"`
}
type LiveRoom struct {
	RoomStatus    int         `json:"roomStatus"`
	LiveStatus    int         `json:"liveStatus"`
	URL           string      `json:"url"`
	Title         string      `json:"title"`
	Cover         string      `json:"cover"`
	Roomid        int         `json:"roomid"`
	RoundStatus   int         `json:"roundStatus"`
	BroadcastType int         `json:"broadcast_type"`
	WatchedShow   WatchedShow `json:"watched_show"`
}
type School struct {
	Name string `json:"name"`
}
type Profession struct {
	Name       string `json:"name"`
	Department string `json:"department"`
	Title      string `json:"title"`
	IsShow     int    `json:"is_show"`
}
type Series struct {
	UserUpgradeStatus int  `json:"user_upgrade_status"`
	ShowUpgradeWindow bool `json:"show_upgrade_window"`
}
type Data struct {
	Mid            int64          `json:"mid"`
	Name           string         `json:"name"`
	Sex            string         `json:"sex"`
	Face           string         `json:"face"`
	FaceNft        int            `json:"face_nft"`
	FaceNftType    int            `json:"face_nft_type"`
	Sign           string         `json:"sign"`
	Rank           int            `json:"rank"`
	Level          int            `json:"level"`
	Jointime       int            `json:"jointime"`
	Moral          int            `json:"moral"`
	Silence        int            `json:"silence"`
	Coins          int            `json:"coins"`
	FansBadge      bool           `json:"fans_badge"`
	FansMedal      FansMedal      `json:"fans_medal"`
	Official       Official       `json:"official"`
	Vip            Vip            `json:"vip"`
	Pendant        Pendant        `json:"pendant"`
	Nameplate      Nameplate      `json:"nameplate"`
	UserHonourInfo UserHonourInfo `json:"user_honour_info"`
	IsFollowed     bool           `json:"is_followed"`
	TopPhoto       string         `json:"top_photo"`
	Theme          Theme          `json:"theme"`
	SysNotice      SysNotice      `json:"sys_notice"`
	LiveRoom       LiveRoom       `json:"live_room"`
	Birthday       string         `json:"birthday"`
	School         School         `json:"school"`
	Profession     Profession     `json:"profession"`
	Tags           interface{}    `json:"tags"`
	Series         Series         `json:"series"`
	IsSeniorMember int            `json:"is_senior_member"`
	McnInfo        interface{}    `json:"mcn_info"`
}
