package main

import (
	"Go_server/config"
	"Go_server/models"
	"Go_server/router"
)

// @title 食谱菜单管理系统靶场
// @version 1.0
// @description 这是一个用Golang写的Web靶场，该系统是以食谱菜单管理系统为场景去编写，一种实战化形式的安全漏洞靶场，其中存在多个安全漏洞，需要我们去探索和发现。该项目旨在帮助安全研究人员和爱好者了解和掌握关于Golang系统的渗透测试和代码审计知识。
// @contact.name Vulnerabilities_Server
// @contact.url https://github.com/A7cc/Vulnerabilities_Server
func main() {
	serverConfig := config.Config.Server
	// 初始化gorm.db
	models.NewGormDB()
	// 运行程序
	r := router.App()
	r.Run(":" + serverConfig.Port)
}
