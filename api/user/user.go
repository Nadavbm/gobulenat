package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	json.NewDecoder(r.Body)

}

func Login(w http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOne(user.Email, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func SetCookie(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "somevalue",
		Path:  "/",
	})

}

func ReadCookie(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		errors.Wrap(err, "unable to read cookie")
	}
	fmt.Fprintln(w, "COOKIE:", c)
}
