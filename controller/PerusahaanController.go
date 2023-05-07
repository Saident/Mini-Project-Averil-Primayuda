package controller

import (
	"net/http"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/Saident/Mini-Project-Averil-Primayuda/utils"

	"github.com/labstack/echo"
)

func GetPerusahaansController(c echo.Context) error {
	var perusahaans []model.Perusahaan

	if err := config.DB.Find(&perusahaans).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success get all perusahaans",
		"perusahaans": perusahaans,
	})
}

func GetPerusahaanController(c echo.Context) error {
	claims, err := utils.GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var perusahaans model.Perusahaan
		if err := config.DB.First(&perusahaans, perusahaan_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "success get perusahaans by id",
			"perusahaans": perusahaans,
		})
	}
	return echo.ErrForbidden
}

func CreatePerusahaanController(c echo.Context) error {
	perusahaans := model.Perusahaan{}
	c.Bind(&perusahaans)

	if err := config.DB.Save(&perusahaans).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create new perusahaans",
		"perusahaans": perusahaans,
	})
}

func UpdatePerusahaanController(c echo.Context) error {
	claims, err := utils.GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var perusahaans model.Perusahaan
		if err := config.DB.First(&perusahaans, perusahaan_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		c.Bind(&perusahaans)

		if err := config.DB.Save(&perusahaans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "success update perusahaan by id",
			"perusahaans": perusahaans,
		})
	}
	return echo.ErrForbidden
}

func DeletePerusahaanController(c echo.Context) error {
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var perusahaan model.Perusahaan
		if err := config.DB.First(&perusahaan, perusahaan_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := config.DB.Delete(&perusahaan).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user by id",
		})
	}
	return echo.ErrForbidden
}
