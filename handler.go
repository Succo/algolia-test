package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (I *Index) queryCounter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]
	fmt.Fprint(w, len(I.getRange(date)))
}

func (I *Index) queryPopularity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprint(w, params)
}
