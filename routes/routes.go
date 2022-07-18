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
		v4.GET("/items", controllers.GetAll())
		v4.GET("/item/:id", controllers.GetOne())
		v4.POST("/item", controllers.CreateOne())
		v4.PUT("/item/:id", controllers.UpdateOne())
		v4.DELETE("item/:id", controllers.DeleteOne())
	}
	return r
}
