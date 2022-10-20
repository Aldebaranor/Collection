package main

import (
	"Collection/dao/mysql"
	"Collection/dao/postgres"
	"Collection/global"
	"Collection/message/mqtt/config"
	"Collection/routers"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strings"
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
	defer mysql.Close()
	defer postgres.ClosePG()
	//Collection.CollectionContr.ReadDb()

	f, _ := os.Create("./logs/logs.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.Routers()
	r.Run(":" + global.ServerSetting.HttpPort)
}
