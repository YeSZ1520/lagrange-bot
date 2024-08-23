package model

type PrivateMessage struct {
	SubType    string    `json:"sub_type"`
	MessageID  int64     `json:"message_id"`
	UserID     int64     `json:"user_id"`
	Message    []Segment `json:"message"`
	RawMessage string    `json:"raw_message"`
	Font       int32     `json:"font"`
	Sender     UserInfo  `json:"sender"`
	BaseEvent
}
