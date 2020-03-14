package main

import (
	"encoding/json"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Item struct {
	ID       int    `pg:"id" json:"id"`
	Title    string `pg:"title" json:"title"`
	Category string `pg:"category" json:"category"`
}

var (
	db *pg.DB
)

func header400(w http.ResponseWriter) {
	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Write([]byte("indddvalid request\n"))
	w.WriteHeader(400)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	var items []Item
	err := db.Model(&items).Select()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		header400(w)
		return
	}
	item := Item{ID: id}
	err = db.Select(&item)
	if err != nil {
		w.Header().Set("Content-type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Write([]byte("Item not found"))
		w.WriteHeader(404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		header400(w)
		return
	}
	if item.Title == "" || item.Category == "" {
		header400(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	item.ID = 0
	err = db.Insert(&item)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		header400(w)
		return
	}
	var item Item
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		header400(w)
		return
	}
	item.ID = id
	if item.Category != "" && item.Title != "" {
		_, err = db.Model(&item).Column([]string{"category", "title"}...).WherePK().Update()
		if err != nil {
			panic(err)
		}
	} else if item.Title != "" {
		_, err = db.Model(&item).Column("title").WherePK().Update()
		if err != nil {
			panic(err)
		}
	} else if item.Category != "" {
		_, err = db.Model(&item).Column("category").WherePK().Update()
		if err != nil {
			panic(err)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		header400(w)
		return
	}
	item := Item{ID: id}
	_, err = db.Model(&item).WherePK().Delete()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func main() {
	for i := 0; i < 4; i++ {
		db = pg.Connect(&pg.Options{
			Addr:     "database1:5432",
			User:     "audrey_horne",
			Password: "my_love_to_audrey_is_enough_strong_to_be_a_password",
			Database: "database1",
		})
		time.Sleep(100000000)
	}
	_, err := db.Exec("SELECT 1")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = createSchema()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func createSchema() error {
	for _, model := range []interface{}{(*Item)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
