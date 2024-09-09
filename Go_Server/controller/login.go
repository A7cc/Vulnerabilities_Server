package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthLogin
// @Summary 处理用户登录
// @Tags 公共方法
// @Param user body define.LoginPassWordRequest true "登录信息"
// @Router /auth/login [post]
func AuthLogin(c *gin.Context) {
	// 创建接收用户输入参数
	in := new(define.LoginPassWordRequest)
	// 绑定参数
	err := c.ShouldBindJSON(in)
	// 判断是否绑定成功
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	if in.Code == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "验证码不能为空",
		})
		return
	}
	// 根据账号和密码查询用户信息
	sysUser, err := models.GetUserByUsernamePassword(in.UserName, in.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	// 生成token
	authorization, err := helper.GenerateToken(sysUser.ID, sysUser.Role_id, sysUser.UserName, define.TokenExpire)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	data := &define.LoginPasswordReply{
		Authorization: "Bearer " + authorization,
	}
	// 1.获取角色信息
	sysRole, err := models.GetRoleDetail(uint(sysUser.Role_id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常" + err.Error(),
		})
		return
	}
	introduce := "这是一个用Golang写的Web靶场，该系统是以食谱菜单管理系统为场景去编写，一种实战化形式的安全漏洞靶场，其中存在多个安全漏洞，需要我们去探索和发现。\n\n该项目旨在帮助安全研究人员和爱好者了解和掌握关于Golang系统的渗透测试和代码审计知识。项目地址：https://github.com/A7cc/Vulnerabilities_Server"
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登陆成功",
		"result": &define.LoginPasswordResponse{
			Uid:           sysUser.ID,
			Authorization: data.Authorization,
			Username:      sysUser.UserName,
			Avatar:        sysUser.Avatar,
			Phone:         sysUser.Phone,
			Sex:           sysUser.Sex,
			Email:         sysUser.Email,
			RoleLevel:     sysRole.Level,
			Role:          sysRole.Name,
			Introduce:     introduce,
			Created_at:    sysUser.CreatedAt.String(),
		},
	})
}

// AuthLoginOut
// @Summary 处理用户注销
// @Tags 公共方法
// @Router /auth/loginout [get]
func AuthLoginOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注销成功",
	})
}
