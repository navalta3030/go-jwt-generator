package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go_jwt_generator/utils"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/token", utils.GeneratePayload)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Listen on port 8020...")
	log.Fatal(http.ListenAndServeTLS(":8020", "server.crt", "server.key", r))
}
