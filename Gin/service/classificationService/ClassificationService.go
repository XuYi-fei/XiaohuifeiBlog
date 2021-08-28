package classificationService

import (
	"Gin/constant"
	"Gin/constant/returnCode"
	"Gin/constant/returnMsg"
	"Gin/dao"
	"Gin/dto/classificationDto"
	"Gin/models"
	"Gin/vo/classificationVO"
	"errors"
)

func GetFieldService(dto *classificationDto.FieldDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	var result []classificationVO.FieldVO
	if dto.Mode == constant.AllBlogs {
		dao.DB.Debug().Model(models.Article{}).Select("field as FieldName, count(*) as num").Group("field").Find(&result)
	} else if dto.Mode == constant.PersonalBlogs {
		dao.DB.Model(models.Article{}).Select("field as FieldName, count(*) as num").Where("user_id = ?", dto.UserId).Group("field").Find(&result)
	}
	if len(result) == 0 {
		return returnCode.NoData, returnMsg.NoData, errors.New("暂无数据"), result
	}
	return returnCode.OK, returnMsg.OK, nil, result
}

func GetClassificationService(dto *classificationDto.ClassificationDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	var result []classificationVO.ClassificationFromFieldVO
	if dto.Mode == constant.AllBlogs {
		if dto.Field != "" {
			dao.DB.Debug().Model(models.Article{}).Select("classification as ClassificationName, count(*) as num").Where("field = ?", dto.Field).Group("classification").Find(&result)
		} else {
			dao.DB.Debug().Model(models.Article{}).Select("classification as ClassificationName, count(*) as num").Group("classification").Find(&result)
		}
	} else if dto.Mode == constant.PersonalBlogs {
		dao.DB.Model(models.Article{}).Select("classification as ClassificationName, count(*) as num").Where("field = ? and user_id = ?", dto.Field, dto.UserId).Group("classification").Find(&result)
	}
	if len(result) == 0 {
		return returnCode.NoData, returnMsg.NoData, errors.New("暂无数据"), result
	}
	return returnCode.OK, returnMsg.OK, nil, result
}

func GetTagBlogsService(dto *classificationDto.TagDto) (returnCode.ReturnCode, returnMsg.ReturnMsg, error, interface{}) {
	var result []classificationVO.TagBlogsVO
	if dto.Mode == constant.AllBlogs {
		dao.DB.Model(models.Article{}).Select("tag, count(*) as num").Where("field = ? and classification = ?", dto.Field, dto.Classification).Group("tag").Find(&result)
		for i := 0; i < len(result); i++ {
			dao.DB.Model(models.Article{}).Where("field = ? and classification = ? and tag = ?", dto.Field, dto.Classification, result[i].Tag).Find(&result[i].Blogs)
		}
	} else if dto.Mode == constant.PersonalBlogs {
		dao.DB.Model(models.Article{}).Select("classification as ClassificationName, count(*) as num").Where("field = ? and classification = ? and user_id = ?", dto.Field, dto.Classification, dto.UserId).Group("classification").Find(&result)
		for i := 0; i < len(result); i++ {
			dao.DB.Model(models.Article{}).Where("field = ? and classification = ? tag = ? and user_id = ?", dto.Field, dto.Classification, result[i].Tag, dto.UserId).Find(&result[i].Blogs)
		}
	}
	if len(result) == 0 {
		return returnCode.NoData, returnMsg.NoData, errors.New("暂无数据"), result
	}
	return returnCode.OK, returnMsg.OK, nil, result
}
