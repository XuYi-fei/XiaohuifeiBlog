package users

import (
	"Gin/authUtils"
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dto/usersDto"
	"Gin/service/usersService"
	"Gin/utils/returnJson"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	//json := make(map[string]interface{})
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m usersDto.LoginInfoDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}

	usersService.Login(c, &m)
}

func LoginStatus(c *gin.Context) {
	claims := c.MustGet("claims").(*authUtils.UserClaims)
	isPass, user, err := usersService.LoginStatus(claims)
	if !isPass {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "自动登录验证通过",
		"data":   user,
	})
}

func UserRegister(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m usersDto.RegisterInfoDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}

	code, msg, returnErr := usersService.UserRegisterService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, nil, returnErr)
	}
}
