package view

import (
	"API-ECHO/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	g := e.Group("/user")
	g.GET("/:id", controller.GetUser)
	g.POST("/", controller.PostUser)
	g.PUT("/:id", controller.PutUser)
	g.DELETE("/:id", controller.DeleteUser)
}
