package router

import (
	"Go_server/config"
	"Go_server/controller"
	_ "Go_server/docs"
	"Go_server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 初始化路由
func App() *gin.Engine {
	r := gin.Default()
	// 跌机恢复
	r.Use(gin.Recovery())
	// 添加跨域中间件
	r.Use(middleware.Cors())
	// 添加日志
	r.Use(middleware.LoggerToFile())
	register(r)
	return r
}

// 路由注册
func register(r *gin.Engine) {
	// 静态资源文件
	r.StaticFS("/static", http.Dir(config.Config.StaticData))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 根据用户名和密码登录路由
	r.POST("/auth/login", controller.AuthLogin)
	r.GET("/auth/loginout", controller.AuthLoginOut)
	// 主页设置组api
	homeApi(r)
	// settings设置组api
	settingsApi(r)
	// 管理员组api
	userApi(r)
	// 角色组api
	roleApi(r)
	// 食物组api
	foodApi(r)
	// 订单组api
	orderApi(r)
	// 用于测试一些新型功能，不外公开，不会有对应前端映射
	otherApi(r)
}

// home设置组
func homeApi(r *gin.Engine) {
	homeapi := r.Group("home")
	{
		// 首页设置 start
		// 获取项目系统的
		homeapi.GET("/get", controller.GetHomeInfo)
		// 更新个人信息
		homeapi.PUT("/updateInfo", controller.UpdateInfo)
		// 更改个人密码
		homeapi.PUT("/updatePwd", controller.UpdatePwd)
		// 更新头像
		homeapi.POST("/upuseravatar", controller.UploadUserAvatar)
		// 获取今日美句
		homeapi.GET("/getsentence", controller.GetSentence)
		// 首页 end
	}
}

// settings设置组
func settingsApi(r *gin.Engine) {
	settingsapi := r.Group("settings").Use(middleware.Auth())
	{
		// 系统设置 start
		// 测试连通性
		settingsapi.POST("/ping", controller.Ping)
		// 获取备份数据库列表
		settingsapi.POST("/getdb", controller.GetBackupsDb)
		// 数据库备份
		settingsapi.GET("/backupsdb", controller.BackupsDb)
		// 数据库删除
		settingsapi.POST("/deletedb", controller.DeleteDb)
		// 数据库下载
		settingsapi.POST("/downdb", controller.DownDb)
		// 系统设置 end
	}
}

// userapi组
func userApi(r *gin.Engine) {
	userapi := r.Group("user").Use(middleware.Auth())
	{
		// 管理员 start
		// 获取管理员列表
		userapi.GET("/get", controller.GetUserList)
		// 添加管理员
		userapi.POST("/add", controller.AddUser)
		// 获取管理员详细信息
		userapi.GET("/detail", controller.GetUserDetail)
		// 更新管理员信息
		userapi.PUT("/update", controller.UpdateUser)
		// 删除管理员信息
		userapi.DELETE("/delete/:id", controller.DeleteUser)
		// 管理员 end
	}
}

// roleApi组
func roleApi(r *gin.Engine) {
	roleapi := r.Group("role").Use(middleware.Auth())
	{
		// 角色 start
		roleapi.GET("/get", controller.GetRoleList)
		// 添加角色
		roleapi.POST("/add", controller.AddRole)
		// 获取角色详细信息
		roleapi.GET("/detail", controller.GetRoleDetail)
		// 更新角色信息
		roleapi.PUT("/update", controller.UpdateRole)
		// 删除角色信息
		roleapi.DELETE("/delete/:id", controller.DeleteRole)
		// 角色 end
	}
}

// foodapi组
func foodApi(r *gin.Engine) {
	foodapi := r.Group("food").Use(middleware.Auth())
	{
		// 食物 start
		// 获取食物列表
		foodapi.GET("/get", controller.GetFoodList)
		// 添加食物
		foodapi.POST("/add", controller.AddFood)
		// 获取食物详细信息
		foodapi.GET("/detail", controller.GetFoodDetail)
		// 更新食物信息
		foodapi.PUT("/update", controller.UpdateFood)
		// 删除食物信息
		foodapi.DELETE("/delete/:id", controller.DeleteFood)
		// 上传食物icon
		foodapi.POST("/upfoodicon", controller.UploadFoodIcon)
		// 上传食物视频
		foodapi.POST("/upfoodvideo", controller.UploadFoodVideo)
		// 食物 end
	}
}

// orderapi组
func orderApi(r *gin.Engine) {
	orderapi := r.Group("order").Use(middleware.Auth())
	{
		// 订单 start
		// 获取订单列表
		orderapi.GET("/get", controller.GetOrderList)
		// 添加订单
		orderapi.POST("/add", controller.AddOrder)
		// 获取订单详细信息
		orderapi.GET("/detail", controller.GetOrderDetail)
		// 删除订单信息
		orderapi.DELETE("/delete/:id", controller.DeleteOrder)
		// 订单 end
	}
}
func otherApi(r *gin.Engine) {
	otherapi := r.Group("other").Use(middleware.Auth())
	{
		// 其他 start
		// 测试原生模板
		otherapi.GET("/ceshitp", controller.CeshiTemplate)
		// 测试上传ZIP并解压功能
		otherapi.POST("/uploadzip", controller.UploadZip)
		// 获取其他详细信息
	}
}
