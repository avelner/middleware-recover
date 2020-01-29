package main

import (
	"fmt"
	"log"
	"net/http"

	"velner.org/middleware/recover"
)

func main() {
	http.Handle("/", recover.Do(http.HandlerFunc(handler)))
	// http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test url (%s)", r.URL.Path[100:])
}
