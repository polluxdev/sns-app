package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "메가야 사랑해요!")
	})

	v := e.Group("/api/v1")

	AuthRoutes(v)
	PostRoutes(v)

	return e
}
