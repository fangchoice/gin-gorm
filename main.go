package main

import (
	"fmt"
	"go_web/conf"
	"go_web/db/mysql"
	"go_web/route"

	logging "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello World!")

	// 路由
	r := route.NewRoute()

	// 运行服务
	err := r.Run(conf.HttpPort)
	if err != nil {
		logging.Fatalln(err)
	}
}

// 初始化配置
func init() {
	conf.Init("./conf/config.ini")
	mysql.Init()
}
