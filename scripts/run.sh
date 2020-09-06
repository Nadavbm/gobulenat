#!/bin/sh
source env.sh

docker-compose down
docker-compose up -d pgsql

go run api/server/main.go