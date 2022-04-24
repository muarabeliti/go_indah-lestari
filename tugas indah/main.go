package main

import (
	"indah_cantik/config"
	"indah_cantik/middleware"
	"indah_cantik/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	// e.POST("/login",controller.LoginController)
	// e.POST("/users",controller.CreateUserController)

	middleware.LogMiddleware(e)

	// e.POST("/auth/login", controller.LoginController)

	// e.POST("/users", controller.CreateUserController)

	// e.GET("/books", c.GetBooksController)
	// e.GET("/books/:id", c.GetBookController)

	// jwtAuth := e.Group("/restricted")
	// jwtAuth.Use(middleware.JWT([]byte(config.JwtSecret)))

	// jwtAuth.GET("/users", controller.GetUsersController)
	// jwtAuth.GET("/users/:id", controller.GetUserController)
	// jwtAuth.DELETE("/users/:id", controller.DeleteUserController)
	// jwtAuth.PUT("/users/:id", controller.UpdateUserController)

	// jwtAuth.POST("/books", controller.CreateBookController)
	// jwtAuth.DELETE("/books/:id", controller.DeleteBookController)
	// jwtAuth.PUT("/books/:id", controller.UpdateBookController)

	// return e

	// e.POST("/log/masuk",controller.LogMasuk)
	// e.POST("/login",controller.LogKeluar)

	// e.GET("/healthcheck",controller.HealthCheck)

	e.Logger.Fatal(e.Start(":8000"))
}
