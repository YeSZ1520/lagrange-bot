package model

type HeartBeat struct {
	Status   map[string]bool `json:"status"`
	Interval int64           `json:"interval"`
	BaseEvent
}
