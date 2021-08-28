package usersService

import (
	"Gin/authUtils"
	"Gin/constant"
	"Gin/dao"
	"Gin/dto/usersDto"
	"Gin/models"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginResult struct {
	Name  string
	Token string
}

func Login(c *gin.Context, userDto *usersDto.LoginInfoDto) {
	//var loginInfo models.LoginInfo
	isPass, users, err := LoginCheck(userDto)
	if isPass {
		generateToken(c, users)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}
}

func LoginCheck(loginInfo *usersDto.LoginInfoDto) (bool, models.Users, error) {
	var user models.Users
	hash := sha256.New()
	hash.Write([]byte(constant.SALT))
	hash.Write([]byte(loginInfo.Password))
	passwd := hex.EncodeToString(hash.Sum(nil))
	dao.DB.Where("user_id = ? AND password = ?", loginInfo.UserID, passwd).Find(&user)
	if user.ID == 0 {
		return false, user, gin.Error{
			Err: errors.New("用户名或密码错误！"),
		}
	}

	return true, user, gin.Error{}
}

func generateToken(c *gin.Context, user models.Users) {
	// 构造SignKey
	j := authUtils.NewJWT()

	// 构造用户claims信息(负荷)
	claims := authUtils.UserClaims{
		UserId: user.UserId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),    // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600*24), // 签名过期时间
			Issuer:    "XuYiFei",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}

	//log.Println(token)
	// 封装一个相应数据,返回用户名和token
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功",
		"data": LoginResult{
			Name:  user.UserName,
			Token: token,
		},
	})
}

func LoginStatus(claim *authUtils.UserClaims) (bool, models.Users, error) {
	userId := claim.UserId
	var user models.Users
	dao.DB.Where("user_id = ?", userId).Find(&user)
	if user.ID == 0 {
		return false, user, errors.New("账号不存在,重新登录")
	}
	return true, user, nil
}
