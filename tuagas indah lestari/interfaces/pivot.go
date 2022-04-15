package interfaces

import "github.com/gin-gonic/gin"

type PivotInterface interface {
	Create(*gin.Context)
}