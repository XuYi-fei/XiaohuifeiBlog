package users

import (
	"Gin/dto/users"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context){
	//json := make(map[string]interface{})
	//b, _ := c.GetRawData()  // 从c.Request.Body读取请求数据
	// 定义map或结构体
	//var m map[string]interface{}
	// 反序列化
	//_ = json.Unmarshal(b, &m)
	var json users.UserLoginDto
	c.Bind(&json)
	//data, _ := ioutil.ReadAll(c.Request.Body)
	////json, _ := c.GetRawData()
	fmt.Println("%v",&json)
	//fmt.Println("%v",data)
	//fmt.Println("%v",m)
	//fmt.Println(2)
}
