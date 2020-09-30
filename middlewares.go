package main

import (
	"io"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestLog := "{'client':" + r.RemoteAddr + ",'server':" + r.URL.Host + r.URL.Path + ",'client-agent':" + r.Header.Get("User-Agent") + ",'method':" + r.Method + "}"
		log.Println(requestLog)
		next(w, r)
	}
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			next(w, r)
		} else {
			io.WriteString(w, `{"authenticated":"false"}`)
			return
		}

	}
}
