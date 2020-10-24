package dat

import (
	"database/sql"
	"fmt"

	"github.com/nadavbm/gobulenat/pkg/logger"
	"go.uber.org/zap"
)

type User struct {
	Id        int    `db:"id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Password  string `db:"password"`
}

func CreateUser(l *logger.Logger) *User {
	u := &User{}
	return u
}

func GetUser(l *logger.Logger) *User {
	logger := logger.DevLogger()

	var u User
	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		logger.Info("could not connect to database", zap.Error(err))
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", u.Id)
	rows, err := db.Query(query)
	if err != nil {
		logger.Info("could not get id from database", zap.Error(err))
	}

	for rows.Next() {
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName)
		if err != nil {
			logger.Info("could not scan users table")
		}
	}
	fmt.Println("user struct:", u)
	return &u
}
