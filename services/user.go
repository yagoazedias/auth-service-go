package services

import (
	"github.com/yagoazedias/rest-api/model"
	"github.com/yagoazedias/rest-api/repository"
)

type User struct {}

func (u *User) Create(user model.User) (*model.User, error) {
	var (
		err error
		newUser *model.User
	)

	userRepository := repository.User{}

	newUser, err = userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}