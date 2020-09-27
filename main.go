package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key-4")
	store = sessions.NewCookieStore(key)
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func login(w http.ResponseWriter, r *http.Request) {

	//client'tan session-name i al ve session storedan ilgili session'ı getir.

	session, _ := store.Get(r, "session-name")
	w.Header().Add("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var authUser User
	err := decoder.Decode(&authUser)
	if err != nil {
		fmt.Println(err)
	}
	username := authUser.Username
	password := authUser.Password
	if username == "admin" && password == "admin" {
		session.Values["authenticated"] = true
		session.Save(r, w)
		io.WriteString(w, `{"authenticated":"true"}`)
		return
	} else {
		session.Values["authenticated"] = false
		session.Save(r, w)
		io.WriteString(w, `{"authenticated":"false"}`)
		return
	}

}

func logout(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Options.MaxAge = 0
	session.Save(r, w)
	io.WriteString(w, `{"authenticated":"false"}`)
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	fmt.Println(session.Values)
	if session.Values["authenticated"] == true {
		io.WriteString(w, `{"message":"index"}`)
		fmt.Println(session.Values)
		return
	} else {
		io.WriteString(w, `{"authenticated":"false"}`)
		fmt.Println(session.Values)
		return
	}

}

func main() {
	store.Options = &sessions.Options{
		//Domain:   "localhost",
		Path:     "/",
		MaxAge:   3600 * 6, //6 hour
		HttpOnly: true,
	}

	r := mux.NewRouter()
	r.HandleFunc("/index", index).Methods("POST")
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("POST")
	fmt.Println("Serving on:8080")
	http.ListenAndServe(":8080", r)
	/*http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	fmt.Println("Serving on:8080")
	http.ListenAndServe(":8080", nil)*/
}
