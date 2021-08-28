package classificationVO

import (
	"time"
)

type FieldVO struct {
	Num       int    `json:"num,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
}

type ClassificationFromFieldVO struct {
	Num                int    `json:"num,omitempty"`
	ClassificationName string `json:"classificationName,omitempty"`
}

type TagArticles struct {
	UserId    string    `json:"userId"`
	BlogId    string    `json:"blogId"`
	Title     string    `json:"title"`
	Digest    string    `json:"digest"`
	Tag       string    `json:"tag"`
	UpdatedAt time.Time `json:"updateAt"`
	Author    string    `json:"author"`
}

type TagBlogsVO struct {
	Tag   string        `json:"tag"`
	Num   int           `json:"num"`
	Blogs []TagArticles `json:"blogs"`
}
