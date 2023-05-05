package controller


import (
	"net/http"

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