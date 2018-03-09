package main

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/valyala/tsvreader"
)

func TestIndex(t *testing.T) {
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

	// Test all the count queries from the subject
	if !reflect.DeepEqual(getDistinct(I.getRange("2015")), countResponse{Count: 573697}) {
		t.Fatalf("Incorrect distinct count for 2015")
	}
	if !reflect.DeepEqual(getDistinct(I.getRange("2015-08")), countResponse{Count: 573697}) {
		t.Fatalf("Incorrect distinct count for 2015-08")
	}
	if !reflect.DeepEqual(getDistinct(I.getRange("2015-08-03")), countResponse{Count: 198117}) {
		t.Fatalf("Incorrect distinct count for 2015-08-03")
	}
	if !reflect.DeepEqual(getDistinct(I.getRange("2015-08-01 00:04")), countResponse{Count: 617}) {
		t.Fatalf("Incorrect distinct count for 2015-08-01 00:04")
	}

	// Test all the popularity queries from the subject
	if !reflect.DeepEqual(getKMostPopular(I.getRange("2015"), 3), popularityResponse{
		Queries: []queryItem{
			queryItem{
				Query: "http%3A%2F%2Fwww.getsidekick.com%2Fblog%2Fbody-language-advice",
				Count: 6675,
			}, queryItem{
				Query: "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F568045",
				Count: 4652,
			}, queryItem{
				Query: "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F379035%3Fsort%3D1",
				Count: 3100,
			},
		},
	}) {
		t.Fatalf("Incorrect popularity values for 2015")
	}
	if !reflect.DeepEqual(getKMostPopular(I.getRange("2015-08-02"), 5), popularityResponse{
		Queries: []queryItem{
			queryItem{
				Query: "http%3A%2F%2Fwww.getsidekick.com%2Fblog%2Fbody-language-advice",
				Count: 2283,
			}, queryItem{
				Query: "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F568045",
				Count: 1943,
			}, queryItem{
				Query: "http%3A%2F%2Fwebboard.yenta4.com%2Ftopic%2F379035%3Fsort%3D1",
				Count: 1358,
			}, queryItem{
				Query: "http%3A%2F%2Fjamonkey.com%2F50-organizing-ideas-for-every-room-in-your-house%2F",
				Count: 890,
			}, queryItem{
				Query: "http%3A%2F%2Fsharingis.cool%2F1000-musicians-played-foo-fighters-learn-to-fly-and-it-was-epic",
				Count: 701,
			},
		},
	}) {
		t.Fatalf("Incorrect popularity values for 2015-08-02")
	}
}
