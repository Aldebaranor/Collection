package service

import (
	"Collection/dao/postgres"
	"Collection/entiy"
	"Collection/global"
	"Collection/message/mqtt/publisher"
	"Collection/utils/file"
	"encoding/json"
	"strconv"
)

var (
	CollectionServ = &CollectionService{}
)

type CollectionService struct {
}

func (t *CollectionService) SendMqtt() (err error) {
	//读取offset
	offset := file.ReadOffset()
	//数据库表结构
	var comments []entiy.Comment
	//SELECT * FROM TABLE ORDER BY column LIMIT step OFFSET offset
	postgres.PgSqlSession.Order(global.DataSourceSetting.Time).Limit(global.DataSourceSetting.Step).Offset(offset).Find(&comments)
	size := len(comments)
	//修改offset
	newOffset, err := strconv.Atoi(offset)
	newOffset += size
	file.WriteOffset(strconv.Itoa(newOffset))
	//
	msg, _ := json.Marshal(comments)
	//发送mqtt
	publisher.Produce(msg)
	return
}
