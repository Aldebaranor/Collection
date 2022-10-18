package routers

import (
	"emptyProject/controller/Hello"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	HelloRouter := routers.Group("/hello")
	{
		HelloRouter.GET("/hello", Hello.Hello)
	}

	return routers
}
