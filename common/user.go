package common

import "github.com/yagoazedias/rest-api/model"

type User struct {}

type Invalidation struct {
	Message string
}

type JwtToken struct {
	Token string `json:"token"`
}

	type Exception struct {
	Message string `json:"message"`
}

func (v User) Validate(user model.User) (bool, Invalidation) {
	if user.Email == "" {
		return false, Invalidation{"Invalidation: email is a required field"}
	} else if user.Name == "" {
		return false, Invalidation{"Invalidation: name is a required field"}
	} else if user.Password == "" {
		return false, Invalidation{"Invalidation: password is a required field"}
	}
	return true, Invalidation{}
}

func (v User) Shorten(user *model.User) *model.User {
	user.Password  = ""
	user.DeletedAt = nil

	return user
}