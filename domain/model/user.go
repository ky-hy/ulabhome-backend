package model

import (
	"time"
)

type UserID int64

type User struct {
	ID       UserID    `json:"id" db:"id"`
	CreateAt time.Time `json:"createdAt" db:"create_at"`
	UpdateAt time.Time `json:"updateAt" db:"update_at"`
}

type Users []*User
