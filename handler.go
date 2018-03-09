package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type countResponse struct {
	Count int `json:count`
}

type queryItem struct {
	Query string `json:query`
	Count int    `json:count`
}

type popularityResponse struct {
	Queries []queryItem `json:queries`
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
	date := params["date"]
	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	queries := I.getRange(date)
	resp := getKMostPopular(queries, size)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getKMostPopular(queries []string, size int) popularityResponse {
	m := make(map[string]int)
	for _, q := range queries {
		m[q] = m[q] + 1
	}

	counted := make([]queryItem, len(m))
	var i int
	for q, freq := range m {
		counted[i] = queryItem{q, freq}
		i++
	}

	quickSelect(queryItemSlice(counted), size)

	return popularityResponse{
		Queries: counted[:size],
	}
}

// queryItemSlice implements the sort.Interface to be able to use quickSelect
type queryItemSlice []queryItem

func (q queryItemSlice) Len() int           { return len(q) }
func (q queryItemSlice) Less(i, j int) bool { return q[i].Count > q[j].Count } // Sign inverted as we sort from biggest to smallest
func (q queryItemSlice) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
