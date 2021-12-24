package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	s := router.PathPrefix("/v1/todo").Subrouter()

	s.HandleFunc("/add", createTask).Methods("POST")
	s.HandleFunc("/update", updateTask).Methods("PUT")
	s.HandleFunc("/delete", deleteTask).Methods("DELETE")
	s.HandleFunc("/find", getTask).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", s))
}

func createTask(w http.ResponseWriter, r *http.Request) {

}

func updateTask(w http.ResponseWriter, r *http.Request) {

}

func deleteTask(w http.ResponseWriter, r *http.Request) {

}

func getTask(w http.ResponseWriter, r *http.Request) {

}
