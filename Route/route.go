package route

import (
	"github.com/labstack/echo"
	m "github.com/Saident/Mini-Project-Averil-Primayuda/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	
	m.LogMiddleware(e)
	return e
}
