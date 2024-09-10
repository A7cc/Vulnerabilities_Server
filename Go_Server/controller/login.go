package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"fmt"

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
	// 判断是否绑定成功
	if err := c.ShouldBindJSON(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	if in.Code == "" {
		helper.ErrorResponse(c, "登录", fmt.Errorf("验证码不能为空"))
		return
	}
	// 根据账号和密码查询用户信息
	sysUser, err := models.GetUserByUsernamePassword(in.UserName, in.Password)
	if err != nil {
		helper.ErrorResponse(c, "登录", err)
		return
	}
	// 生成token
	authorization, err := helper.GenerateToken(sysUser.ID, sysUser.Role_id, sysUser.UserName, define.TokenExpire)
	if err != nil {
		helper.ErrorResponse(c, "生成token", err)
		return
	}
	// 获取角色信息
	sysRole, err := models.GetRoleDetail(uint(sysUser.Role_id))
	if err != nil {
		helper.ErrorResponse(c, "获取角色信息", err)
		return
	}
	introduce := "这是一个用Golang写的Web靶场，该系统是以食谱菜单管理系统为场景去编写，一种实战化形式的安全漏洞靶场，其中存在多个安全漏洞，需要我们去探索和发现。\n\n该项目旨在帮助安全研究人员和爱好者了解和掌握关于Golang系统的渗透测试和代码审计知识。项目地址：https://github.com/A7cc/Vulnerabilities_Server"
	helper.SuccessResponse(c, "登录", &define.LoginPasswordResponse{
		Uid:           sysUser.ID,
		Authorization: "Bearer " + authorization,
		Username:      sysUser.UserName,
		Avatar:        sysUser.Avatar,
		Phone:         sysUser.Phone,
		Sex:           sysUser.Sex,
		Email:         sysUser.Email,
		RoleLevel:     sysRole.Level,
		Role:          sysRole.Name,
		Introduce:     introduce,
		Created_at:    sysUser.CreatedAt.String(),
	})
}

// AuthLoginOut
// @Summary 处理用户注销
// @Tags 公共方法
// @Router /auth/loginout [get]
func AuthLoginOut(c *gin.Context) {
	helper.SuccessResponse(c, "注销", nil)
}
