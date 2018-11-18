package services

import (
	"github.com/yagoazedias/rest-api/common"
	"github.com/yagoazedias/rest-api/model"
	"github.com/yagoazedias/rest-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

type UserView struct{}

func (u *User) Create(user model.User) (*model.User, error) {

	var common common.User

	userRepository := repository.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	newUser, err := userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return common.Shorten(newUser), nil
}

func (u *UserView) GetUsers() ([]model.UserView, error) {

	userRepository := repository.UserView{}

	users, err := userRepository.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}
