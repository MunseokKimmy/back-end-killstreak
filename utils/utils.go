package utils

import "net/http"

func CheckMethodGet(r *http.Request) bool {
	return r.Method == "GET"
}

func CheckMethodPost(r *http.Request) bool {
	return r.Method == "POST"
}
