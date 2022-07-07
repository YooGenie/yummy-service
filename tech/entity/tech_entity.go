package entity

import (
	"encoding/json"
	"time"
)

type Tech struct {
	Id        int64           `xorm:"id pk autoincr" `
	Name      string          `xorm:"name" `
	Created   json.RawMessage `xorm:"json 'created'" json:"created"`
	Updated   json.RawMessage `xorm:"json 'updated'" json:"updated"`
	DeletedAt time.Time       `xorm:"deleted" `
}

func (Tech) TableName() string {
	return "techs"
}
