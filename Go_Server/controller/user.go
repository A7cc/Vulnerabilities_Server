package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 获取用户列表信息
// @Tags 鉴权接口-用户相关方法
// @Param Authorization header string true "Authorization"
// @Param GetUserListRequest body define.GetUserListRequest true "获取用户查询参数"
// @Router /user/get [get]
func GetUserList(c *gin.Context) {
	in := &define.GetUserListRequest{QueryRequest: NewQueryRequest()}
	if err := c.ShouldBindQuery(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	var (
		cnt  int64
		list = make([]*define.GetUserListReply, 0)
	)
	if err := models.GetUserList(in.Keyword, in.Status).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error; err != nil {
		helper.ErrorResponse(c, "获取用户列表信息", err)
		return
	}
	helper.SuccessResponse(c, "获取用户列表信息", gin.H{
		"list":  list,
		"count": cnt,
	})
}

// AddUser
// @Summary 新增用户信息
// @Tags 鉴权接口-用户相关方法
// @Param Authorization header string true "Authorization"
// @Param AddUserRequest body define.AddUserRequest true "新增用户信息"
// @Router /user/add [post]
func AddUser(c *gin.Context) {
	in := new(define.AddUserRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	// 1.判断用户名是否存在
	var cnt int64
	// 大于0说明存在用户
	if err := models.DB.Model(new(models.SysUser)).Where("username = ?", in.Username).Count(&cnt).Error; cnt > 0 && err == nil {
		helper.ErrorResponse(c, "新增用户信息", fmt.Errorf("用户名已经存在"))
		return
	}
	// 大于0说明存在角色
	if err := models.DB.Model(new(models.SysRole)).Where("id = ?", in.Role_id).Count(&cnt).Error; cnt == 0 && err == nil {
		helper.ErrorResponse(c, "新增用户信息", fmt.Errorf("角色不存在"))
		return
	}
	// 验证手机
	if ok := helper.ValidatePhone(in.Phone); !ok {
		helper.ErrorResponse(c, "新增用户信息", fmt.Errorf("手机号码不正确"))
		return
	}
	// 验证邮箱
	if ok := helper.ValidateEmail(in.Email); !ok {
		helper.ErrorResponse(c, "新增用户信息", fmt.Errorf("邮箱不正确"))
		return
	}
	// 2.保存数据
	if err := models.DB.Create(&models.SysUser{
		UserName: in.Username,
		PassWord: in.Password,
		Phone:    in.Phone,
		Status:   in.Status,
		Role_id:  in.Role_id,
		Sex:      in.Sex,
		Email:    in.Email,
		Remarks:  in.Remarks,
	}).Error; err != nil {
		helper.ErrorResponse(c, "新增用户信息", err)
		return
	}
	helper.SuccessResponse(c, "新增用户信息", nil)
}

// GetUserDetail
// @Summary 获取管理员详细信息
// @Tags 鉴权接口-用户相关方法
// @Param Authorization header string true "Authorization"
// @Param id query int true "获取用户ID"
// @Router /user/detail [get]
func GetUserDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		helper.ErrorResponse(c, "获取管理员详细信息", fmt.Errorf("ID不能为空"))
		return
	}
	Uid, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorResponse(c, "获取管理员详细信息", err)
		return
	}

	// 获取管理员信息
	sysUser, err := models.GetUserDetail(uint(Uid))
	if err != nil {
		helper.ErrorResponse(c, "获取管理员详细信息", err)
		return
	}
	data := new(define.GetUserDetailReply)
	// 赋值
	data.ID = sysUser.ID
	data.Username = sysUser.UserName
	data.Password = sysUser.PassWord
	data.Sex = sysUser.Sex
	data.Status = sysUser.Status
	data.Role_id = sysUser.Role_id
	data.Phone = sysUser.Phone
	data.Email = sysUser.Email
	data.Remarks = sysUser.Remarks
	// 返回管理员信息
	helper.SuccessResponse(c, "获取管理员详细信息", data)
}

// UpdateUser
// @Summary 修改用户信息
// @Tags 鉴权接口-用户相关方法
// @Param Authorization header string true "Authorization"
// @Param UpdateUserRequest body define.UpdateUserRequest true "更新用户信息参数"
// @Router /user/update [put]
func UpdateUser(c *gin.Context) {
	in := new(define.UpdateUserRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	// 验证ID是否存在
	if _, err := models.GetUserDetail(in.ID); err != nil {
		helper.ErrorResponse(c, "修改用户信息", err)
		return
	}
	// 1.判断用户名是否已存在
	var cnt int64
	if err := models.DB.Model(new(models.SysUser)).Where("id != ? AND username = ?", in.ID, in.Username).Count(&cnt).Error; err != nil {
		helper.ErrorResponse(c, "修改用户信息", err)
		return
	}
	if cnt > 0 {
		helper.ErrorResponse(c, "修改用户信息", fmt.Errorf("用户名已经存在"))
		return
	}
	// 验证手机
	if ok := helper.ValidatePhone(in.Phone); !ok {
		helper.ErrorResponse(c, "修改用户信息", fmt.Errorf("手机号码不正确"))
		return
	}
	// 验证邮箱
	if ok := helper.ValidateEmail(in.Email); !ok {
		helper.ErrorResponse(c, "修改用户信息", fmt.Errorf("邮箱不正确"))
		return
	}
	// 判断用户是否为1
	if in.ID == 1 && (in.Role_id != 1 || !in.Status) {
		helper.ErrorResponse(c, "修改用户信息", fmt.Errorf("admin最高管理员用户角色或封禁状态不能被修改"))
		return
	}
	// 修改数据
	if err := models.DB.Model(new(models.SysUser)).Where("id = ?", in.ID).Updates(map[string]any{
		"username": in.Username,
		"password": in.Password,
		"sex":      in.Sex,
		"role_id":  in.Role_id,
		"status":   in.Status,
		"phone":    in.Phone,
		"email":    in.Email,
		"remarks":  in.Remarks,
	}).Error; err != nil {
		helper.ErrorResponse(c, "修改用户信息", err)
		return
	}
	helper.SuccessResponse(c, "修改用户信息", nil)
}

// DeleteUser
// @Summary 删除用户信息
// @Tags 鉴权接口-用户相关方法
// @Param Authorization header string true "Authorization"
// @Param id path int true "删除用户ID"
// @Router /user/delete/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ErrorResponse(c, "删除用户信息", fmt.Errorf("ID不能为空"))
		return
	}
	Fid, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorResponse(c, "删除用户信息", err)
		return
	}
	// 获取用户信息
	sysUser, err := models.GetUserDetail(uint(Fid))
	if err != nil {
		helper.ErrorResponse(c, "删除用户信息", err)
		return
	}
	// 判断用户是否为1
	if sysUser.ID == 1 {
		helper.ErrorResponse(c, "删除用户信息", fmt.Errorf("admin最高管理员用户不能被删除"))
		return
	}
	if err = models.DB.Where("id = ?", id).Delete(new(models.SysUser)).Error; err != nil {
		helper.ErrorResponse(c, "删除用户信息", err)
		return
	}
	helper.SuccessResponse(c, "删除用户信息", nil)
}
