package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("req", r.URL)
		fmt.Fprintln(w, "ola")
	})
	log.Println("Test server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	log.Println("fin")
}
