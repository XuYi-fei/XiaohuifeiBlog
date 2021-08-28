package adminService

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dao"
	"Gin/dto/sentencesDto"
	"Gin/models"
	"errors"
)

func DeleteSentencesTmpAdminService(sentenceDto *sentencesDto.SentenceUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	dao.DB.Where("id = ?", sentenceDto.ID).Delete(&models.SentencesTmp{})
	return returnCode.OK, returnMsg.BlogDeleteSuccess, nil, nil
}

func PassSentencesTmpAdminService(sentenceDto *sentencesDto.SentenceUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	UserId := sentenceDto.UserId
	var user models.Users
	dao.DB.Where("user_id = ?", UserId).Find(&user)
	if user.UserId == "" {
		return returnCode.UserNotFound, returnMsg.UserNotFound, errors.New("用户不存在"), nil
	}

	newSentence := models.NewSentence(sentenceDto.Content, sentenceDto.Author, sentenceDto.UserId, sentenceDto.UserName)
	dao.DB.Where("id = ?", sentenceDto.ID).Delete(&models.SentencesTmp{})
	dao.DB.Create(&newSentence)
	return returnCode.OK, returnMsg.CheckPassed, nil, nil
}
