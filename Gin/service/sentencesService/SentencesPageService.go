package sentencesService

import (
	"Gin/constant"
	"Gin/constant/returnCode"
	"Gin/dao"
	"Gin/dto/sentencesDto"
	"Gin/models"
	"Gin/vo/sentencesVO"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSentencesPageNumService(c *gin.Context, sentenceDto *sentencesDto.SentencePageDto) {
	var pageNumVo sentencesVO.SentencesPageNumVO
	var num int64
	if sentenceDto.Mode == constant.AllBlogs {
		dao.DB.Model(&models.Sentence{}).Count(&num)
	} else if sentenceDto.Mode == constant.AllBlogsTmp {
		dao.DB.Model(&models.SentencesTmp{}).Count(&num)
	} else if sentenceDto.Mode == constant.PersonalBlogs {
		dao.DB.Model(&models.Sentence{}).Where("user_id = ?", sentenceDto.UserId).Count(&num)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": returnCode.JsonInvalid,
			"msg":    "Mode设置有误",
			"data":   nil,
		})
		return
	}
	num = (num / 11) + 1
	pageNumVo.PageNum = int(num)
	c.JSON(http.StatusOK, gin.H{
		"status": returnCode.OK,
		"msg":    "查询成功",
		"data":   pageNumVo,
	})
}
