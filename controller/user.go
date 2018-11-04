package controller

import (
	"net/http"
	"github.com/yagoazedias/rest-api/model"
	"encoding/json"
	"fmt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Printf(user.Email)

	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}