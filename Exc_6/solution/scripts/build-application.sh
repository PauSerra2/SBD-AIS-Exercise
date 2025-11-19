#!/bin/sh
set -e
cd /app

# Baixar i ajustar dependències
go mod tidy

# Compilar binari per Linux
CGO_ENABLED=0 GOOS=linux go build -o /app/ordersystem
