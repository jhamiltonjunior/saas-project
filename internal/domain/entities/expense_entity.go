package entities

import (
	"time"
)

type ExpenseID int

type Expense struct {
	ID           ExpenseID `json:"id" gorm:"primary_key;type:int(11);not null;auto_increment;unique_index"`
	Name         string    `json:"name" gorm:"type:varchar(100);not null"`
	Value        float64   `json:"value" gorm:"type:decimal(10,2);not null"`
	CreateAt     time.Time `json:"create_at" gorm:"autoCreateTime;type:DATETIME;default:CURRENT_TIMESTAMP;not null"`
	UpdateAt     time.Time `json:"update_at" gorm:"type:datetime; default:null"`
	DeleteAt     time.Time `json:"delete_at" gorm:"type:datetime; default:null"`
	Active       uint8     `json:"active" gorm:"default:1;type:tinyint(1);not null;"`
	PaidOut      uint8     `json:"paid_out" gorm:"default:0;type:tinyint(1);not null;"`
	UserID       int32     `json:"user_id" gorm:"type:int(11);not null;"`
	RecurrenceID int32     `json:"recurrence_id" gorm:"type:int(11);not null;"`
	User         User
	Recurrence   Recurrence
}

type CreditCardExpense struct {
	ID           int   `json:"id" gorm:"primary_key;type:int(11);not null;auto_increment;unique_index"`
	CreditCardID int32 `json:"credit_card_id" gorm:"type:int(11);not null;"`
	ExpenseID    int32 `json:"expense_id" gorm:"type:int(11);not null;"`
	CreditCard   CreditCard
	Expense      Expense
}
