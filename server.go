package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zachlatta/boolr/database"
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

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
