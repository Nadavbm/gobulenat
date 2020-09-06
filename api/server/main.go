package main

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"

	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/pkg/logger"
)

var db *sql.DB
var tpl = template.Must(template.ParseGlob("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/templates/*html"))

func init() {
	logger := logger.DevLogger()
	dat.InitDB(logger)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homePage)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		errors.Wrap(err, "cannot listen and serve")
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute home template")
	}

	return
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute home template")
	}

	return
}

func profilePage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute home template")
	}

	return
}

func signupPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute home template")
	}

	return
}
