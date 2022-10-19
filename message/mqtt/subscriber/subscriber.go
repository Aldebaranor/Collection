package subscriber

import (
	"Collection/global"
	"Collection/message/mqtt/config"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Subscribe() {
	config.MqttClient.Subscribe(global.MqttSetting.Topic, global.MqttSetting.QoS, subCallBackFunc)
}

func subCallBackFunc(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("订阅：当前的话题是 [%s]；信息是[%s] \n", msg.Topic(), string(msg.Payload()))
	fmt.Println(string(msg.Payload()))
}
