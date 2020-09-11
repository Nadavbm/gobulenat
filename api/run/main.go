package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/api/server"
	"github.com/nadavbm/gobulenat/pkg/logger"
)

var db *sql.DB

func init() {
	logger := logger.DevLogger()

	dat.InitDB(logger)
}

func main() {
	l := logger.DevLogger()
	s := server.NewServer(l)
	err := s.Run()
	if err != nil {
		panic(err)
	}
}
