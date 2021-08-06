package routers

import (
	"Gin/controller/users"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.Static("/static", "static")

	userGroup := r.Group("/users")
	{
		userGroup.Any("/login", users.UserLogin)
	}

	return r
}

