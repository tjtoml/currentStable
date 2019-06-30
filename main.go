package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type hook struct {
	Type   string `json:"type"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Active string `json:"active"`
}

type payload struct {
	Zen    string `json:"zen"`
	HookId int    `json:"hook_id"`
	Hook   hook   `json:"hook"`
}

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, r.RequestURI)
}

func githubEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	load := payload{}
	err := json.NewDecoder(r.Body).Decode(&load)
	if err != nil {
		panic(err)
	}

	loadJson, err := json.Marshal(load)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(loadJson)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/api/v1/github/endpoint", githubEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":80", r))
}
