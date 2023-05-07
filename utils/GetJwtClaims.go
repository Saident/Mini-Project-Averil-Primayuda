package utils

import (
	"net/http"
	"strings"

	"github.com/Saident/Mini-Project-Averil-Primayuda/constants"
	"github.com/labstack/echo"

	"github.com/dgrijalva/jwt-go"
)

func GetJwtClaims(c echo.Context) (jwt.MapClaims, bool) {
	authHeader := c.Request().Header.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		return nil, false
	}
	jwtToken := authHeaderParts[1]

	hmacSecretString := jwtToken
	hmacSecret := []byte(constants.SECRET_JWT)
	token, err := jwt.Parse(hmacSecretString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		echo.NewHTTPError(http.StatusBadRequest, "Invalid JWT Token")
		return nil, false
	}
}
