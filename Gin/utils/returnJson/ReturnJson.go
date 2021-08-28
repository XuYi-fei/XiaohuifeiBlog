package returnJson

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnFailedJson(c *gin.Context, msg returnMsg.ReturnMsg, code returnCode.ReturnCode, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"success": false,
		"msg":     msg,
		"data":    nil,
		"error":   err.Error(),
	})
	return
}

func ReturnSuccessJson(c *gin.Context, msg returnMsg.ReturnMsg, data interface{}, err error) {
	c.JSON(http.StatusOK, gin.H{
		"status":  returnCode.OK,
		"success": true,
		"msg":     msg,
		"data":    data,
		"error":   nil,
	})
}
