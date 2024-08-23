package model

type GroupUploadNotice struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	File    file  `json:"file"`
	BaseEvent
}

type file struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	BusID int64  `json:"busid"`
}
