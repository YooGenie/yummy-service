package entity

import (
	"encoding/json"
	"time"
)

type Member struct {
	Id        int64           `xorm:"id pk autoincr" `
	Email     string          `xorm:"email" `
	Password  string          `xorm:"password" `
	Name      string          `xorm:"name" `
	Mobile    string          `xorm:"mobile" `
	Closed    bool            `xorm:"closed" `
	Created   json.RawMessage `xorm:"json 'created'"`
	Updated   json.RawMessage `xorm:"json 'updated'" `
	DeletedAt time.Time       `xorm:"deleted" `
}

func (Member) TableName() string {
	return "members"
}
