package main

import (
	"bitbucket.org/bemobidev/discount-server/repository"
	"github.com/gorilla/mux"
	"github.com/yagoazedias/rest-api/controller"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/user/", controller.CreateUser).Methods("POST")
	router.HandleFunc("/user/", controller.GetUsers).Methods("GET")
	router.HandleFunc("/login/", controller.Login).Methods("POST")
	router.HandleFunc("/user/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

	defer repository.Close()
}
