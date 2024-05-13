package entities

import (
	"time"

	"gorm.io/gorm"
)

type CreditCard struct {
	ID        int32          `json:"id" gorm:"primary_key;type:int(11);not null;auto_increment;unique_index"`
	Name      string         `json:"name" gorm:"type:varchar(200);not null;"`
	Value     float64        `json:"value" gorm:"type:decimal(10,2);not null;"`
	DueDate   uint8          `json:"due_date" gorm:"type:tinyint(2);not null;"`
	Active    uint8          `json:"active" gorm:"default:1;type:tinyint(1);not null;"`
	UserID    int32          `json:"user_id" gorm:"type:int(11);not null;"`
	FlagID    int32          `json:"flag_id" gorm:"type:int(11);not null;"`
	BankID    int32          `json:"bank_id" gorm:"type:int(11);not null;"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP; not null"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User      User
	Flag      FlagCard
	Bank      Bank
}
