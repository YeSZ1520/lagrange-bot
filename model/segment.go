package model

import "strings"

type Segment struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

type NodeSegmentData struct {
	UserID   string    `json:"user_id"`
	NickName string    `json:"nickname"`
	Content  []Segment `json:"content"`
}

type NodeSegment struct {
	Type string          `json:"type"`
	Data NodeSegmentData `json:"data"`
}

func ToCQCode(s Segment) string {
	cq := "[CQ:" + s.Type
	for k, v := range s.Data {
		if v != "" {
			v = strings.Replace(v, "&", "&amp;", -1)
			v = strings.Replace(v, "[", "&#91;", -1)
			v = strings.Replace(v, "]", "&#93;", -1)
			cq += "," + k + "=" + v
		}
	}
	cq += "]"
	return cq
}

func ToCQCodes(s []Segment) string {
	cqs := ""
	for _, seg := range s {
		cqs += ToCQCode(seg)
	}
	return cqs
}
