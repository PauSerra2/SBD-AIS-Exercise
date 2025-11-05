#!/bin/bash
docker run -d \
  --name orders-db \
  -e POSTGRES_DB=order \
  -e POSTGRES_USER=docker \
  -e POSTGRES_PASSWORD=docker \
  -p 5432:5432 \
  -v pgdata:/var/lib/postgresql \
  postgres:18

echo "✅ PostgreSQL container started"