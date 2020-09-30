package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	admin = User{Username: "admin", Password: "admin", Role: "admin"}
	user  = User{Username: "user", Password: "user", Role: "user"}
	users = []User{admin, user}
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/index", Logger(Index)).Methods("POST")
	r.HandleFunc("/login", Logger(Login)).Methods("POST")
	r.HandleFunc("/logout", Logger(Logout)).Methods("POST")
	r.HandleFunc("/auto", Logger(AutoLogin)).Methods("POST")
	r.HandleFunc("/proxy", Logger(ProxyPost)).Methods("POST")
	fmt.Println("Serving on:8080")
	http.ListenAndServe(":8080", r)
}
