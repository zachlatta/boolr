package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zachlatta/boolr/database"
	"github.com/zachlatta/boolr/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := database.Init("postgres",
		os.ExpandEnv("postgres://docker:docker@$DB_1_PORT_5432_TCP_ADDR/docker"))
	if err != nil {
		panic(err)
	}
	defer database.Close()

	r := mux.NewRouter()

	r.Handle("/users", handler.AppHandler(handler.CreateUser)).Methods("POST")
	r.Handle("/users/login", handler.AppHandler(handler.Login)).Methods("POST")
	r.Handle("/users/{id}", handler.AppHandler(handler.GetUser)).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
