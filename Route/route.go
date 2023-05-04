package route

import (
	"github.com/labstack/echo"
	m "github.com/Saident/Mini-Project-Averil-Primayuda/middleware"

	"github.com/Saident/Mini-Project-Averil-Primayuda/constants"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//Non-JWT Route
	
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	//JWT Route

	m.LogMiddleware(e)
	return e
}
