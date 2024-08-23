package model

type FriendRequest struct {
	UserID  int64  `json:"user_id"`
	Comment string `json:"comment"`
	Flag    string `json:"flag"`
	BaseEvent
}
