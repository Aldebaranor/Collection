package routers

import (
	"EmptyProject/controller/Hello"
	"EmptyProject/controller/Mqtt"
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
		MqttRouter.GET("/sub", Mqtt.MqttControl.Subscribe)
		MqttRouter.POST("/pub", Mqtt.MqttControl.Produce)
	}

	return routers
}
