package api

import "net/http"

type RouterInterface interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	Prefix(path string)
}

type myRouter struct {
	*http.ServeMux
}

func NewRouter() RouterInterface {
	return &myRouter{http.NewServeMux()}
}

func (r *myRouter) Prefix(path string) {
	var handler http.Handler = r.ServeMux
	r.Handle(path+"/", http.StripPrefix(path, handler))
}
