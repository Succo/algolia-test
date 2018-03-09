package main

import "sort"

type Index struct {
	timestamps []string
	queries    []string
	size       int
}

func newIndex() *Index {
	return &Index{
		timestamps: make([]string, 0),
		queries:    make([]string, 0),
	}
}

func (I *Index) add(t string, q string) {
	I.timestamps = append(I.timestamps, t)
	I.queries = append(I.queries, q)
}

func (I *Index) done() {
	I.size = len(I.queries)
	sort.Sort(I)
}

func (I *Index) getRange(date string) []string {
	length := len(date)
	min := sort.Search(I.size, func(i int) bool {
		if I.timestamps[i][:length] >= date {
			return true
		}
		return false
	})
	max := sort.Search(I.size, func(i int) bool {
		if I.timestamps[i][:length] > date {
			return true
		}
		return false
	})
	return I.queries[min:max]
}

// Implements the sort interface
func (I *Index) Len() int { return I.size }
func (I *Index) Swap(i, j int) {
	I.timestamps[i], I.timestamps[j] = I.timestamps[j], I.timestamps[i]
	I.queries[i], I.queries[j] = I.queries[j], I.queries[i]
}
func (I *Index) Less(i, j int) bool { return I.timestamps[i] < I.timestamps[j] }
