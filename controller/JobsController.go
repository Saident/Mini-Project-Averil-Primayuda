package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/constants"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetJobsController(c echo.Context) error {
	var jobs []model.Jobs
	claims, err := GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	if role == "user" || role == "admin" {
		if err := config.DB.Find(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all jobs",
			"jobs":    jobs,
		})
	}
	return echo.ErrBadRequest
}

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
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
