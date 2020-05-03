package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"go_jwt_generator/constants"
	"go_jwt_generator/routes"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/token", routes.GenerateJWTFromPayload).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Println("Listen on port 8020...")
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(constants.Port, handlers.CORS()(loggedRouter)))
}
