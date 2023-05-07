package controller

import (
	"net/http"
	"strconv"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/labstack/echo"
)

func GetJobsController(c echo.Context) error {
	var jobs []model.Jobs
	claims, bool := GetJwtClaims(c)
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

	claims, bool := GetJwtClaims(c)
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
