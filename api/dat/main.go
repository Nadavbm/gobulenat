package dat

import (
	"database/sql"
	"fmt"

	"github.com/nadavbm/gobulenat/pkg/env"
	"github.com/nadavbm/gobulenat/pkg/logger"
	"go.uber.org/zap"
)

var db *sql.DB

func InitDB(logger *logger.Logger) {

	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		logger.Panic("could not open connection to database", zap.Error(err))
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Panic("could not ping database")
	}
	logger.Info("connected to database: " + env.DatabaseDB + "on host:" + env.DatabaseHost)

	_, err = db.Exec(migration)
	if err != nil {
		logger.Panic("could not run db migrations", zap.Error(err))
	}
	logger.Info("db migration completed")
}

func GetDBConnString() string {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable", env.DatabaseUser, env.DatabasePass, env.DatabaseHost, env.DatabasePort, env.DatabaseDB)
	return conn
}

func getDBConnStr() string {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", env.DatabaseHost, env.DatabasePort, env.DatabaseUser, env.DatabasePass, env.DatabaseDB)
	return conn
}
