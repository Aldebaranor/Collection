package main

import (
	"emptyProject/global"
	"emptyProject/mapper/mysql"
	"emptyProject/mapper/postgres"
	"emptyProject/routers"
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

}

func main() {
	var err error
	if strings.Compare(global.DataSourceSetting.Source, "mysql") == 0 {
		err = mysql.ConnectDB()
	} else {
		err = postgres.ConnectDB()
	}
	if err != nil {
		log.Printf("%+v", err)
	}
	defer mysql.Close()
	f, _ := os.Create("./logs/logs.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.Routers()
	r.Run(":" + global.ServerSetting.HttpPort)
}
