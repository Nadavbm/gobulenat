package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var tpl = template.Must(template.ParseGlob("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/templates/*html"))

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	logger := logger.DevLogger()

	u := NewUser()

	if r.Method == http.MethodPost {
		// get form values
		email := r.FormValue("username")
		password := r.FormValue("password")
		confirm := r.FormValue("confirm")
		fname := r.FormValue("first_name")
		lname := r.FormValue("last_name")

		if password != confirm {
			logger.Info("passwords does not much. redirecting..")
			http.Redirect(w, r, "/signup", 302)
			return
		}
		fmt.Println("user submit a form:", email, fname, lname, password, confirm)

		u = &User{
			FirstName: fname,
			LastName:  lname,
			Email:     email,
			Password:  password,
		}

		fmt.Println("user struct is:", u)
		conn := dat.GetDBConnString()
		db, err := sql.Open("postgres", conn)
		if err != nil {
			logger.Panic("could not open connection to database")
		}
		defer db.Close()

		query := fmt.Sprintf("INSERT INTO users(first_name, last_name, email, password) VALUES ('%s', '%s', '%s', '%s');", u.FirstName, u.LastName, u.Email, u.Password)
		fmt.Println("query is:", query)
		_, err = db.Exec(query)
		if err != nil {
			fmt.Println("ERROR:", err)
			logger.Info("could not execute in database:", zap.String("query:", query))
		}
		logger.Info("execute in database:", zap.String("query:", query))
	}

	err := tpl.ExecuteTemplate(w, "signup.html", u)
	if err != nil {
		errors.Wrap(err, "could not execute signup html template")
	}

	fmt.Println("user struct: ", u, "\nfirst name: ", u.FirstName, "\tlast name: ", u.LastName)

	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute home html template")
	}
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute about html template")
	}
}
