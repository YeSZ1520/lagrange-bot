package model

type GroupMessage struct {
	SubType    string      `json:"sub_type"`
	MessageID  int32       `json:"message_id"`
	GroupID    int64       `json:"group_id"`
	UserID     int64       `json:"user_id"`
	Anonymous  anonymous   `json:"anonymous"`
	Message    []Segment   `json:"message"`
	RawMessage string      `json:"raw_message"`
	Font       int32       `json:"font"`
	Sender     groupSender `json:"sender"`
	BaseEvent
}

type anonymous struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

type groupSender struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Card     string `json:"card"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	Area     string `json:"area"`
	Level    string `json:"level"`
	Role     string `json:"role"`
	Title    string `json:"title"`
}
