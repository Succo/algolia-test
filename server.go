package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func serve() {
	router := mux.NewRouter()

	router.HandleFunc("/1/queries/count/{date}", queryCounter).Methods("GET")
	router.HandleFunc("/1/queries/popular/{date}", queryPopularity).Methods("GET")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	log.Fatal(srv.ListenAndServe())
}

func queryCounter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprint(w, params)
}

func queryPopularity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprint(w, params)
}
