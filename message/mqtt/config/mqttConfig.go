package config

import (
	"EmptyProject/global"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
)

var (
	MqttClient mqtt.Client
)

func InitMqtt() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(global.MqttSetting.Address)
	opts.SetUsername(global.MqttSetting.UserName)
	opts.SetPassword(global.MqttSetting.Password)
	MqttClient = mqtt.NewClient(opts)
	if token := MqttClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("订阅MQTT失败")
	}
}
