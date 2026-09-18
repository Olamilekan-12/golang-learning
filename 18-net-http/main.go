package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello from GO!")
}

func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "stranger"
	}
	if len(name) > 20 {
		http.Error(w, "name too long", http.StatusBadRequest)
		return
	}
	_, _ = fmt.Fprintf(w, "Hello, %s!\n", name)
}

func main() {
	http.HandleFunc("GET /{$}", hello)
	http.HandleFunc("GET /greet", greet)

	fmt.Println("server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
