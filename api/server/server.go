package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadavbm/gobulenat/pkg/logger"
)

type Server struct {
	Logger     *logger.Logger
	Mux        *http.ServeMux
	HTTPServer *http.Server
}

func NewServer(logger *logger.Logger) *Server {
	server := &Server{
		Logger: logger,
	}

	server.HTTPServer = &http.Server{}
	return server
}

func (s *Server) Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/", HomePage)
	r.HandleFunc("/signup", SignupHandler)
	//r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/about", AboutPage)
	//r.HandleFunc("/logout", LogoutHandler)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	return err
}
