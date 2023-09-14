package api

import (
	"net/http"
)

func (r router) AuthEndPoints() {
	r.HandleFunc("/register", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("register"))
	})

	r.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("regsiter"))
	})

	r.HandleFunc("/regsiter", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("regsiter"))
	})
}
