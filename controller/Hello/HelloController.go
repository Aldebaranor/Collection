package Hello

import "github.com/gin-gonic/gin"

func Hello(context *gin.Context) {
	context.JSON(200, gin.H{
		"code": 20,
		"msg":  "hello,world",
	})
}
