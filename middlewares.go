package main

import (
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
