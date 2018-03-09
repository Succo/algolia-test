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

// queryCounter is the handler for request of distinct query count
func (I *Index) queryCounter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]
	resp := getDistinct(I.getRange(date))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getDistinct counts the number of distinct values in a list of queries
func getDistinct(queries []string) countResponse {
	m := make(map[string]bool)
	for _, q := range queries {
		m[q] = true
	}

	return countResponse{Count: len(m)}
}

// queryPopularitty is the handler for request of most popular queries
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

// getKMostPopular return the K most popular queries from a list of queries
// first it counts the frequencies of all queries with a map
// then it extract the k biggest frequencies using quick select
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
