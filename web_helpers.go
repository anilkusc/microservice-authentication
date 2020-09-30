package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

func Logout(w http.ResponseWriter, r *http.Request) {

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

func Index(w http.ResponseWriter, r *http.Request) {
	sessionName, _ := r.Cookie("session-name")
	sessionJSON, _ := RedisGet(sessionName.Value)
	io.WriteString(w, sessionJSON)
	return
}

func AutoLogin(w http.ResponseWriter, r *http.Request) {
	sessionName, _ := r.Cookie("session-name")
	sessionJSON, _ := RedisGet(sessionName.Value)
	io.WriteString(w, sessionJSON)
	return
}

func ProxyPost(w http.ResponseWriter, r *http.Request) {
	var proxy Proxy
	err := json.NewDecoder(r.Body).Decode(&proxy)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := strings.NewReader(proxy.Data)
	req, err := http.NewRequest("POST", proxy.Destination, data)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
	io.WriteString(w, string(body))
	return
}
