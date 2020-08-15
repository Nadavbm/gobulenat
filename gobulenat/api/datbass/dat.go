package datbass

import "github.com/nadavbm/gobulenat/gobulenat/pkg/logger"

func InitDB(logger *logger.Logger) {
	logger.Info("could not initialize db")
}

func RunMigrations(logger *logger.Logger) {
	logger.Info("could not run db migrations")
}
