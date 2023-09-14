package api

import "net/http"

type router struct {
	*http.ServeMux
}

func NewRouter() router {
	return router{ServeMux: http.NewServeMux()}
}

func (r router) Prefix(path string) router {
	router := NewRouter()
	var handler http.Handler = router
	r.Handle(path+"/", http.StripPrefix(path, handler))
	return router
}

func (r router) Version(v string) *router {
	router := NewRouter()
	var handler http.Handler = router
	r.Handle(v+"/", http.StripPrefix(v, handler))
	return &router
}
