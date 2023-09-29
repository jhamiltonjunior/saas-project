#!/bin/bash

# This script is used to run the database migration tool to rollback the migrations.

# Load the environment variables
source .env

# Run the migration tool to rollback the migrations
go run cmd/migrate/main.go -dir db/migrations -database "$DATABASE_URL" down