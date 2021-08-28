package usersService

import (
	"Gin/constant"
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dao"
	"Gin/dto/usersDto"
	"Gin/models"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func UserRegisterService(userDto *usersDto.RegisterInfoDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error) {
	var user models.Users

	dao.DB.Where("user_id = ?", userDto.UserId).Find(&user)
	if user.ID != 0 {
		return returnCode.UserIdDuplicate, returnMsg.UserIdDuplicate, errors.New(returnMsg.UserIdDuplicate)
	}

	dao.DB.Where("user_name = ?", userDto.UserName).Find(&user)
	if user.ID != 0 {
		return returnCode.UserNameDuplicate, returnMsg.UserNameDuplicate, errors.New(returnMsg.UserNameDuplicate)
	}

	hash := sha256.New()
	hash.Write([]byte(constant.SALT))
	hash.Write([]byte(userDto.Password))
	passwd := hex.EncodeToString(hash.Sum(nil))
	newUser := models.NewUsers(userDto.UserName, passwd, userDto.UserId, userDto.Question, userDto.Answer)

	dao.DB.Create(&newUser)
	return returnCode.OK, returnMsg.RegisterOK, nil
}
