package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	admin = User{Username: "admin", Password: "admin", Role: "admin"}
	user  = User{Username: "user", Password: "user", Role: "user"}
	users = []User{admin, user}
)

func login(w http.ResponseWriter, r *http.Request) {
	var auth = false
	var role string
	var session Session
	w.Header().Add("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var authUser User
	err := decoder.Decode(&authUser)
	if err != nil {
		fmt.Println(err)
	}
	username := authUser.Username
	password := authUser.Password

	for _, user := range users {
		if username == user.Username && password == user.Password {
			auth = true
			role = user.Role
			break
		}
	}
	if auth == true {
		session.Name = CreateUuid()
		session.Role = role
		session.Auth = "true"
		json, err := json.Marshal(session)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 21600=6 hours
		RedisSetEx(session.Name, string(json), "21600")
		value, err := RedisGet(session.Name)
		if err != nil {
			fmt.Println(err)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "session-name",
			Value:   session.Name,
			Expires: time.Now().Add(21600 * time.Second),
		})
		io.WriteString(w, value)
		return
	} else {
		io.WriteString(w, `{"authenticated":"false"}`)
		return
	}

}

func logout(w http.ResponseWriter, r *http.Request) {

	sessionName, _ := r.Cookie("session-name")
	RedisDelete(sessionName.Value)
	http.SetCookie(w, &http.Cookie{
		Name:    "session-name",
		Value:   "logout",
		Expires: time.Now().Add(21600 * time.Second),
	})
	io.WriteString(w, `{"authenticated":"false"}`)
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	sessionName, _ := r.Cookie("session-name")
	sessionJSON, err := RedisGet(sessionName.Value)
	if err != nil {
		log.Println(err)
		return
	}

	if sessionJSON == "" {
		io.WriteString(w, `{"authenticated":"false"}`)
		return
	} else {
		io.WriteString(w, sessionJSON)
		return

	}

}

func autoLogin(w http.ResponseWriter, r *http.Request) {
	sessionName, err := r.Cookie("session-name")
	if err != nil {
		io.WriteString(w, `{"authenticated":"false"}`)
		return
	}
	sessionJSON, err := RedisGet(sessionName.Value)
	if err != nil {
		io.WriteString(w, `{"authenticated":"false"}`)
		return
	}

	if sessionJSON != "" {
		io.WriteString(w, sessionJSON)
		return
	} else {
		io.WriteString(w, `{"authenticated":"false"}`)
		return
	}

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/index", index).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("POST")
	r.HandleFunc("/auto", autoLogin).Methods("POST")
	fmt.Println("Serving on:8080")
	http.ListenAndServe(":8080", r)
}
