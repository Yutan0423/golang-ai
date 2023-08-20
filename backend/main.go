package main

import (
	"backend/route"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ai", route.AIHandler).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
