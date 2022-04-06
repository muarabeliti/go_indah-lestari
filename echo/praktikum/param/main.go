package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("hello indah.")

	e := echo.New()
	e.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello saya adalah %s", name)

		return ctx.String(http.StatusOK, data)
	})
	e.Start(":1301")
}
