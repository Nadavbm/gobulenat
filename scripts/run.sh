#!/bin/sh
export DATABASE_USER=postgres
export DATABASE_PASSWORD=bulenat1234
export DATABASE_DB=gobulenat
export DATABASE_PORT=5432
export DATABASE_HOST=localhost

docker-compose down
docker-compose up -d pgsql

sleep 10

go run api/run/main.go