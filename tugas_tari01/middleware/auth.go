package middleware

import (
	"github.com/labstack/echo/v4"
)

func AuthLogin(email, password string, c echo.Context) (bool, error) {
	if email == "user" && password == "user" {
		return true, nil
	}
	return false, nil
}

func BasicAuth(email, password string, c echo.Context) (bool, error) {
	if email == "user" && password == "user" {
		return true, nil
	}
	return false, nil
}
