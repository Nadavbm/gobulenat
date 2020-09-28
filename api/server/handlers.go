package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var tpl = template.Must(template.ParseGlob("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/templates/*html"))

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger := logger.DevLogger()

	u := NewUser()

	session, _ := store.Get(r, "cookie-name")

	conn := dat.GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		logger.Panic("could not open connection to database")
	}
	defer db.Close()

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")

		u = &User{
			Id:       0,
			Email:    email,
			Password: password,
		}

		if email == "" && password == "" {
			http.Redirect(w, r, "/", 302)
			logger.Info("no email or password provided. enter your credentials please.")
			return
		}

		emailScan := 0
		rows := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1;", email)
		err := rows.Scan(&emailScan)
		if emailScan == 0 || err == sql.ErrNoRows {
			http.Redirect(w, r, "/", 302)
			logger.Info("ERROR: email was not found in the database. redirecting to login page.", zap.Error(err))
			return
		}

		query := fmt.Sprintf("SELECT id,email, password FROM users WHERE email = '%s'", email)
		rowss, err := db.Query(query)
		if err != nil {
			http.Redirect(w, r, "/", 302)
			logger.Info("could not get login credentials from database", zap.Error(err))
			return
		}

		//expiresAt := time.Now().Add(time.Minute * 100000).Unix()
		user := new(User)
		for rowss.Next() {
			err := rowss.Scan(&user.Id, &user.Email, &user.Password)
			if err != nil {
				logger.Info("could not scan users table")
			}
		}

		fmt.Println("from database:", user.Email, user.Password, "from form:", u.Email, u.Password)

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
		if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
			logger.Info("incorrect password", zap.String("for email:", u.Email))
			http.Redirect(w, r, "/", 302)
			return
		}

		url := fmt.Sprintf("/profile/%d", user.Id)
		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Save(r, w)

		setSession(email, w)
		http.Redirect(w, r, url, 302)

	}

	err = tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute login html template")
	}

	//clearSession(w)
	//http.Redirect(w, r, "/", 302)
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
		fmt.Println("bcrypt password:", pass)

		conn := dat.GetDBConnString()
		db, err := sql.Open("postgres", conn)
		if err != nil {
			logger.Panic("could not open connection to database")
		}
		defer db.Close()

		query := fmt.Sprintf("INSERT INTO users(first_name, last_name, email, password) VALUES ('%s', '%s', '%s', '%s');", u.FirstName, u.LastName, u.Email, u.Password)
		_, err = db.Exec(query)
		if err != nil {
			logger.Info("ERROR:", zap.Error(err))
			logger.Info("could not execute in database:", zap.String("query:", query))
		}
		logger.Info("execute in database:", zap.String("query:", query))

		http.Redirect(w, r, "/", 302)
		return
	}

	err := tpl.ExecuteTemplate(w, "signup.html", u)
	if err != nil {
		errors.Wrap(err, "could not execute signup html template")
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute home html template")
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}
