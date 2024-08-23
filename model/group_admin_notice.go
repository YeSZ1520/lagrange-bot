package model

type GroupAdminNotice struct {
	SubType string `json:"sub_type"`
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	BaseEvent
}
