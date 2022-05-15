package routes

import (
	"myproject/contens"
	c "myproject/controller"
	m "myproject/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// user routing
	m.LogMiddleware(e)

	e.POST("/auth/login", c.LoginController)
	e.POST("/users", c.CreateUserController)

	e.GET("/callApi", c.CallApi)

	jwtAuth := e.Group("/restricted")
	jwtAuth.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	jwtAuth.GET("/users", c.GetUsersController)
	jwtAuth.GET("/users/:id", c.GetUserController)
	jwtAuth.DELETE("/users/:id", c.DeleteUserController)
	jwtAuth.PUT("/users/:id", c.UpdateUserController)

	return e
}
