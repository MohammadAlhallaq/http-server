package main

import (
	"geteway/internal/api"
	"gitlab.com/0x4149/logz"
	"net/http"
	"time"
)

const local bool = true

var srv *http.Server

func init() {

	if local {
		logz.VerbosMode()
	}
	logz.Run()

}

func main() {

	router := api.NewRouter()

	//router.Prefix("/auth").AuthEndPoints()

	if local {
		srv = &http.Server{
			Addr:         ":6969",
			Handler:      router,
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}
	}

	logz.Info("server is running")
	logz.Fatal(srv.ListenAndServe())
}
