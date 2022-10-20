package service

import (
	"Collection/utils/file"
	"GinHello/mapper"
	"GinHello/message/mqtt/producer"
	"strings"
)

var (
	CollectionServ = &CollectionService{}
)

type CollectionService struct {
}

func (t *CollectionService) sendMqtt() (err error) {
	//读取offset
	offset := file.ReadOffset()
	//拼接sql
	var build strings.Builder
	build.WriteString("" + offset)
	mapper.SqlSession.Exec(build.String())
	//发送mqtt
	producer.Produce("")
	return
}
