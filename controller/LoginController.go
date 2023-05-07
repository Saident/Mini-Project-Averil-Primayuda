package controller

import (
	"strings"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/constants"
	"github.com/Saident/Mini-Project-Averil-Primayuda/middleware"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"

	"net/http"
)

func LoginUser(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := middleware.CreateToken(user.Nama, user.Email, "user", user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loginResponse := model.LoginResponse{Email: user.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Login",
		"user":    loginResponse,
	})
}

func LoginPerusahaan(c echo.Context) error {
	perusahaan := model.Perusahaan{}
	c.Bind(&perusahaan)

	if err := config.DB.Where("email = ? AND password = ?", perusahaan.Email, perusahaan.Password).First(&perusahaan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := middleware.CreateToken(perusahaan.Nama, perusahaan.Email, "perusahaan", perusahaan.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loginResponse := model.LoginResponse{Email: perusahaan.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success Login",
		"perusahaan": loginResponse,
	})
}

func LoginAdmin(c echo.Context) error {
	admin := model.Admin{}
	c.Bind(&admin)

	if err := config.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := middleware.CreateToken("admin", admin.Email, "admin", admin.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loginResponse := model.LoginResponse{Email: admin.Email, Token: token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Login",
		"admin":   loginResponse,
	})
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
		echo.NewHTTPError(http.StatusBadRequest, "Invalid JWT Token")
		return nil, false
	}
}
