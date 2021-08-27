package routes

import (
	"github.com/labstack/echo"
	"github.com/polluxdev/sns-app/app/interfaces"
)

func PostRoutes(router *echo.Group, controller interfaces.PostController) {
	b := router.Group("/posts")

	b.POST("", controller.Index)
}
