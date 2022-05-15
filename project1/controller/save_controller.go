package controllers

import (
	"myproject/config"
   
	"myproject/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SaveResep(c echo.Context) error {
	DB := config.Connect()
	saveResep := models.SaveResep{}
	c.Bind(&saveResep)

	if err := DB.Save(&saveResep).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    saveResep,
	})
}
