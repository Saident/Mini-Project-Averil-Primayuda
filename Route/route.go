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

	//JWT Route
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))

	//users
	eJwt.PUT("/user/update", controller.UpdateUserController)
	eJwt.DELETE("/user/delete", controller.DeleteUserController)

	eJwt.POST("/user/lampiran/post/:lampiran_tipe", controller.PostLampiranController)
	eJwt.GET("/user/lampiran/:lampiran_id", controller.GetUserLampiranByUserController)
	eJwt.GET("/user/lampiran", controller.GetLampiranListController)

	eJwt.GET("/jobs", controller.GetJobsController)
	eJwt.GET("/jobs/:job_id", controller.GetJobByIdController)

	eJwt.POST("/lamaran/:perusahaan_id/:job_id", controller.PostLamaranController)
	eJwt.GET("/lamaran/status", controller.GetLamaranStatusController)

	//perusahaan
	eJwt.PUT("/perusahaan/update", controller.UpdatePerusahaanController)
	eJwt.DELETE("/perusahaan/delete", controller.DeletePerusahaanController)

	eJwt.GET("/perusahaan/jobs", controller.GetJobByPerusahaanController)
	eJwt.POST("/perusahaan/jobs/post", controller.PostJobsController)
	eJwt.PUT("/perusahaan/jobs/update/:job_id", controller.UpdateJobByPerusahaanController)

	eJwt.GET("/perusahaan/lamaran", controller.GetAllLamaranByPerusahaanController)
	eJwt.GET("/perusahaan/lamaran/:lamaran_id", controller.GetLamaranByIdPerusahaanController)
	eJwt.POST("/perusahaan/lamaran/validate/:lamaran_id", controller.ValidateLamaranController)

	e.GET("/perusahaan/lamaran/lampiran/:lamaran_id", controller.GetUserLampiranByPerusahaanController)

	//admins
	eJwt.POST("/admin/jobs/validate/:job_id", controller.ValidateJobsController)

	m.LogMiddleware(e)
	return e
}
