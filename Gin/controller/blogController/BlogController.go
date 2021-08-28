package blogController

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dto/blogsDto"
	"Gin/service/blogsService"
	"Gin/utils/returnJson"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func BlogUpload(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m blogsDto.ArticlesDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}

	code, msg, returnErr, data := blogsService.BlogsUploadService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func GetLatestBlogs(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m blogsDto.BlogsAllDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}

	code, msg, returnErr, data := blogsService.GetLatestBlogsService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func GetBlogByBlogId(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m blogsDto.BlogByIdDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}
	code, msg, returnErr, data := blogsService.GetBlogByBlogIDService(m.BlogId)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func UpdateBlogById(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m blogsDto.ArticleDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}

	code, msg, returnErr, data := blogsService.UpdateBlogByIdService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func DeleteBlogByTitle(c *gin.Context) {
	b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
	// 定义map或结构体
	var m blogsDto.ArticleDto
	//var m map[string]interface{}
	// 反序列化
	err := json.Unmarshal(b, &m)
	if err != nil {
		returnJson.ReturnFailedJson(c, returnMsg.JsonInvalid, returnCode.JsonInvalid, err)
		return
	}

	code, msg, returnErr, data := blogsService.DeleteBlogByTitleService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func PassBlogs(c *gin.Context) {
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
	code, msg, returnErr, data := blogsService.PassBlogsService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func DeleteBlogs(c *gin.Context) {
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
	code, msg, returnErr, data := blogsService.DeleteBlogsService(&m)
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}

func GetBlogsInfo(c *gin.Context) {
	code, msg, returnErr, data := blogsService.GetBlogsInfoService()
	if returnErr != nil {
		returnJson.ReturnFailedJson(c, msg, code, returnErr)
	} else {
		returnJson.ReturnSuccessJson(c, msg, data, returnErr)
	}
}
