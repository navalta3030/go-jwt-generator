package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	utils "./utils"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/token", utils.GeneratePayload).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Listen on port 8020...")
	log.Fatal(http.ListenAndServe(":8020", r))
}
