package main

import (
	"net/http"
	"time"

	"phuket/logging"
	"phuket/server"

	"github.com/gorilla/mux"
)

func main() {
	logging.LogInfo("Starting server")

	r := mux.NewRouter()
	r.HandleFunc("/souvenir/{id}", server.SouvenirHandler)

	logging.LogInfo("Handlers set up")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logging.LogInfo("Server created")
	logging.LogInfo("Listening on " + srv.Addr)

	logging.LogFatal(srv.ListenAndServe())
}
