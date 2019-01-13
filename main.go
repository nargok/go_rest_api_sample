package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	// people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "1", Firstname: "Taro", Lastname: "Yamada", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Jiro", Lastname: "Koyama", Address: &Address{City: "City Y", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Sabro", Lastname: "Tanaka"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	fmt.Println("The server is runnning on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
