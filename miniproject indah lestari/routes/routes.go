package routes

import (
	"project/config"
	"project/controller"

	// our_middleware "project/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetUsersController)
	e.GET("/users/:id", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.POST("/login", controller.LoginController)

	// jwt grup
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(config.JwtSecret)))
	r.GET("/user/:id", controller.GetUserDetailController)

	// eAuth := e.Group("")
	// eAuth.Use(middleware.BasicAuth(our_middleware.BasicAuthDB))

	// e.Pre(middleware.RemoveTrailingSlash())
	// our_middleware.LogMiddleware(e)

	// auth := e.Group("auth")
	// auth.POST("/login",controller.LoginController)

	// jwtAuthGroup := e.Group("jwt_auth")
	// jwtAuthGroup.Use(middleware.JWT([]byte(config.JwtSecret)))
	// // jwtAuthGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// // 	SigningMethod: jwt.SigningMethodES256.Name,
	// // 	SigningKeys: []byte(config.JwtSecret),
	// // 	ErrorHandlerWithContext: func(err error, c echo.Context) error {
	// // 		return c.JSON(http.StatusForbidden,map[string]interface{}{
	// // 			"status" : "error login",
	// // 			"message" : e ,
	// // 		})

	// // 	},
	// // }))

	// jwtAuthGroup.GET("/produk",controller.GetBooksController)
	// basicAuthGroup := e.Group("basic_auth")
	// basicAuthGroup.Use(middleware.BasicAuth(our_middleware.AuthLogin))
	// basicAuthGroup.GET("/produk",controller.GetBooksController)

	// e.GET("",controller.GetBooksController,middleware.BasicAuth(our_middleware.AuthLogin))
	return e
}
