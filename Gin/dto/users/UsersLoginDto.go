package users

type UserLoginDto struct {
	UserId string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}