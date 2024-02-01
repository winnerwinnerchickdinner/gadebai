package handlers

import (
	"net/http"
)

func getTokenHandler(w http.ResponseWriter, r *http.Request) {
	// title := r.URL.Path[len("/auth/token"):]
	// body := r.FormValue("body")
	w.Write([]byte(`{"num":6.13,"strs":["a","b"]}`))
}

func RegistAll() {
	http.HandleFunc("/auth/token", getTokenHandler)
}
