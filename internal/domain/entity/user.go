package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(320);unique;not null"`
	Password  string    `json:"password" gorm:"type:varchar(72);not null"`
	IsAdmin   bool      `json:"is_admin" gorm:"type:boolean;not null;default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;autoUpdateTime"`
}
