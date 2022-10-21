package main

import (
	"Collection/controller/Collection"
	"Collection/dao/mysql"
	"Collection/dao/postgres"
	"Collection/global"
	"Collection/message/mqtt/config"
	"Collection/routers"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	config.InitMqtt()
}

func main() {

	var err error
	if strings.Compare(global.DataSourceSetting.Source, "mysql") == 0 {
		err = mysql.ConnectDB()
	} else {
		err = postgres.ConnectPGDB()
	}
	if err != nil {
		log.Printf("%+v", err)
	}
	//golang执行定时任务时，数据库链接不需要添加关闭
	//defer mysql.Close()
	//defer postgres.ClosePG()
	//定时任务
	c := cron.New()
	spec := "*/1 * * * * *"
	c.AddFunc(spec, func() {
		Collection.CollectionContr.ReadPgDb()
		time.Sleep(1 * time.Second)
	})
	go c.Start()
	defer c.Stop()

	f, _ := os.Create("./logs/logs.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.Routers()
	r.Run(":" + global.ServerSetting.HttpPort)
}
