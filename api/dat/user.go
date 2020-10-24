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

func (u *User) CreateUser(l *logger.Logger) *User {
	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		l.Panic("could not open connection to database")
	}
	defer db.Close()

	query := fmt.Sprintf("INSERT INTO users(first_name, last_name, email, password) VALUES ('%s', '%s', '%s', '%s');", u.FirstName, u.LastName, u.Email, u.Password)
	_, err = db.Exec(query)
	if err != nil {
		l.Info("ERROR:", zap.Error(err))
		l.Info("could not execute in database:", zap.String("query:", query))
	}
	l.Info("execute in database:", zap.String("query:", query))
	return u
}

func GetUserById(l *logger.Logger, id int) (*User, error) {
	u := &User{}

	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		l.Info("could not connect to database", zap.Error(err))
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id,email,first_name,last_name FROM users WHERE id = %d", id)
	rows, err := db.Query(query)
	if err != nil {
		l.Info("could not get id from database", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&u.Id, &u.Email, &u.FirstName, &u.LastName)
		if err != nil {
			l.Info("could not scan users table")
			return nil, err
		}
	}

	return u, nil
}
