package services

import (
	"github.com/yagoazedias/rest-api/common"
	"github.com/yagoazedias/rest-api/model"
	"github.com/yagoazedias/rest-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type User struct{}

func (u *User) Create(user model.User) (*model.User, error) {

	var (
		err     error
		newUser *model.User
		common  common.User
	)

	userRepository := repository.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	newUser, err = userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return common.Shorten(newUser), nil
}

func (u *User) GetUsers() ([]model.User, error) {

	userRepository := repository.User{}

	users, err := userRepository.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}
