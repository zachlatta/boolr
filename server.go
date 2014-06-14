package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/zachlatta/boolr/database"
	"github.com/zachlatta/boolr/handler"
)

func httpLog(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		handler.ServeHTTP(w, r)
		log.Printf("Completed in %s", time.Now().Sub(start).String())
	})
}

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
	r.Handle("/users/{id}/booleans",
		handler.AppHandler(handler.GetUserBooleans)).Methods("GET")

	r.Handle("/booleans",
		handler.AppHandler(handler.CreateBoolean)).Methods("POST")
	r.Handle("/booleans/{id}",
		handler.AppHandler(handler.GetBoolean)).Methods("GET")
	r.Handle("/booleans/{id}/switch",
		handler.AppHandler(handler.SwitchBoolean)).Methods("PUT")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, httpLog(http.DefaultServeMux)))
}
