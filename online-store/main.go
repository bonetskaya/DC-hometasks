package main

import (
	"./item"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	items   map[int]item.Item
	mu      sync.Mutex
	counter int
)

func getVals() []item.Item {
	mu.Lock()
	vals := make([]item.Item, len(items))
	i := 0
	for _, v := range items {
		vals[i] = v
		i++
	}
	mu.Unlock()
	return vals
}

func getVal(id int) (item.Item, bool) {
	it, ok := items[id]
	return it, ok
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getVals())
}

func header400(w http.ResponseWriter) {
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Write([]byte("invalid request\n"))
	w.WriteHeader(400)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		header400(w)
		return
	}
	if it, ok := getVal(id); ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(it)
		return
	}
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.Write([]byte("Item not found\n"))
	w.WriteHeader(404)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var it item.Item
	err := json.NewDecoder(r.Body).Decode(&it)
	if err != nil {
		header400(w)
		return
	}
	if it.Title == "" || it.Category == "" {
		header400(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	it.ID = counter
	counter++
	mu.Unlock()
	json.NewEncoder(w).Encode(it)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		header400(w)
		return
	}
	mu.Lock()
	if it, ok := getVal(id); ok {
		err := json.NewDecoder(r.Body).Decode(&it)
		if err != nil {
			header400(w)
			mu.Unlock()
			return
		}
		it.ID = id
		items[id] = it
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(it)
		return
	}
	mu.Lock()
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.Write([]byte("Item not found\n"))
	w.WriteHeader(404)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		header400(w)
		return
	}
	mu.Lock()
	if it, ok := getVal(id); ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(it)
		delete(items, id)
		mu.Unlock()
		return
	}
	mu.Unlock()
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.Write([]byte("Item not found\n"))
	w.WriteHeader(404)
}

func main() {
	r := mux.NewRouter()
	items = map[int]item.Item{
		1: item.Item{ID: 1, Title: "Gin", Category: "Drink"},
		2: item.Item{ID: 2, Title: "Milk", Category: "Drink"},
	}
	counter = 3
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
