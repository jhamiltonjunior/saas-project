package entities

import (
	"time"
)

type Recurrence struct {
	ID       int       `json:"id" gorm:"primary_key;type:int(11);not null;auto_increment;unique_index"`
	Name     string    `json:"name" gorm:"type:enum('Único', 'Diario', 'Semanal', 'Mensal', 'Anual');not null"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime;type:DATETIME;default:CURRENT_TIMESTAMP;not null"`
	UpdateAt time.Time `json:"update_at" gorm:"type:datetime; default:null"`
	DeleteAt time.Time `json:"delete_at" gorm:"type:datetime; default:null"`
	Active   uint8     `json:"active" gorm:"default:1;type:tinyint(1);not null;"`
}
