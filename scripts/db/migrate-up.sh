#!/bin/bash

# This script runs the database migration tool to apply the migrations.

# Load the environment variables
source .env

# Run the migration tool
go run cmd/migrate/main.go up