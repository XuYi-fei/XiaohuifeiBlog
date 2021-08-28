package sentencesService

import (
	"Gin/constant"
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dao"
	"Gin/dto/sentencesDto"
	"Gin/models"
	"errors"
)

func SentenceUploadService(sentenceDto *sentencesDto.SentenceUploadDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	UserId := sentenceDto.UserId
	var user models.Users
	dao.DB.Where("user_id = ?", UserId).Find(&user)
	if user.UserId == "" {
		return returnCode.LoginStatusInvalid, returnMsg.LoginStatusInvalid, errors.New(returnMsg.LoginStatusInvalid), nil
	}

	if user.Privilege == 1 {
		newSentence := new(models.Sentence)
		newSentence.Content = sentenceDto.Content
		newSentence.UserId = sentenceDto.UserId
		newSentence.Author = sentenceDto.Author
		newSentence.UserName = user.UserName

		dao.DB.Create(&newSentence)
		return returnCode.OK, returnMsg.UploadSuccessful, nil, nil
	} else {
		newSentence := new(models.SentencesTmp)
		newSentence.Content = sentenceDto.Content
		newSentence.UserId = sentenceDto.UserId
		newSentence.Author = sentenceDto.Author
		newSentence.UserName = user.UserName
		dao.DB.Create(&newSentence)
		return returnCode.OK, returnMsg.UploadWaiting, nil, nil
	}
}

func GetAllSentencesService(sentenceDto *sentencesDto.SentencesAllDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	var sentences []*models.Sentence
	var sentencesTmp []*models.SentencesTmp
	if sentenceDto.Mode == constant.AllBlogsTmp {
		dao.DB.Order("id desc").Offset((sentenceDto.PageNo - 1) * 10).Limit(sentenceDto.PageNo * 10).Find(&sentencesTmp)
	} else {
		if sentenceDto.Mode == constant.AllBlogs {
			if sentenceDto.OrderType == "time" {
				dao.DB.Order("created_time desc").Offset((sentenceDto.PageNo - 1) * 10).Limit(sentenceDto.PageNo * 10).Find(&sentences)
			}
		} else if sentenceDto.Mode == constant.PersonalBlogs {
			if sentenceDto.OrderType == "time" {
				dao.DB.Where("user_id = ?", sentenceDto.UserId).Order("created_time desc").Offset((sentenceDto.PageNo - 1) * 10).Limit(sentenceDto.PageNo * 10).Find(&sentences)
			}
		} else {
			return returnCode.ModeInvalid, returnMsg.ModeInvalid, errors.New("选择的mode无效"), nil
		}
	}

	if len(sentences) == 0 && len(sentencesTmp) == 0 {
		return returnCode.NoData, returnMsg.NoData, errors.New("暂无数据"), nil
	}
	if len(sentences) > 0 {
		return returnCode.OK, returnMsg.OK, nil, sentences
	} else {
		return returnCode.OK, returnMsg.OK, nil, sentencesTmp
	}
}

func GetSentenceShowService() (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	var sentences []*models.SentenceShow
	dao.DB.Find(&sentences)
	if len(sentences) == 0 {
		return returnCode.NoData, returnMsg.NoData, errors.New("暂无数据"), nil
	}
	return returnCode.OK, returnMsg.OK, nil, sentences
}

func PassSentencesService(sentenceDto *sentencesDto.SentenceUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
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

func DeleteSentencesPersonalService(sentenceDto *sentencesDto.SentenceUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	dao.DB.Where("id = ?", sentenceDto.ID).Delete(&models.Sentence{})
	return returnCode.OK, returnMsg.SentenceDeleteSuccess, nil, nil
}
