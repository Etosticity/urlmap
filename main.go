package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Etosticity/urlmap/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.RootHandler)
	r.HandleFunc("/{shortenID:[a-z]}", routes.RedirectHandler)
	r.HandleFunc("/api/shorten", routes.ShortenHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
