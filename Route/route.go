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
	//login
	e.POST("/login/user", controller.LoginUser)
	e.POST("/login/perusahaan", controller.LoginPerusahaan)
	e.POST("/login/admin", controller.LoginAdmin)

	//register
	e.POST("/register", controller.CreateUserController)
	e.POST("/register/perusahaan", controller.CreatePerusahaanController)

	//Testing
	e.GET("/users", controller.GetUsersController)
	e.GET("/perusahaan", controller.GetPerusahaansController)

	//JWT Route
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	//TODO : add /update & /delete

	//users
	eJwt.GET("/jobs", controller.GetJobsController)

	eJwt.POST("/lamaran/:perusahaan_id/:job_id", controller.PostLamaranController)
	eJwt.GET("/lamaran/status", controller.GetLamaranStatusController)

	//perusahaan
	eJwt.POST("/jobs/post", controller.PostJobsController)
	eJwt.GET("/jobs/perusahaan", controller.GetJobByPerusahaanController)

	eJwt.GET("/lamaran/perusahaan", controller.GetAllLamaranByPerusahaanController)
	eJwt.GET("/lamaran/perusahaan/:lamaran_id", controller.GetLamaranByIdController)
	eJwt.POST("lamaran/perusahaan/validate/:lamaran_id", controller.ValidateLamaranController)

	m.LogMiddleware(e)
	return e
}
