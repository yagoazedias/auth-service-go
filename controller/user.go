package controller

import (
	"encoding/json"
	"github.com/yagoazedias/rest-api/common"
	"github.com/yagoazedias/rest-api/model"
	"github.com/yagoazedias/rest-api/services"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var (
		user   model.User
		common common.User
	)

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	isValid, invalidation := common.Validate(user)

	if !isValid {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(invalidation)
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

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userService = services.UserView{}

	users, err := userService.GetUsers()

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {}

func DeleteUser(w http.ResponseWriter, r *http.Request) {}
