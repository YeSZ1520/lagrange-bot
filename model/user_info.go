package model

type UserInfo struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}
