package server

import (
	"net/http"
	"text/template"

	"github.com/pkg/errors"
)

var tpl = template.Must(template.ParseGlob("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/templates/*html"))

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

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute signup html template")
	}
}
