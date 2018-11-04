package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/yagoazedias/rest-api/controller"
	"bitbucket.org/bemobidev/discount-server/repository"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/user/", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

	defer repository.Close()
}