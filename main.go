package main

import (
	"fmt"
	"os"

	"zdmin/datasource"
	"zdmin/routes"
	"zdmin/setting"
)

func main() {
	var filePath string
	if len(os.Args) <= 1 {
		filePath = "./config.yaml"
	} else {
		filePath = os.Args[1]
	}
	//加载配置
	setting.Init(filePath)

	//初始化日志

	// 初始化数据库
	datasource.Init()
	defer datasource.Close() // 程序退出关闭数据库连接

	r := routes.Init()
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		panic(err)
	}

}
