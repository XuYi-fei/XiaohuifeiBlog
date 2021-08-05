package main

import (
	"Gin/controller/users"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println(1)
	users.Test()
	r := gin.Default()
	r.Run()
}