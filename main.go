package main

import (
	"log"
	"os"

	"github.com/valyala/tsvreader"
)

func main() {
	file, err := os.Open("hn_logs.tsv")
	if err != nil {
		log.Fatal(err)
	}

	r := tsvreader.New(file)
	I := newIndex()
	for r.Next() {
		I.add(r.String(), r.String())
	}
	if err := r.Error(); err != nil {
		log.Printf("Failed to parse tsv: %s", err)
		return
	}

	I.done()

	log.Printf("Done indexing, index size %d\n", I.size)
	serve(I)
}
