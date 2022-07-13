package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()
	v4 := r.Group("/v4")
	{
		v4.GET("/test", func(c *gin.Context) {
			c.String(200, "You Did it!!!")
		})
	}
	return r
}
