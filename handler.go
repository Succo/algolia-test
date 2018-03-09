package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type countResponse struct {
	Count int `json:count`
}

func (I *Index) queryCounter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]
	resp := countResponse{Count: len(I.getRange(date))}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (I *Index) queryPopularity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprint(w, params)
}
