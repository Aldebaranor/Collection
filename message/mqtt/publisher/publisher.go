package publisher

import (
	"EmptyProject/global"
	"EmptyProject/message/mqtt/config"
)

func Produce(msg string) {
	config.MqttClient.Publish(global.MqttSetting.Topic, global.MqttSetting.QoS, true, msg)
}
