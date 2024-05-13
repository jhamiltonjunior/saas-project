package entities

import (
	"time"

	"gorm.io/gorm"
)

type FlagCard struct {
	ID        int32          `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string         `json:"name" gorm:"type:varchar(255); not null"`
	Active    uint8          `json:"active" gorm:"type:tinyint(1); not null"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
