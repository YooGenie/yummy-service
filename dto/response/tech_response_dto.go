package dto

import (
	"encoding/json"
)

type TechSummary struct {
	Id      int64           `json:"id"`
	Name    string          `json:"name"`
	Created json.RawMessage `xorm:"json 'created'" json:"created"`
	Updated json.RawMessage `xorm:"json 'updated'" json:"updated"`
}
