package classificationController

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dto/classificationDto"
	"Gin/service/classificationService"
	"Gin/utils/returnJson"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func GetField(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	var m classificationDto.FieldDto
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := classificationService.GetFieldService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func GetClassification(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	var m classificationDto.ClassificationDto
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := classificationService.GetClassificationService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func GetTagBlogs(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	var m classificationDto.TagDto
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := classificationService.GetTagBlogsService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}
