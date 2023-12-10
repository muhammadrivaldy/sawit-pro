package models

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	Id          int       `json:"id" gorm:"column:id"`
	FullName    string    `json:"full_name" gorm:"column:full_name"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number"`
	CreatedBy   int       `json:"created_by" gorm:"column:created_by"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedBy   null.Int  `json:"updated_by" gorm:"column:updated_by"`
	UpdatedAt   null.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   null.Time `json:"deleted_at" gorm:"column:deleted_at"`
}
