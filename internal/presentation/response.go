package presentation

import (
	"encoding/json"
	"net/http"
)

func RenderJson(w http.ResponseWriter, data any, status int) {
	w.Header().Add("Response-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
