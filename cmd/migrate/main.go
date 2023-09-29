package main

import (
	"log"

	"github.com/my-saas-app/internal/infrastructure/database"
	"github.com/my-saas-app/scripts/db"
)

func main() {
	dbConn, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	migrator := db.NewMigrator(dbConn)
	if err := migrator.MigrateUp(); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}
}