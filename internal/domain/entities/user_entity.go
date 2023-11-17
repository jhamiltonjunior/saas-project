package entities

import (
	"time"
)

type User struct {
	ID       int       `json:"id" gorm:"primaryKey;autoIncrement:true;not null;unique_index"`
	Name     string    `json:"name" gorm:"type:varchar(200);not null"`
	Email    string    `json:"email" gorm:"type:varchar(200);not null;unique_index"`
	Password string    `json:"password" gorm:"type:varchar(200);not null"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time `json:"update_at" gorm:"type:timestamp; null"`
	DeleteAt time.Time `json:"delete_at" gorm:"type:timestamp; null"`
	Active   bool      `json:"active" gorm:"default:true"`
}
