package routes

import (
	"sns-app/controllers"

	"github.com/labstack/echo/v4"
)

func PostRoutes(router *echo.Group) {
	b := router.Group("/posts")

	b.POST("", controllers.GetPosts)
}
