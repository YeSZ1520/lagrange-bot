package model

type GroupRecallNotice struct {
	GroupID    int64 `json:"group_id"`
	UserID     int64 `json:"user_id"`
	OperatorID int64 `json:"operator_id"`
	MessageID  int64 `json:"message_id"`
	BaseEvent
}
