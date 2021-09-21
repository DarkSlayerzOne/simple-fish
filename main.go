package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/james.bondoc.dev/simple-fish-api/app/data"
)

func getFishes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fishes := data.GetFishes()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fishes)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/sea-store/v1/fish", getFishes).Methods("GET")
	log.Fatal(http.ListenAndServe(":8090", r))
}
