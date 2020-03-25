package utils

import (
	"encoding/json"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(code)
}
