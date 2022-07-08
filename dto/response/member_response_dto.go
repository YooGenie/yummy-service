package dto

import (
	"time"
)

type MemberSummary struct {
	ID        int64     `xorm:"id" json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
