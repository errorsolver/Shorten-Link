package router

import (
	"shorten-link/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouters() *gin.Engine {
	r := gin.Default()
	r.GET("/:url", controllers.Get)
	r.POST("/", controllers.Create)
	return r
}
