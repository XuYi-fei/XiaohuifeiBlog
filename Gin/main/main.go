package main

import (
	"Gin/dao"
	"Gin/routers"
)

func main(){
	err := dao.InitMySQL()
	if err != nil{
		panic(err)
	}
	//defer dao.DB.Close()
	err = dao.InitMySQL()
	if err != nil {
		return
	}
	//r := dao.DB
	r := routers.SetupRouter()


	r.Run(":18081")
}