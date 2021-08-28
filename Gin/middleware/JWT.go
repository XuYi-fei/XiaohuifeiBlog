package middleware

import (
	"Gin/authUtils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "登录信息已过期或无效,请重新登录",
				"data":   nil,
			})
			c.Abort()
			return
		}

		//log.Print("get token:", token)

		// 初始化一个JWT对象实例，并且根据结构体方法解析token
		j := authUtils.NewJWT()
		// 解析token中包含的相关信息(有效载荷)
		claims, err := j.ParseToken(token)

		if err != nil {
			// token 过期
			if err == fmt.Errorf("token过期") {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "token授权已过期",
					"data":   nil,
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   nil,
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
