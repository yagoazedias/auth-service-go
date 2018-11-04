package repository

import (
	"github.com/yagoazedias/rest-api/model"
	"github.com/pkg/errors"
	"github.com/yagoazedias/rest-api/environment"
	"fmt"
)

type User struct {}

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