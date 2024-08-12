package entities

import (
	"time"
)

type Bank struct {
	ID       int       `json:"id" gorm:"primary_key;type:int(11);not null;auto_increment"`
	Name     string    `json:"name" gorm:"type:varchar(100);not null"`
	Image    string    `json:"image" gorm:"type:varchar(250);null"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime;type:DATETIME;not null"`
	UpdateAt time.Time `json:"update_at" gorm:"type:datetime; default:null"`
	DeleteAt time.Time `json:"delete_at" gorm:"type:datetime; default:null"`
	Active   int       `json:"active" gorm:"default:1;type:tinyint(1);not null;"`
}
