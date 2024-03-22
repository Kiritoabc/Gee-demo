package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
	if err != nil {
		return
	}
}
