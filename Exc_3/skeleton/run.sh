#!/bin/sh

# --- Configuration ---
IMAGE_NAME="ordersystem"
DB_CONTAINER="db"
APP_CONTAINER="orderservice"
DB_PASSWORD="docker"
DB_USER="docker"
DB_NAME="order"
DB_PORT=5432
APP_PORT=8080

# --- Cleanup old containers ---
echo "🧹 Cleaning up old containers..."
docker rm -f $DB_CONTAINER $APP_CONTAINER 2>/dev/null

# --- Build the image ---
echo "⚙️  Building Docker image..."
docker build -t $IMAGE_NAME .

# --- Run PostgreSQL container ---
echo "🐘 Starting PostgreSQL..."
docker run -d \
  --name $DB_CONTAINER \
  -e POSTGRES_USER=$DB_USER \
  -e POSTGRES_PASSWORD=$DB_PASSWORD \
  -e POSTGRES_DB=$DB_NAME \
  -p $DB_PORT:5432 \
  postgres:15

# --- Wait for DB to start ---
echo "⏳ Waiting for database to be ready..."
sleep 5

# --- Run the Go application container ---
echo "🚀 Starting $APP_CONTAINER..."
docker run -d \
  --name $APP_CONTAINER \
  --link $DB_CONTAINER:db \
  -e POSTGRES_USER=$DB_USER \
  -e POSTGRES_PASSWORD=$DB_PASSWORD \
  -e POSTGRES_DB=$DB_NAME \
  -e POSTGRES_TCP_PORT=$DB_PORT \
  -e DB_HOST=db \
  -p $APP_PORT:8080 \
  $IMAGE_NAME

# --- Status ---
echo "✅ Containers are up and running!"
docker ps
echo ""
echo "🌍 Application: http://localhost:$APP_PORT"
echo "🐘 Database: localhost:$DB_PORT (user=$DB_USER, password=$DB_PASSWORD, db=$DB_NAME)"
