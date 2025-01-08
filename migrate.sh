#!/bin/bash
set -e

# Run migrations
echo "Running database migrations..."
go run migrate/migrate.go

# Start your application
echo "Starting application..."
exec "$@"