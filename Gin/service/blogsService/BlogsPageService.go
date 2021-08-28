package blogsService

import (
	"Gin/constant"
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dao"
	"Gin/dto/blogsDto"
	"Gin/models"
	"Gin/vo/blogsVO"
	"github.com/gin-gonic/gin"
)

func GetBlogsPageNumService(c *gin.Context, pageNumDto *blogsDto.BlogPageDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	var pageNumVo blogsVO.BlogsPageNumVO
	var num int64
	if pageNumDto.Mode == constant.AllBlogs {
		dao.DB.Model(&models.Article{}).Count(&num)
	} else if pageNumDto.Mode == constant.AllBlogsTmp {
		dao.DB.Model(&models.ArticleTmp{}).Count(&num)
	} else if pageNumDto.Mode == constant.PersonalBlogs {
		dao.DB.Model(&models.Article{}).Where("user_id = ?", pageNumDto.UserId).Count(&num)
	}
	num = (num / 11) + 1
	pageNumVo.PageNum = int(num)
	return returnCode.OK, returnMsg.OK, nil, pageNumVo
}
