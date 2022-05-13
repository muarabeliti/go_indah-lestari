package middleware

import (
	"project/config"

	"github.com/labstack/echo/v4"
)

var db = config.DB

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user model.user
	if err := db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
