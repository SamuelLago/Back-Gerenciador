package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	New()
}

func New() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", greet)

	http.ListenAndServe(":5555", mux)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
