package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Listening port 5005")
	log.Fatal(http.ListenAndServe(":5005", router))
}
