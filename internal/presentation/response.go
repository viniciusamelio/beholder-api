package presentation

import (
	"beholder/internal/exception"
	"encoding/json"
	"net/http"
)

func RenderJson(w http.ResponseWriter, data any, status int) {
	w.Header().Add("Response-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func HandleException(w http.ResponseWriter, err exception.Exception) {
	RenderJson(w, map[string]interface{}{
		"message": err.Message(),
	}, 400)
}
