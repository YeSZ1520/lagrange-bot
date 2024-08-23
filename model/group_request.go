package model

type GroupRequest struct {
	SubType string `json:"sub_type"`
	GroupID int64  `json:"group_id"`
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
	BaseEvent
}
