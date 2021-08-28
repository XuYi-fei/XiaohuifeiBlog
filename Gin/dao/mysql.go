package dao

import (
	"Gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:xuyifei@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}
	err = DB.AutoMigrate(&models.Users{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&models.Article{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.ArticleTmp{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.Sentence{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.SentencesTmp{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.SentenceShow{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.Messages{})
	if err != nil {
		panic(err)
	}
	//err = DB.AutoMigrate(&models.Classification{})
	//if err != nil {
	//	panic(err)
	//}
	//err = DB.AutoMigrate(&models.Field{})
	//if err != nil {
	//	panic(err)
	//}
	return err
}
