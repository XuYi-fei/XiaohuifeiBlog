package adminController

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dto/sentencesDto"
	"Gin/service/adminService"
	"Gin/utils/returnJson"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// 后台管理全部美句
func DeleteSentencesAdminController(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	var m sentencesDto.SentenceUpdateDto
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := adminService.DeleteSentencesAdminService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}
