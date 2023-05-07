package controller

import (
	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/middleware"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"

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
