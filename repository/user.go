package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/yagoazedias/rest-api/environment"
	"github.com/yagoazedias/rest-api/model"
)

type User struct{}
type UserView struct{}

func (u User) CreateUser(user model.User) (*model.User, error) {

	env := environment.Postgres{}
	fmt.Printf(env.ConnectionString())
	MustOpen(env.ConnectionString(), false)
	defer db.Close()

	if result := db.Create(&user); result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not create user")
	}

	return &user, nil
}

func (u UserView) GetUsers() ([]model.UserView, error) {

	var users []model.UserView
	env := environment.Postgres{}
	fmt.Printf(env.ConnectionString())
	MustOpen(env.ConnectionString(), false)
	defer db.Close()

	if result := db.Find(&users); result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not find users")
	}

	return users, nil
}

func (u User) FindUserByEmail(email string) (*model.User, error) {

	var user model.User
	env := environment.Postgres{}
	fmt.Printf(env.ConnectionString())
	MustOpen(env.ConnectionString(), false)
	defer db.Close()

	if result := db.First(&user, "email = ?", email); result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not find user")
	}

	return &user, nil
}