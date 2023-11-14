package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func NewGormConnection (dsn string) (*gorm.DB, error) {
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}