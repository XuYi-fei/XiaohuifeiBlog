package models

// Todo Model
type Users struct {
	ID int `gorm:"column:id"`
	UserName string `gorm:"AUTO_INCREMENT;column:user_name"`
	Password string `gorm:"column:pass_word"`
	UserID string `gorm:"column:user_id"`
}

// Todo 增删改查
func create(users *Users) error{
	return nil
}