package Collection

import (
	"Collection/service"
	"github.com/gin-gonic/gin"
)

var (
	CollectionContr = &CollectionController{}
)

type CollectionController struct {
}

func (p *CollectionController) ReadDb(ctx *gin.Context) {
	service.CollectionServ.SendMqtt()
	return
}

func (p *CollectionController) ReadPgDb() {
	service.CollectionServ.SendMqtt()
	return
}
