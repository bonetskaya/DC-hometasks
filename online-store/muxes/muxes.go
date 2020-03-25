package muxes

import (
	"github.com/gorilla/mux"
	"net/http"
	m "online-store/models"
	u "online-store/utils"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	switch code, items := m.GetAllItems(); code {
	case m.DBError:
		u.Respond(w, u.Message(false, "Internal Server Error"), 500)
	case m.OK:
		u.Respond(w, map[string]interface{}{"data": items}, 200)
	case m.NotFound:
		u.Respond(w, u.Message(false, "Not Found"), 404)
	}
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	switch code, item := m.GetItem(params["id"]); code {
	case m.InvalidData:
		u.Respond(w, u.Message(false, "Invalid request"), 400)
	case m.NotFound:
		u.Respond(w, u.Message(false, "Internal Server Error"), 500)
	case m.OK:
		u.Respond(w, map[string]interface{}{"data": item}, 200)
	}
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	switch code, id := m.CreateItem(r.Body); code {
	case m.InvalidData:
		u.Respond(w, u.Message(false, "Invalid request"), 400)
	case m.DBError:
		u.Respond(w, u.Message(false, "Internal Server Error"), 500)
	case m.OK:
		u.Respond(w, map[string]interface{}{"id": id}, 201)
	}
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	switch code := m.UpdateItem(r.Body, params["id"]); code {
	case m.InvalidData:
		u.Respond(w, u.Message(false, "Invalid request"), 400)
	case m.DBError:
		u.Respond(w, u.Message(false, "Internal Server Error"), 500)
	case m.OK:
		u.Respond(w, u.Message(false, "Success"), 200)
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	switch code := m.UpdateItem(r.Body, params["id"]); code {
	case m.InvalidData:
		u.Respond(w, u.Message(false, "Invalid request"), 400)
	case m.DBError:
		u.Respond(w, u.Message(false, "Internal Server Error"), 500)
	case m.OK:
		u.Respond(w, u.Message(false, "Success"), 200)
	}
}
