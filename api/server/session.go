package server

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

/*
func HasCookie(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("bulenat-cookie")
	if err != nil {
		return false
	}
	cs := SetCookieStore()
	session, err := cs.Get(r, "bulenat-session")
	if err != nil {
		http.Error(w, "", http.Internal)
		return false
	}

	s, ok := ds[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	return true
}
*/
func SetCookieStore() *sessions.CookieStore {
	cs := sessions.NewCookieStore([]byte("bulenat-session"))
	cs.Options.MaxAge = 180 // 3 minutes
	cs.Options.HttpOnly = true
	cs.Options.Secure = false
	return cs
}
