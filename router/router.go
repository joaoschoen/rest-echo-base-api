package router

import (
	"API-ECHO/view"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {

	view.UserRoutes(e)

}
