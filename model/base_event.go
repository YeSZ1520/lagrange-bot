package model

type BaseEvent struct {
	Time   int64 `json:"time"`
	SelfID int64 `json:"self_id"`
}
