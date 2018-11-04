package controller

import (
	"net/http"
	"github.com/yagoazedias/rest-api/model"
	"encoding/json"
	"github.com/yagoazedias/rest-api/services"
)

type ErrorResponse struct {
	Message  string
	ErrorMsg string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	userService := services.User{}

	newUser, err := userService.Create(user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(newUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}