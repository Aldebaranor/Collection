package publisher

import (
	"Collection/global"
	"Collection/message/mqtt/config"
)

func Produce(msg interface{}) {
	config.MqttClient.Publish(global.MqttSetting.Topic, global.MqttSetting.QoS, true, msg)
}
