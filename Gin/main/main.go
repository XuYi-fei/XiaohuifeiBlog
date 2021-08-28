package main

import (
	"Gin/dao"
	"Gin/routers"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	//defer dao.DB.Close()
	//r := dao.DB
	r := routers.SetupRouter()

	r.Run(":18081")
}
