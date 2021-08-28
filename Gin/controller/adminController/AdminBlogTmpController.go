package adminController

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dto/blogsDto"
	"Gin/service/adminService"
	"Gin/utils/returnJson"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func PassBlogsTmpAdminController(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m blogsDto.BlogUpdateDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := adminService.PassBlogsTmpAdminService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func DeleteBlogsTmpAdminController(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m blogsDto.BlogUpdateDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := adminService.DeleteBlogsTmpAdminService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}
