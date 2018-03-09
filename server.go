package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func serve(I *Index) {
	router := mux.NewRouter()

	router.HandleFunc("/1/queries/count/{date}", I.queryCounter).Methods("GET")
	router.HandleFunc("/1/queries/popular/{date}", I.queryPopularity).Methods("GET")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	log.Fatal(srv.ListenAndServe())
}
