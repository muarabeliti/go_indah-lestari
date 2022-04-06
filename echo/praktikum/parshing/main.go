package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("saya indah lestari.")

	e := echo.New()
	e.POST("/", func(ctx echo.Context) error {

		name := ctx.Param("hello")
		message := ctx.Param("indah")

		return ctx.String(http.StatusOK, "Hello saya indah lestari")

	})
	e.Start(":1000")
}
