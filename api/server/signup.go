package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Token      string `json:"token"`
	AuthMethod string `json:"auth_method"`
}

func NewUser() *User {
	user := &User{}
	return user
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	logger := logger.DevLogger()

	u := NewUser()

	if r.Method == http.MethodPost {
		// get form values
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirm := r.FormValue("confirm")
		fname := r.FormValue("first_name")
		lname := r.FormValue("last_name")

		if password != confirm {
			logger.Info("passwords does not much. redirecting..")
			http.Redirect(w, r, "/signup", 302)
			return
		}

		submitForm := fmt.Sprintf("email: %s,\t Name: %s %s, \nPassword: %s, \tConfim:%s", email, fname, lname, password, confirm)
		logger.Info("form submitted:", zap.String("details - ", submitForm))
		u = &User{
			FirstName: fname,
			LastName:  lname,
			Email:     email,
			Password:  password,
		}

		pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			logger.Info("ERROR:", zap.Error(err))
			json.NewEncoder(w).Encode(err)
		}
		u.Password = string(pass)

		user := &dat.User{
			FirstName: fname,
			LastName:  lname,
			Email:     email,
			Password:  string(pass),
		}

		user.CreateUser(logger)

		http.Redirect(w, r, "/login", 302)
		return
	}

	err := tpl.ExecuteTemplate(w, "signup.html", u)
	if err != nil {
		errors.Wrap(err, "could not execute signup html template")
	}
}
