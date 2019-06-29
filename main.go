package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, r.RequestURI)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
