package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("saya indah lestari.")

	e := echo.New()
	e.Any("/", func(c echo.Context) error {

		return c.JSON(http.StatusOK, "saya indah lestari")
	})
	e.Start(":1000")
}
