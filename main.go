package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Method("GET")
	router.HandleFunc("/people/{id}", GetPerson).Method("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Method("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Method("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
