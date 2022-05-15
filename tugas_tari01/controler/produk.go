package controller

import (
	"net/http"
	"tugas_tari01/config"
	"tugas_tari01/model"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var logs []model.Product

	if err := config.DB.Find(&logs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, logs)
}

func GetBookController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	log_product := model.LogProduk{}

	config.DB.Where("ID = ?", id).First(&log_product)

	if log_product.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	return c.JSON(http.StatusOK, log_product)
}

func CreateBookController(c echo.Context) error {
	DB := config.Connect()
	log_product := model.LogProduk{}
	c.Bind(&log_product)

	if err := config.DB.Save(&log_product).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, log_product)
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")

	config.DB.Delete(&model.Book{}, id)

	return c.JSON(http.StatusOK, nil)
}

func UpdateBookController(c echo.Context) error {
	DB := config.Connect()
	id := c.Param("id")
	product := model.Product{}

	config.DB.Where("ID = ?", id).First(&product)

	if product.ID == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}

	payload := model.Product{}
	c.Bind(&payload)

	product.ID = payload.ID
	product.Code = payload.Code
	product.ProductTypeId = payload.ProductTypeId
	product.Name = payload.Name
	product.Description = payload.Description
	product.Price = payload.Price
	product.Qty = payload.Qty
	config.DB.Save(&product)

	return c.JSON(http.StatusOK, product)
}
