package model

type ApiResult[T any] struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
	Data    T      `json:"data"`
	Echo    string `json:"echo"`
}

type Cookies struct {
	Cookies string `json:"Cookies"`
}

type CsrfToken struct {
	CsrfToken int32 `json:"token"`
}

type Credentials struct {
	Cookies   string `json:"Cookies"`
	CsrfToken int32  `json:"csrf_token"`
}

type FriendInfo struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

type LoginInfo struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
}

type GroupInfo struct {
	GroupId        int64  `json:"group_id"`
	GroupName      string `json:"group_name"`
	MemberCount    int32  `json:"member_count"`
	MaxMemberCount int32  `json:"max_member_count"`
}

type GroupMemberInfo struct {
	GroupId         int64  `json:"group_id"`
	UserId          int64  `json:"user_id"`
	Nickname        string `json:"nickname"`
	Card            string `json:"card"`
	Sex             string `json:"sex"`
	Age             int32  `json:"age"`
	Area            string `json:"area"`
	JoinTime        int32  `json:"join_time"`
	LastSendTime    int32  `json:"last_sent_time"`
	Level           string `json:"level"`
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int64  `json:"title_expire_time"`
	CardChangeAble  bool   `json:"card_changeable"`
}

type CanSend struct {
	Yes bool `json:"yes"`
}

type LagrangeStatus struct {
	AppInitialized bool `json:"app_initialized"`
	AppEnabled     bool `json:"app_enabled"`
	PluginsGood    bool `json:"plugins_good"`
	AppGood        bool `json:"app_good"`
	Online         bool `json:"online"`
	Good           bool `json:"good"`
	Memory         int  `json:"memory"`
}

type LagrangeVersionInfo struct {
	AppName         string `json:"app_name"`
	AppVersion      string `json:"app_version"`
	ProtocolVersion string `json:"protocol_version"`
	NtProtocol      string `json:"nt_protocol"`
}

type CurrentTalkative struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
	DayCount int32  `json:"day_count"`
}

type TypeList struct {
	UserID      int64  `json:"user_id"`
	NickName    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type GroupHonorInfo struct {
	GroupID          int64            `json:"group_id"`
	CurrentTalkative CurrentTalkative `json:"current_talkative"`
	TalkativeList    []TypeList       `json:"talkative_list"`
	PerformerList    []TypeList       `json:"performer_list"`
	LegendList       []TypeList       `json:"legend_list"`
	StrongNewbieList []TypeList       `json:"strong_newbie_list"`
	EmotionList      []TypeList       `json:"emotion_list"`
}

type MessageId struct {
	MessageId int64 `json:"message_id"`
}

type ForwardResult struct {
	MessageId int64  `json:"message_id"`
	ForwardId string `json:"forward_id"`
}

type Message struct {
	Time        int64     `json:"time"`
	MessageType string    `json:"message_type"`
	MessageID   int64     `json:"message_id"`
	RealID      int64     `json:"real_id"`
	Sender      UserInfo  `json:"sender"`
	Message     []Segment `json:"message"`
}

type LagrangeNode struct {
	Name    string    `json:"name"`
	Uin     string    `json:"uin"`
	Content []Segment `json:"content"`
}

type ResID struct {
	ResID int64 `json:"resid"`
}
