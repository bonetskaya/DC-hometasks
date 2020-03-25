package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	m "online-store/muxes"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/items", m.GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", m.GetItem).Methods("GET")
	r.HandleFunc("/items", m.CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", m.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", m.DeleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
