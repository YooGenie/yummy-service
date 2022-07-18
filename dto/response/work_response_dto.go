package dto

import (
	"encoding/json"
)

type WorkSummary struct {
	Id          int64           `json:"id" `
	Date        string          `json:"date"`
	ProjectName string          `json:"projectName"`
	TechId      int64           `json:"techID" `
	TechName    string          `json:"techName`
	LinkUrl     string          `json:"linkURL"`
	Content     string          `json:"content"`
	Retrospect  string          `json:"retrospect"`
	Created     json.RawMessage `xorm:"json 'created'" json:"created"`
	Updated     json.RawMessage `xorm:"json 'updated'" json:"updated"`
}
