package route

import (
	"github.com/Saident/Mini-Project-Averil-Primayuda/controller"
	m "github.com/Saident/Mini-Project-Averil-Primayuda/middleware"
	"github.com/labstack/echo"

	"github.com/Saident/Mini-Project-Averil-Primayuda/constants"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//Non-JWT Route

	//Testing
	e.GET("/users", controller.GetUsersController)
	e.POST("/login", controller.LoginUser)
	e.POST("/loginPerusahaan", controller.LoginPerusahaan)
	e.POST("/register", controller.CreateUserController)
	e.POST("/registerPerusahaan", controller.CreatePerusahaanController)
	e.GET("/jobs", controller.GetJobsController)
	e.POST("/postjobs", controller.PostJobsController)

	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	//JWT Route

	m.LogMiddleware(e)
	return e
}
