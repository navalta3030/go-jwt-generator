package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go_jwt_generator/constants"
	"go_jwt_generator/controllers"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/token", controllers.GenerateJWTFromPayload).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Listen on port 8020...")
	log.Fatal(http.ListenAndServe(constants.Port, r))
}
