package server

import (
	"database/sql"
	"encoding/gob"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/nadavbm/gobulenat/pkg/logger"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Server struct {
	Logger     *logger.Logger
	Mux        *http.ServeMux
	HTTPServer *http.Server
}

func NewServer(l *logger.Logger, listenAddress string) *Server {
	s := &Server{
		Logger: l,
	}

	r, err := CreateApiRouter(l)
	if err != nil {
		l.Error("failed to create mux router")
		panic(err)
	}

	s.Mux = http.NewServeMux()
	s.Mux.Handle("/", r)

	s.HTTPServer = &http.Server{
		Addr: listenAddress,
	}

	return s
}

func (s *Server) Run() error {
	logger := logger.Logger{}

	err := s.HTTPServer.ListenAndServe()
	if err != nil {
		logger.Error("cannot run http server - listen and serve", zap.Error(err))
	}

	return nil
}

func CreateApiRouter(l *logger.Logger) (*mux.Router, error) {
	r := mux.NewRouter()

	r.HandleFunc("/", RootHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST", "GET")
	r.HandleFunc("/logout", LogoutHandler)
	r.HandleFunc("/profile/{id}", ProfileHandler).Methods("GET")
	r.HandleFunc("/signup", SignupHandler).Methods("POST", "GET")

	r.PathPrefix("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/static/").Handler(http.StripPrefix("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/static/", http.FileServer(http.Dir("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/static/"))))

	http.Handle("/", r)
	return r, nil
}

func InitCS(l *logger.Logger) {
	l.Info("initialize cookie store")
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(UserSession{})

	l.Info("initialize html tempalte")
	tpl = template.Must(template.ParseGlob("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/templates/*html"))
}
