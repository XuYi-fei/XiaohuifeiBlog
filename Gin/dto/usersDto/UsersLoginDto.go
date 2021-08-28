package usersDto

type LoginInfoDto struct {
	Password string `json:"password"`
	UserID   string `json:"userId"`
}

type RegisterInfoDto struct {
	Password string `json:"password,omitempty"`
	UserId   string `json:"userId,omitempty"`
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
	UserName string `json:"userName,omitempty"`
}
