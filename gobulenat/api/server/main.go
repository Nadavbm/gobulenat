package main

import (
	"fmt"

	"github.com/nadavbm/gobulenat/gobulenat/pkg/logger"

	"github.com/nadavbm/gobulenat/gobulenat/api/datbass"
	"github.com/nadavbm/gobulenat/gobulenat/pkg/env"
)

func main() {
	l := logger.DevLogger()

	datbass.InitDB(l)
	datbass.RunMigrations(l)

	fmt.Println(env.DatabaseUser)
	fmt.Println(env.DatabasePort)
}
