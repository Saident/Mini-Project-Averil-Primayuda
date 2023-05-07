package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/jinzhu/gorm"
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
	claims, bool := GetJwtClaims(c)
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
	claims, bool := GetJwtClaims(c)
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

func PostLamaranController(c echo.Context) error {
	lamarans := model.Lamaran{}
	perusahaan_id, err := strconv.Atoi(c.Param("perusahaan_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	job_id, err := strconv.Atoi(c.Param("job_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {

		if err := config.DB.Where("user_id = ? AND job_id = ?", user_id, job_id).First(&lamarans).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				lamarans.UserID = int(user_id)
				lamarans.PerusahaanID = perusahaan_id
				lamarans.JobID = job_id
				c.Bind(&lamarans)

				if err := config.DB.Save(&lamarans).Error; err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}
				return c.JSON(http.StatusOK, map[string]interface{}{
					"message":  "success create new lamarans",
					"lamarans": lamarans,
				})
			} else {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "You already Applied for the Job")
		}
	}
	return echo.ErrForbidden
}

func GetLamaranStatusController(c echo.Context) error {
	claims, bool := GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {
		var lamarans model.Lamaran
		if err := config.DB.Where("user_id = ?", user_id).Find(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success get all lamarans",
			"lamarans": lamarans,
		})
	}
	return echo.ErrForbidden
}

func PostLampiranController(c echo.Context) error {
	return config.DB.Error
}

func GetLampiranListController(c echo.Context) error {
	return config.DB.Error
}

func GetLampiranByIdController(c echo.Context) error {
	return config.DB.Error
}
