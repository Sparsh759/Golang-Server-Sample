package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", ping)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	message := "Success"
	fmt.Fprintf(w, message)
}
