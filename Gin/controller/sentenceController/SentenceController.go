package sentenceController

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dto/sentencesDto"
	"Gin/service/sentencesService"
	"Gin/utils/returnJson"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func SentenceUpload(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	var m sentencesDto.SentenceUploadDto
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := sentencesService.SentenceUploadService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func GetAllSentences(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m sentencesDto.SentencesAllDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}

	code, msg, returnErr, data := sentencesService.GetAllSentencesService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func GetSentencesShow(c *gin.Context) {
	code, msg, returnErr, data := sentencesService.GetSentenceShowService()
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

// 通过审核
func PassSentences(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m sentencesDto.SentenceUpdateDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}

	code, msg, returnErr, data := sentencesService.PassSentencesService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

// 个人管理自己的全部美句
func DeleteSentencesPersonalController(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	var m sentencesDto.SentenceUpdateDto
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := sentencesService.DeleteSentencesPersonalService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}
