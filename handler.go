package main

import (
	"encoding/json"
	"net/http"
	"sort"
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
	distinct := make(map[string]bool)
	for _, q := range queries {
		distinct[q] = true
	}

	return countResponse{Count: len(distinct)}
}

// queryPopularitty is the handler for request of most popular queries
func (I *Index) queryPopularity(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date := params["date"]
	size, err := strconv.Atoi(r.FormValue("size"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
// those k items are sorted using sorting functions from the standard library
func getKMostPopular(queries []string, size int) popularityResponse {
	freqs := make(map[string]int)
	for _, q := range queries {
		freqs[q]++
	}

	counted := make([]queryItem, len(freqs))
	var i int
	for q, freq := range freqs {
		counted[i] = queryItem{q, freq}
		i++
	}

	size = min(size, len(counted))

	quickSelect(queryItemSlice(counted), size)
	sort.Sort(queryItemSlice(counted[:size]))

	return popularityResponse{
		Queries: counted[:size],
	}
}

// queryItemSlice implements the sort.Interface to be able to use quickSelect
type queryItemSlice []queryItem

func (q queryItemSlice) Len() int           { return len(q) }
func (q queryItemSlice) Less(i, j int) bool { return q[i].Count > q[j].Count } // Sign inverted as we sort from biggest to smallest
func (q queryItemSlice) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
