#!/bin/bash

# Set environment variables for local development
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=postgres
export DB_PORT=5432

echo "Environment variables set for local development:"
echo "DB_HOST: $DB_HOST"
echo "DB_USER: $DB_USER"
echo "DB_NAME: $DB_NAME"
echo "DB_PORT: $DB_PORT"

echo ""
echo "Starting application..."
go run main.go 