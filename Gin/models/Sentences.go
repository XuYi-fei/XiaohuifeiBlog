package models

import "time"

type Sentence struct {
	ID        int       `gorm:"column:id" json:"id"`
	Content   string    `gorm:"column:content" json:"content"`
	Author    string    `gorm:"column:author;default:''" json:"author"`
	UserId    string    `gorm:"column:user_id" json:"userId"`
	UserName  string    `gorm:"column:user_name" json:"userName"`
	CreatedAt time.Time `gorm:"column:created_time" json:"createdTime"`
}

func NewSentence(content string, author string, userId string, userName string) *Sentence {
	return &Sentence{Content: content, Author: author, UserId: userId, UserName: userName}
}

type SentencesTmp struct {
	ID       int    `gorm:"column:id" json:"id"`
	Content  string `gorm:"column:content" json:"content"`
	Author   string `gorm:"column:author;default:''" json:"author"`
	UserId   string `gorm:"column:user_id" json:"userId"`
	UserName string `gorm:"column:user_name" json:"userName"`
	//CreatedAt time.Time `gorm:"column:created_time" json:"createdTime"`
}

type SentenceShow struct {
	ID      int    `gorm:"column:id" json:"id"`
	Content string `gorm:"column:content" json:"content"`
}
