package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/Saident/Mini-Project-Averil-Primayuda/utils"
	"github.com/jinzhu/gorm"

	"github.com/labstack/echo"
)

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

	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {
		if err := config.DB.Where("user_id = ? AND job_id = ?", user_id, job_id).First(&lamarans).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				lamarans.Lamaran_status = "Belum Divalidasi"
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
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {
		var lamarans []model.Lamaran
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

func GetAllLamaranByPerusahaanController(c echo.Context) error {
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var lamarans []model.Lamaran
		if err := config.DB.Where("perusahaan_id = ?", perusahaan_id).Find(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success get all lamarans",
			"lamarans": lamarans,
		})
	}
	return echo.ErrForbidden
}

func GetLamaranByIdPerusahaanController(c echo.Context) error {
	lamaran_id, err := strconv.Atoi(c.Param("lamaran_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var lamarans model.Lamaran
		var users model.User
		if err := config.DB.Where("perusahaan_id = ? AND id = ?", perusahaan_id, lamaran_id).First(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := config.DB.First(&users, lamarans.UserID).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"lamaran": lamarans,
			"pelamar": users,
		})
	}
	return echo.ErrForbidden
}

func ValidateLamaranController(c echo.Context) error {
	lamaran_id, err := strconv.Atoi(c.Param("lamaran_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := int(claims["id"].(float64))

	if role == "perusahaan" {
		var lamarans model.Lamaran
		if err := config.DB.Where("perusahaan_id = ? AND id = ?", perusahaan_id, lamaran_id).First(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Bind(&lamarans)

		if err := config.DB.Save(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success save lamarans",
			"lamarans": lamarans,
		})
	}
	return echo.ErrForbidden
}
