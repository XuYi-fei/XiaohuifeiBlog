package models

import "time"

type Article struct {
	Id             int       `gorm:"AUTO_INCREMENT;column:id" json:"id,omitempty"`
	BlogId         string    `gorm:"column:blog_id" json:"blogId,omitempty"`
	UserId         string    `gorm:"column:user_id" json:"userId,omitempty"`
	Title          string    `gorm:"column:title" json:"title,omitempty"`
	Digest         string    `gorm:"column:digest" json:"digest,omitempty"`
	Content        string    `gorm:"column:content" json:"content,omitempty"`
	Author         string    `gorm:"column:author" json:"author,omitempty"`
	Field          string    `gorm:"field;default:其他" json:"field,omitempty"`
	Classification string    `gorm:"column:classification;default:其他" json:"classification,omitempty"`
	Tag            string    `gorm:"column:tag;default:其他" json:"tag,omitempty"`
	Hot            uint8     `gorm:"column:hot;comment:热度;default:0" json:"hot"`
	CreatedAt      time.Time `json:"createAt"`
	UpdatedAt      time.Time `json:"updateAt"`
}

func NewArticle(blogId string, userId string, title string, digest string, content string,
	author string, field string, classification string, tag string) *Article {
	return &Article{BlogId: blogId, UserId: userId, Title: title, Digest: digest, Content: content, Author: author, Field: field, Classification: classification,
		Tag: tag}
}

func NewBlogs() interface{} {
	return []Article{}
}

type ArticleTmp struct {
	ID             int    `gorm:"AUTO_INCREMENT;column:id" json:"id,omitempty"`
	BlogId         string `gorm:"column:blog_id" json:"blogId,omitempty"`
	UserID         string `gorm:"column:user_id" json:"userId,omitempty"`
	Title          string `gorm:"column:title" json:"title,omitempty"`
	Digest         string `gorm:"column:digest" json:"digest,omitempty"`
	Content        string `gorm:"column:content" json:"content,omitempty"`
	Author         string `gorm:"column:author" json:"author,omitempty"`
	Field          string `gorm:"field;default:其他" json:"field,omitempty"`
	Classification string `gorm:"column:classification;default:其他" json:"classification,omitempty"`
	Tag            string `gorm:"column:tag;default:其他" json:"tag,omitempty"`
}

func NewArticleTmp(blogId string, userID string, title string, digest string, content string, author string, field string, classification string, tag string) *ArticleTmp {
	return &ArticleTmp{BlogId: blogId, UserID: userID, Title: title, Digest: digest, Content: content, Author: author, Field: field, Classification: classification, Tag: tag}
}

//type Field struct {
//	ID             int       `gorm:"AUTO_INCREMENT;column:id" json:"id,omitempty"`
//	UserID         string    `gorm:"column:user_id" json:"userId,omitempty"`
//	Num 		   int       `gorm:"column:num" json:"num"`
//	FieldName      string    `gorm:"column:field_name" json:"fieldName"`
//}
//
//type Classification struct {
//	ID             int       `gorm:"AUTO_INCREMENT;column:id" json:"id,omitempty"`
//	UserID         string    `gorm:"column:user_id" json:"userId,omitempty"`
//	Num 		   int       `gorm:"column:num" json:"num"`
//	ClassificationName string `gorm:"column:classification_name" json:"classificationName"`
//}
//
