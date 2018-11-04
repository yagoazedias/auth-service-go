package repository

import (
	"github.com/yagoazedias/rest-api/model"
	"github.com/pkg/errors"
)

func CreateUser(user model.User) (*model.User, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, errors.Wrap(err, "could not create user")
	}

	return &user, nil
}