package routers

import (
	"emptyProject/controller/Hello"
	"emptyProject/controller/Mqtt"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	HelloRouter := routers.Group("/hello")
	{
		HelloRouter.GET("/hello", Hello.Hello)
	}
	MqttRouter := routers.Group("/mqtt")
	{
		MqttRouter.GET("/sub", controller.MqttControl.Subscribe)
		MqttRouter.POST("/pub", controller.MqttControl.Produce)
	}

	return routers
}
