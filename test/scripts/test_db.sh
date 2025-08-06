#!/bin/bash

# Kill and remove old container if it exists
docker rm -f postgres-test 2>/dev/null || true

# Start new PostgreSQL container
docker run --name postgres-test \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 6432:5432 \
  -d postgres:latest

echo "PostgreSQL starts..."
sleep 3

# Create database
winpty docker exec -it postgres-test \
  psql -U postgres -d postgres \
  -c "CREATE DATABASE productapp"
echo "Database productapp created..."
sleep 3

# Create table
winpty docker exec -it postgres-test \
  psql -U postgres -d productapp \
  -c "
  CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DOUBLE PRECISION NOT NULL,
    discount DOUBLE PRECISION,
    store VARCHAR(255) NOT NULL
  );"
echo "Table products created..."
