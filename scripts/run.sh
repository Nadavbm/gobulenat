#!/bin/sh
source env.sh

docker-compose down
docker-compose up -d pgsql

sleep 10

go run api/run/main.go