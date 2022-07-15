package entity

import (
	"encoding/json"
	"time"
)

type Work struct {
	ID          int64           `xorm:"id pk autoincr" `
	Date        string          `xorm:"date" `
	ProjectName string          `xorm:"project_name" `
	TechID      int64           `xorm:"tech_id" `
	TechName    string          `xorm:"tech_name" `
	LinkURL     string          `xorm:"link_url" `
	Content     string          `xorm:"content" `
	Retrospect  string          `xorm:"retrospect" `
	Created     json.RawMessage `xorm:"json 'created'" json:"created"`
	Updated     json.RawMessage `xorm:"json 'updated'" json:"updated"`
	DeletedAt   time.Time       `xorm:"deleted" `
}

func (Work) TableName() string {
	return "works"
}
