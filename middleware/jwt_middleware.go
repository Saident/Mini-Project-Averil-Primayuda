package middleware

import (
	"github.com/Saident/Mini-Project-Averil-Primayuda/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(name string, email string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["name"] = name
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}