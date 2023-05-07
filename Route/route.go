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
	e.POST("/uploadlampiran/:lampiran_tipe", controller.PostLampiranController)
	e.GET("/getuserlampiran/:lamaran_id", controller.GetUserLampiranByPerusahaanController)

	//JWT Route
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	//TODO : add /update & /delete

	//users
	eJwt.GET("/jobs", controller.GetJobsController)
	eJwt.GET("/jobs/:job_id", controller.GetJobByIdController)

	eJwt.POST("/lamaran/:perusahaan_id/:job_id", controller.PostLamaranController)
	eJwt.GET("/lamaran/status", controller.GetLamaranStatusController)

	//perusahaan
	eJwt.POST("/perusahaan/jobs/post", controller.PostJobsController)
	eJwt.GET("/perusahaan/jobs", controller.GetJobByPerusahaanController)

	eJwt.GET("/perusahaan/lamaran", controller.GetAllLamaranByPerusahaanController)
	eJwt.GET("/perusahaan/lamaran/:lamaran_id", controller.GetLamaranByIdController)
	eJwt.POST("/perusahaan/lamaran/validate/:lamaran_id", controller.ValidateLamaranController)

	//admins
	eJwt.POST("/admin/jobs/validate/:job_id", controller.ValidateJobsController)

	m.LogMiddleware(e)
	return e
}
