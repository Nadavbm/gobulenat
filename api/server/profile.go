package server

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/pkg/logger"
	"github.com/pkg/errors"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	l := logger.DevLogger()

	session, err := store.Get(r, "bulenat-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUserSession(session)

	u, err := dat.GetUserById(l, user.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.ExecuteTemplate(w, "home.html", u)
	if err != nil {
		errors.Wrap(err, "could not execute home html template")
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

}

func getUserSession(s *sessions.Session) UserSession {
	val := s.Values["user"]
	var user = UserSession{}
	user, ok := val.(UserSession)
	if !ok {
		return UserSession{Authenticated: false}
	}
	return user
}
