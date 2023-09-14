package api

import "net/http"

func (r router) UserEndPoints() {

	r.HandleFunc("/profile", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("user profile"))
	})
}
