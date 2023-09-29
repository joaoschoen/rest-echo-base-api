package view

import (
	"API-ECHO/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(server *echo.Echo) {
	group := server.Group("/user")
	group.GET("/:id", controller.GetUser)
	group.GET("/list", controller.GetUserList)
	group.POST("/", controller.PostUser)
	group.PUT("/:id", controller.PutUser)
	group.DELETE("/:id", controller.DeleteUser)
}
