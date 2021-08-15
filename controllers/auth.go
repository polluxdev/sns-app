package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("User with email : %s", c.FormValue("email")))
}

// func Signup(c echo.Context) (err error) {
// 	u := new(models.User)

// 	err = c.Bind(u)

// 	if err != nil {
// 		return
// 	}

// 	return c.JSON(http.StatusCreated, u)
// }
