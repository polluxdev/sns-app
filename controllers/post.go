package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {
	return c.String(http.StatusOK, "Post List")
}
