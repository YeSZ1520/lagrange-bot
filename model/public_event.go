package model

type PushEvent struct {
	PostType      string `json:"post_type"`
	MessageType   string `json:"message_type"`
	NoticeType    string `json:"notice_type"`
	RequestType   string `json:"request_type"`
	MetaEventType string `json:"meta_event_type"`
	BaseEvent
}
