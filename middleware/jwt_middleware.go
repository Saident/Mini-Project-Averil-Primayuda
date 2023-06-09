package middleware

import (
	"time"

	"github.com/Saident/Mini-Project-Averil-Primayuda/constants"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(name string, email string, role string, roleID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["nama"] = name
	claims["email"] = email
	claims["role"] = role
	claims["id"] = roleID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
