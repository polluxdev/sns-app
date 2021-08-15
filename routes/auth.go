package routes

import (
	"sns-app/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(router *echo.Group) {
	b := router.Group("/auth")

	b.POST("/login", controllers.Login)
}
