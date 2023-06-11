package tool

import (
	"github.com/gin-gonic/gin"
)

// RespErrorWithData 数据错误
func RespErrorWithData(c *gin.Context, date interface{}) {
	c.JSON(200, gin.H{
		"status": false,
		"info":   date,
	})
}

// RespSuccessful 成功反馈
func RespSuccessful(c *gin.Context, info string) {
	c.JSON(200, gin.H{
		"status": true,
		"info":   info + "成功",
	})
}



// RespInternalError 服务器错误
func RespInternalError(c *gin.Context) {
	c.JSON(500, gin.H{
		"status": false,
		"info":   "服务器错误",
	})
}
