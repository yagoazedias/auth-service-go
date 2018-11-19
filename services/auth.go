package services

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

type Auth struct {}

func (u *Auth) ValidateJwt(tokenString string) (bool, bool) {
	token , _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	_, ok := token.Claims.(jwt.MapClaims)
	return ok, token.Valid
}