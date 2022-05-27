package utils

import (
	"encoding/json"
	"net/http"
)

func Error405CheckGETMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return true
	}
	return false
}

func Error405CheckPOSTMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return true
	}
	return false
}

//Could be a SQL error
func Error500Check(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		http.Error(w, err.Error(), 500)
		return true
	}
	return false
}

//Bad Request
func Error400Check(err error, w http.ResponseWriter) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}

func DecodeRequest(request any, w http.ResponseWriter, r *http.Request) any {
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return request
}
