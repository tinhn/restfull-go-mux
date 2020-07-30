package main

import (
	"log"
	"net/http"

	service "./controller/userservices"
	"github.com/gorilla/mux"
)

func init() {
	// mdb.CheckConnection()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", service.AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users/{id}", service.FindUserEndpoint).Methods("GET")
	r.HandleFunc("/users", service.CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users/{id}", service.UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", service.DeleteUserEndPoint).Methods("DELETE")

	if err := http.ListenAndServe(":3333", r); err != nil {
		log.Fatal(err)
	}
}
