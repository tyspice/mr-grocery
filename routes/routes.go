package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tyspice/mr-grocery/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	v4 := r.Group("/v4")
	{
		v4.GET("/test", controllers.GetTests())
	}
	return r
}
