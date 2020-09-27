package main

import (
	"github.com/nadavbm/gobulenat/api/server"
	"github.com/nadavbm/gobulenat/pkg/logger"
)

func main() {
	l := logger.DevLogger()

	s := server.NewServer(l)
	//s := server.Server{}
	err := s.Run()
	if err != nil {
		panic(err)
	}
}
