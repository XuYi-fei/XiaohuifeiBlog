package models

// Todo Model

type Users struct {
	ID        int    `gorm:"AUTO_INCREMENT;column:id" json:"id"`
	UserName  string `gorm:"column:user_name" json:"userName"`
	Password  string `gorm:"column:password" json:"password"`
	UserId    string `gorm:"column:user_id" json:"userId"`
	Sex       string `gorm:"column:sex;default:暂未填写" json:"sex"`
	Age       string `gorm:"column:age;default:'暂未填写'" json:"age"`
	Privilege int    `gorm:"column:privilege;default:0" json:"privilege"`
	RealName  string `gorm:"column:real_name;default:'暂未填写'" json:"realName"`
	BlogNum   int    `gorm:"column:blog_num;default:0" json:"blogNum"`
	Question  string `gorm:"column:question" json:"question"`
	Answer    string `gorm:"column:answer" json:"answer"`
}

func NewUsers(userName string, password string, userID string, question string, answer string) *Users {
	return &Users{UserName: userName, Password: password, UserId: userID, Question: question, Answer: answer}
}

// Todo 增删改查
func create(users *Users) error {
	return nil
}
