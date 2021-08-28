package blogsVO

import (
	"Gin/models"
	"time"
)

type BlogVO struct {
	UserId         string    `json:"userId"`
	BlogId         string    `json:"blogId"`
	Title          string    `json:"title"`
	Digest         string    `json:"digest"`
	Content        string    `json:"content"`
	Tag            string    `json:"tag"`
	CreatedTime    time.Time `json:"createdTime"`
	Author         string    `json:"author"`
	Field          string    `json:"field"`
	Classification string    `json:"classification"`
}

func NewBlogVOFromModel(blog models.Article, str string) *BlogVO {
	return &BlogVO{
		UserId:         blog.UserId,
		BlogId:         blog.BlogId,
		Title:          blog.Title,
		Digest:         blog.Digest,
		Tag:            blog.Tag,
		Field:          blog.Field,
		Classification: blog.Classification,
		Author:         blog.Author,
		Content:        str,
		CreatedTime:    blog.CreatedAt,
	}
}

type BlogTmpVO struct {
	Id             int    `json:"id"`
	BlogId         string `json:"blogId"`
	Title          string `json:"title"`
	Digest         string `json:"digest"`
	Content        string `json:"content"`
	UserId         string `json:"userId"`
	Tag            string `json:"tag"`
	Author         string `json:"author"`
	Field          string `json:"field"`
	Classification string `json:"classification"`
	FilePath       string `json:"filePath"`
}

type BlogInfos struct {
	BlogNum           int `json:"blogsNum"`
	FieldNum          int `json:"fieldsNum"`
	ClassificationNum int `json:"classificationNum"`
	TagNum            int `json:"tagNum"`
}

// 返回分页数量
type BlogsPageNumVO struct {
	PageNum int `json:"pageNum"`
}
