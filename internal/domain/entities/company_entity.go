package entities

import (
	"time"
)

type Company struct {
	ID       int       `json:"id" gorm:"primary_key;type:int(11);not null;auto_increment"`
	Name     string    `json:"name" gorm:"type:varchar(100)"`
	CNPJ     string    `json:"cnpj" gorm:"type:varchar(100);unique"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime;type:DATETIME"`
	UpdateAt time.Time `json:"update_at" gorm:"type:datetime; default:null"`
	DeleteAt time.Time `json:"delete_at" gorm:"type:datetime; default:null"`
	Active   int       `json:"active" gorm:"default:1;type:tinyint(1);not null;"`
}
