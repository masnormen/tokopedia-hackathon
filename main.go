package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
	return
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health", homePage).Methods("GET")

	_ = http.ListenAndServe(":9090", r)
}
