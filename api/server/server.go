package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nadavbm/gobulenat/api/dat"
	"github.com/nadavbm/gobulenat/pkg/logger"

	_ "github.com/lib/pq"
)

var db *sql.DB

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

func init() {
	dat.InitDB()
}

func (s *Server) Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/profile/{id}", ProfileHandler).Methods("GET")
	r.HandleFunc("/signup", SignupHandler).Methods("POST", "GET")
	r.HandleFunc("/", LoginHandler).Methods("POST", "GET")
	r.HandleFunc("/logout", LogoutHandler)

	/* examples:
	   r.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	   r.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
	   r.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
	   r.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	   r.HandleFunc("/api/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

	   	s := r.PathPrefix("/auth").Subrouter()
		s.Use(auth.JwtVerify)
		s.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
		s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
		s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
		s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	*/

	r.PathPrefix("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/static/").Handler(http.StripPrefix("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/static/", http.FileServer(http.Dir("/home/rodriguez/go/src/github.com/nadavbm/gobulenat/api/server/static/"))))

	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
	return err
}
