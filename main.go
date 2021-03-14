package main

import (
	"go_api/model/mysql_model"
	"go_api/router"
)

// @title Go API接口文档
// @version 1.0
// @description HTTP API服务

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 1065680788@qq.com

// @host 127.0.0.1
// @BasePath /v1
func main() {
	mysql_model.Init()
	defer mysql_model.CloseDB()
	r := router.InitRouter()

	_ = r.Run()

}
