package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func presonHandler(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "Alice", Age: 30}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Println("encode failed:", err)
	}
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Println("encode failed:", err)
	}
}

func main() {
	p := Person{
		Name: "Alice",
		Age:  30,
	}

	data, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	input := `{
		"name" : "Bob",
		"age" : 25
	}`

	var q Person

	err = json.Unmarshal([]byte(input), &q)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /person", presonHandler)
	http.HandleFunc("POST /person", createPerson)

	fmt.Println(q.Name, q.Age)
	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
