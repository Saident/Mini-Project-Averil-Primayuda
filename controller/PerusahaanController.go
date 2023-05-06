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
		"message":     "success get all perusahaans",
		"perusahaans": perusahaans,
	})
}

// TODO : add get data from JWT, remove id
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
		"message":     "success get perusahaans by id",
		"perusahaans": perusahaans,
	})
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

// TODO : add get data from JWT, remove id
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
		"message":     "success update perusahaan by id",
		"perusahaans": perusahaans,
	})
}

// TODO : add get data from JWT
func PostJobsController(c echo.Context) error {
	jobs := model.Jobs{}

	claims, err := GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaanId := claims["id"].(float64)

	if role == "perusahaan" {
		jobs.PerusahaanID = int(perusahaanId)
		c.Bind(&jobs)
		if err := config.DB.Save(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new jobs",
			"jobs":    jobs,
		})
	}
	return echo.ErrForbidden
}

// TODO : add get data from JWT, remove perusahaan_id
func GetJobByPerusahaanController(c echo.Context) error {
	perusahaan_id, err := strconv.Atoi(c.Param("perusahaan_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var jobs []model.Jobs
	if err := config.DB.Find(&jobs, perusahaan_id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get jobs by perusahaan id",
		"jobs":    jobs,
	})
}

// TODO : add get data from JWT
func UpdateJobByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("JobId"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	perusahaan_id, err := strconv.Atoi(c.Param("PerusahaanId"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var jobs model.Jobs

	if err := config.DB.First(&jobs, id, perusahaan_id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	c.Bind(&jobs)

	if err := config.DB.Save(&jobs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update job by id",
		"job":     jobs,
	})
}

func GetAllLamaranByPerusahaanController(c echo.Context) error {
	return config.DB.Error
}

func GetLamaranByIdController(c echo.Context) error {
	return config.DB.Error
}

func ValidateLamaranController(c echo.Context) error {
	return config.DB.Error
}

// TODO : add get user lampiran
