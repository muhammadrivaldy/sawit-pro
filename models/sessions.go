package models

import (
	"time"
)

type Session struct {
	Id      int       `json:"id" gorm:"column:id"`
	UserId  int       `json:"user_id" gorm:"column:user_id"`
	LoginAt time.Time `json:"login_at" gorm:"column:login_at"`
}

func (Session) TableName() string {
	return "trx_sessions"
}
