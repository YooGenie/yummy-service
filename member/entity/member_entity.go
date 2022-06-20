package entity

import (
	"time"
)

type Member struct {
	Id        int64     `xorm:"id pk autoincr" `
	Email     string    `xorm:"email" `
	Password  string    `xorm:"password" `
	Name      string    `xorm:"name" `
	Mobile    string    `xorm:"mobile" `
	Role      string    `xorm:"role" `
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at" `
	DeletedAt time.Time `xorm:"deleted" `
}

func (Member) TableName() string {
	return "members"
}
