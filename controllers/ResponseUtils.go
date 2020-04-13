package controllers

import (
	"encoding/json"
	"net/http"
)

func responseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
