package dat

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/nadavbm/gobulenat/pkg/env"
	"github.com/nadavbm/gobulenat/pkg/logger"
)

var db *sql.DB
var err error

func InitDB(logger *logger.Logger) error {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable", env.DatabaseUser, env.DatabasePass, env.DatabaseHost, env.DatabasePort, env.DatabaseDB)
	fmt.Println(conn)
	db, err = sql.Open("postgres", conn)
	if err != nil {
		logger.Panic("could not open connection to database")
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Panic("could not ping database")
		return err
	}
	logger.Info("connected to database: " + env.DatabaseDB + "on host:" + env.DatabaseHost)

	_, err = db.Exec(migration)
	if err != nil {
		logger.Info("could not run db migrations")
		return err
	}
	logger.Info("db migrations completed successfully")
	return nil
}
