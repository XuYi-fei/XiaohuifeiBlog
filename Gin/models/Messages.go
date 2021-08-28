package models

import "time"

type Messages struct {
	Id           int       `gorm:"AUTO_INCREMENT;column:id" json:"id,omitempty"`
	FromUserId   string    `gorm:"column:from_user_id" json:"fromUserId"`
	FromUserName string    `gorm:"column:from_user_name" json:"fromUserName"`
	UserName     string    `gorm:"column:user_name" json:"userName"`
	UserId       string    `gorm:"column:user_id" json:"userId,omitempty"`
	Content      string    `gorm:"column:content" json:"content,omitempty"`
	IsRead       bool      `gorm:"column:is_read" json:"isRead,omitempty"`
	Title        string    `gorm:"column:title" json:"title,omitempty"`
	Digest       string    `gorm:"column:digest" json:"digest,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"createdAt"`
}
