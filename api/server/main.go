package main

import (
	"fmt"

	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/pkg/logger"

	"github.com/nadavbm/gobulenat/pkg/env"
)

func main() {
	l := logger.DevLogger()

	dat.InitDB(l)
	dat.RunMigrations(l)

	fmt.Println(env.DatabaseUser)
	fmt.Println(env.DatabasePort)
}
