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
		v4.GET("/getItems", controllers.GetAll())
		v4.GET("/getItem/:id", controllers.GetOne())
		v4.POST("/addItem", controllers.CreateOne())
	}
	return r
}
