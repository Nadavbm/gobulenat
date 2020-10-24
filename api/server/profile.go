package server

import (
	"net/http"

	"github.com/pkg/errors"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		errors.Wrap(err, "could not execute home html template")
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

}
