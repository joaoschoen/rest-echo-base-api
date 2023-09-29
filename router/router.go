package router

import (
	"API-ECHO/view"

	"github.com/labstack/echo/v4"
)

func InitRoutes(server *echo.Echo) {

	view.UserRoutes(server)
}
