package adminService

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dao"
	"Gin/dto/blogsDto"
	"Gin/models"
	"errors"
)

func DeleteBlogsTmpAdminService(blogDto *blogsDto.BlogUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	dao.DB.Where("id = ?", blogDto.Id).Delete(&models.ArticleTmp{})
	return returnCode.OK, returnMsg.BlogDeleteSuccess, nil, nil
}

func PassBlogsTmpAdminService(blogDto *blogsDto.BlogUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	UserId := blogDto.UserId
	var user models.Users
	dao.DB.Where("user_id = ?", UserId).Find(&user)
	if user.UserId == "" {
		return returnCode.UserNotFound, returnMsg.UserNotFound, errors.New("用户不存在"), nil
	}

	newBlog := models.NewArticle(blogDto.BlogId, blogDto.UserId, blogDto.Title, blogDto.Digest, blogDto.Content,
		blogDto.Author, blogDto.Field, blogDto.Classification, blogDto.Tag)
	dao.DB.Create(&newBlog)
	dao.DB.Where("id = ?", blogDto.Id).Delete(&models.ArticleTmp{})
	return returnCode.OK, returnMsg.CheckPassed, nil, nil
}
