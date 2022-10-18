package main

import (
	"emptyProject/global"
	"emptyProject/mapper"
	"emptyProject/routers"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func init() {
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

}

func main() {
	err := mapper.ConnectDB()
	if err != nil {
		log.Printf("%+v", err)
	}
	defer mapper.Close()
	f, _ := os.Create("./logs/logs.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(global.ServerSetting.RunMode)
	r := routers.Routers()
	r.Run(":" + global.ServerSetting.HttpPort)
}
