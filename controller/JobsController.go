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

	if err := config.DB.Find(&jobs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all jobs",
		"jobs":  jobs,
	})
}

func GetJobController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var jobs model.Jobs
	if err := config.DB.First(&jobs, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get jobs by id",
		"jobs":  jobs,
	})
}

func CreateJobController(c echo.Context) error {
	jobs := model.Jobs{}
	c.Bind(&jobs)

	if err := config.DB.Save(&jobs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new jobs",
		"jobs":  jobs,
	})
}

func UpdateJobController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	var jobs model.Jobs

	if err := config.DB.First(&jobs, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	c.Bind(&jobs)

	if err := config.DB.Save(&jobs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update job by id",
		"job":   jobs,
	})
}