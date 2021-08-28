package adminService

import (
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dao"
	"Gin/dto/sentencesDto"
	"Gin/models"
)

func DeleteSentencesAdminService(sentenceDto *sentencesDto.SentenceUpdateDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	dao.DB.Where("id = ?", sentenceDto.ID).Delete(&models.Sentence{})
	return returnCode.OK, returnMsg.SentenceDeleteSuccess, nil, nil
}
