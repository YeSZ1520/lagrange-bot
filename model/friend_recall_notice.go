package model

type FriendRecallNotice struct {
	UserID    int64 `json:"user_id"`
	MessageID int64 `json:"message_id"`
	BaseEvent
}
