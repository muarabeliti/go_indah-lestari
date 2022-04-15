package interfaces

import "github.com/gin-gonic/gin"

type LogProductInterface interface {
	CreateFromIncrease(*gin.Context)
	CreateFromDecrease(*gin.Context)
}
