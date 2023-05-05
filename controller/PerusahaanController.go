package controller

import (
	"net/http"
	"strconv"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/labstack/echo"
)

func GetPerusahaansController(c echo.Context) error {
	var perusahaans []model.Perusahaan

	if err := config.DB.Find(&perusahaans).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all perusahaans",
		"perusahaans":  perusahaans,
	})
}

func GetPerusahaanController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var perusahaans model.Perusahaan
	if err := config.DB.First(&perusahaans, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get perusahaans by id",
		"perusahaans":  perusahaans,
	})
}

func CreatePerusahaanController(c echo.Context) error {
	perusahaans := model.Perusahaan{}
	c.Bind(&perusahaans)

	if err := config.DB.Save(&perusahaans).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new perusahaans",
		"perusahaans":  perusahaans,
	})
}

func UpdatePerusahaanController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var perusahaans model.Perusahaan

	if err := config.DB.First(&perusahaans, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	c.Bind(&perusahaans)

	if err := config.DB.Save(&perusahaans).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update perusahaan by id",
		"perusahaans":   perusahaans,
	})
}