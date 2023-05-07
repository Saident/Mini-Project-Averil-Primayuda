package controller

import (
	"net/http"
	"strconv"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/Saident/Mini-Project-Averil-Primayuda/utils"

	"github.com/labstack/echo"
)

func GetJobsController(c echo.Context) error {
	var jobs []model.Jobs
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	if role == "user" || role == "admin" {
		if err := config.DB.Where("status = ?", "Tervalidasi").Find(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all jobs",
			"jobs":    jobs,
		})
	}
	return echo.ErrBadRequest
}

func GetJobByIdController(c echo.Context) error {
	var jobs []model.Jobs

	job_id, err := strconv.Atoi(c.Param("job_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	if role == "user" || role == "admin" {
		if err := config.DB.Where("status = ?", "Tervalidasi").First(&jobs, job_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get job by id",
			"job":     jobs,
		})
	}
	return echo.ErrBadRequest
}

func PostJobsController(c echo.Context) error {
	jobs := model.Jobs{}

	claims, err := utils.GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		jobs.PerusahaanID = int(perusahaan_id)
		jobs.Status = "Belum Divalidasi"
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

func GetJobByPerusahaanController(c echo.Context) error {
	claims, err := utils.GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	var jobs []model.Jobs
	if role == "perusahaan" {
		if err := config.DB.Where("perusahaan_id = ?", perusahaan_id).Find(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get jobs by perusahaan id",
			"jobs":    jobs,
		})
	}
	return echo.ErrForbidden
}

func UpdateJobByPerusahaanController(c echo.Context) error {
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	job_id, err := strconv.Atoi(c.Param("job_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	if role == "perusahaan" {
		var jobs model.Jobs
		if err := config.DB.Where("job_id = ? AND perusahaan_id = ?", job_id, perusahaan_id).First(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
	return echo.ErrForbidden
}
