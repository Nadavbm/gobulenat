package main

import (
	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/api/server"
	"github.com/nadavbm/gobulenat/pkg/logger"
)

func main() {
	l := logger.DevLogger()

	// initialize db connection and db migrations
	dat.InitDB()

	// initialize cookie store
	//server.InitCS()

	l.Info("starting server on port 8081")
	s := server.NewServer(l, ":8081")

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
