package routers

import (
	"Collection/controller/Mqtt"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	routers := gin.Default()
	MqttRouter := routers.Group("/mqtt")
	{
		MqttRouter.GET("/sub", Mqtt.MqttControl.Subscribe)
		MqttRouter.POST("/pub", Mqtt.MqttControl.Produce)
	}

	return routers
}
