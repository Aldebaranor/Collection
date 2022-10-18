package publisher

import (
	"emptyProject/global"
	"emptyProject/message/mqtt/config"
)

func Produce(msg string) {
	config.MqttClient.Publish(global.MqttSetting.Topic, global.MqttSetting.QoS, true, msg)
}
