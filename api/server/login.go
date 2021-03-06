package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserSession struct {
	UserId        int
	Authenticated bool
}

// store will hold all session data
var store *sessions.CookieStore

// tpl holds all parsed templates
var tpl *template.Template

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	logger := logger.DevLogger()

	u := NewUser()

	// get client session
	session, err := store.Get(r, "bulenat-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
			Email:    email,
			Password: password,
		}

		if email == "" && password == "" {
			http.Redirect(w, r, "/login", 302)
			session.AddFlash("please enter your credentials")
			logger.Info("no email or password provided. enter your credentials please.")
			return
		}

		emailScan := 0
		rows := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1;", email)
		err := rows.Scan(&emailScan)
		if emailScan == 0 || err == sql.ErrNoRows {
			http.Redirect(w, r, "/login", 302)
			session.AddFlash("we can't find your email in the database, please login again")
			logger.Info("email was not found in the database. redirecting to login page.", zap.Error(err))
			return
		}

		query := fmt.Sprintf("SELECT id,email, password FROM users WHERE email = '%s'", email)
		rowss, err := db.Query(query)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
			session.AddFlash("we lost your credentials - please sign up again")
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

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
		if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
			logger.Info("incorrect password", zap.String("for email:", u.Email))
			session.AddFlash("hi", u.FirstName, "the password was incorrect")
			http.Redirect(w, r, "/login", 302)
			return
		}

		url := fmt.Sprintf("/profile/%d", user.Id)

		// Set user as authenticated
		userSession := &UserSession{
			UserId:        user.Id,
			Authenticated: true,
		}

		session.Values["user"] = userSession
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//setSession(email, w)
		http.Redirect(w, r, url, 302)

	}

	err = tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute login html template")
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "bulenat-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = UserSession{}
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", 302)
}
