package routers

import (
	"myapp/middlewares"

	"myapp/controllers"

	"github.com/gin-gonic/gin"
)

func ApiRoute(router *gin.Engine) {

	router.Use(middlewares.AuthMiddleware())

	apiRouter := router.Group("/api")

	// User
	apiRouter.POST("/register", controllers.UserController.Register)
	apiRouter.POST("/login", controllers.UserController.Login)

	// Log Product
	apiRouter.POST("/log-product/increase", controllers.LogProductController.CreateFromIncrease)
	apiRouter.POST("/log-product/decrease", controllers.LogProductController.CreateFromDecrease)

	// Product
	apiRouter.GET("/product", controllers.ProductController.GetAll)
	apiRouter.GET("/product/:id", controllers.ProductController.GetByID)
	apiRouter.POST("/product", controllers.ProductController.Create)
	apiRouter.PUT("/product/:id", controllers.ProductController.Update)
	apiRouter.DELETE("/product/:id", controllers.ProductController.Delete)

	// Product Type
	apiRouter.GET("/product-type", controllers.ProductTypeController.GetAll)
	apiRouter.GET("/product-type/:id", controllers.ProductTypeController.GetByID)
	apiRouter.POST("/product-type", controllers.ProductTypeController.Create)
	apiRouter.PUT("/product-type/:id", controllers.ProductTypeController.Update)
	apiRouter.DELETE("/product-type/:id", controllers.ProductTypeController.Delete)
}
