package main

import (
	// "log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"my-saas-app/internal/domain/entities"
	// "my-saas-app/internal/infrastructure/database"
	// "my-saas-app/scripts/db"
)

func main() {
	dsn := "root:0000@tcp(localhost:3306)/my_saas_app?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.FlagCard{})
	db.AutoMigrate(&entities.Company{})
	db.AutoMigrate(&entities.Bank{})
	db.AutoMigrate(&entities.CreditCard{})
	db.AutoMigrate(&entities.Remuneration{})
	// dbConn, err := database.Connect()
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }

	// migrator := db.NewMigrator(dbConn)
	// if err := migrator.MigrateUp(); err != nil {
	// 	log.Fatalf("failed to apply migrations: %v", err)
	// }
}
