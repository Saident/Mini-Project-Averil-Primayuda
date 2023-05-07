package controller

import (
	"net/http"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/Saident/Mini-Project-Averil-Primayuda/utils"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {
		var users model.User
		if err := config.DB.First(&users, user_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get users by id",
			"users":   users,
		})
	}
	return echo.ErrForbidden
}

func CreateUserController(c echo.Context) error {
	users := model.User{}
	c.Bind(&users)

	if err := config.DB.Save(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new users",
		"users":   users,
	})
}

func UpdateUserController(c echo.Context) error {
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {
		var users model.User
		if err := config.DB.First(&users, user_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		c.Bind(&users)

		if err := config.DB.Save(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update user by id",
			"user":    users,
		})
	}
	return echo.ErrForbidden
}

func DeleteUserController(c echo.Context) error {
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {
		var user model.User
		if err := config.DB.First(&user, user_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := config.DB.Delete(&user).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user by id",
		})
	}
	return echo.ErrForbidden
}
