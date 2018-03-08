package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/valyala/tsvreader"
)

func main() {
	file, err := os.Open("hn_logs.tsv")
	if err != nil {
		log.Fatal(err)
	}
	r := tsvreader.New(file)
	var i int

	for r.Next() {
		date := r.DateTime()
		query := r.String()
		if i%1000 == 0 {
			fmt.Printf("date=%s, query=%s\n", date.Format(time.UnixDate), query)
		}
		i++
	}
	if err := r.Error(); err != nil {
		fmt.Printf("unexpected error: %s", err)
	}
}
