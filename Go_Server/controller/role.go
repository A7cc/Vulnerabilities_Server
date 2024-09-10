package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRoleList
// @Summary 获取角色列表信息
// @Tags 鉴权接口-角色相关方法
// @Param Authorization header string true "Authorization"
// @Param GetRoleListRequest body define.GetRoleListRequest true "获取角色查询参数"
// @Router /role/get [get]
func GetRoleList(c *gin.Context) {
	in := &define.GetRoleListRequest{QueryRequest: NewQueryRequest()}
	if err := c.ShouldBindQuery(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	var (
		cnt  int64
		list = make([]*define.GetRoleListReply, 0)
	)

	if err := models.GetRoleList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error; err != nil {
		helper.ErrorResponse(c, "获取角色列表信息", err)
		return
	}
	helper.SuccessResponse(c, "获取角色列表信息", gin.H{
		"list":  list,
		"count": cnt,
	})
}

// AddRole
// @Summary 新增角色信息
// @Tags 鉴权接口-角色相关方法
// @Param Authorization header string true "Authorization"
// @Param AddRoleRequest body define.AddRoleRequest true "添加角色信息"
// @Router /role/add [post]
func AddRole(c *gin.Context) {
	in := new(define.AddRoleRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	// 1.判断角色是否存在
	var cnt int64
	// 大于0说明存在用户
	if err := models.DB.Model(new(models.SysRole)).Where("name = ?", in.Name).Count(&cnt).Error; cnt > 1 && err == nil {
		helper.ErrorResponse(c, "新增角色信息", fmt.Errorf("用户名已经存在"))
		return
	}
	// 2.保存数据
	if err := models.DB.Create(&models.SysRole{
		Name:    in.Name,
		Level:   in.Level,
		Remarks: in.Remarks,
	}).Error; err != nil {
		helper.ErrorResponse(c, "新增角色信息", err)
		return
	}
	helper.SuccessResponse(c, "新增角色信息", nil)
}

// GetRoleDetail
// @Summary 获取角色详细信息
// @Tags 鉴权接口-角色相关方法
// @Param Authorization header string true "Authorization"
// @Param id query int true "获取角色ID"
// @Router /role/detail [get]
func GetRoleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		helper.ErrorResponse(c, "获取角色详细信息", fmt.Errorf("id不能为空"))
		return
	}
	Uid, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorResponse(c, "获取角色详细信息", err)
		return
	}
	// 获取角色信息
	sysRole, err := models.GetRoleDetail(uint(Uid))
	if err != nil {
		helper.ErrorResponse(c, "获取角色详细信息", err)
		return
	}
	data := new(define.GetRoleDetailReply)
	// 赋值
	data.ID = sysRole.ID
	data.Name = sysRole.Name
	data.Level = sysRole.Level
	data.Remarks = sysRole.Remarks
	// 返回角色信息
	helper.SuccessResponse(c, "获取角色详细信息", data)
}

// UpdateRole
// @Summary 修改角色信息
// @Tags 鉴权接口-角色相关方法
// @Param Authorization header string true "Authorization"
// @Param UpdateRoleRequest body define.UpdateRoleRequest true "修改角色信息参数"
// @Router /role/update [put]
func UpdateRole(c *gin.Context) {
	in := new(define.UpdateRoleRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	// 验证ID是否存在
	if _, err := models.GetRoleDetail(in.ID); err != nil {
		helper.ErrorResponse(c, "修改角色信息", err)
		return
	}
	// 1.判断角色是否已存在
	var cnt int64
	if err := models.DB.Model(new(models.SysRole)).Where("id != ? AND name = ?", in.ID, in.Name).Count(&cnt).Error; err != nil {
		helper.ErrorResponse(c, "修改角色信息", err)
		return
	}
	if cnt > 0 {
		helper.ErrorResponse(c, "修改角色信息", fmt.Errorf("角色名已经存在"))
		return
	}
	// 获取角色信息
	sysRole, err := models.GetRoleDetail(in.ID)
	if err != nil {
		helper.ErrorResponse(c, "修改角色信息", err)
		return
	}
	// 判断用户是否为1
	if sysRole.ID == 1 && in.Level != 1 {
		helper.ErrorResponse(c, "修改角色信息", fmt.Errorf("admin最高管理角色等级不能被修改"))
		return
	}
	// 修改数据
	if err := models.DB.Model(new(models.SysRole)).Where("id = ?", in.ID).Updates(map[string]any{
		"name":    in.Name,
		"level":   in.Level,
		"remarks": in.Remarks,
	}).Error; err != nil {
		helper.ErrorResponse(c, "修改角色信息", err)
		return
	}
	helper.SuccessResponse(c, "修改角色信息", nil)
}

// DeleteRole
// @Summary 删除角色信息
// @Tags 鉴权接口-角色相关方法
// @Param Authorization header string true "Authorization"
// @Param id path int true "删除角色ID"
// @Router /role/delete/{id} [delete]
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ErrorResponse(c, "删除角色信息", fmt.Errorf("id不能为空"))
		return
	}
	Fid, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorResponse(c, "删除角色信息", err)
		return
	}
	// 获取角色信息
	sysRole, err := models.GetRoleDetail(uint(Fid))
	if err != nil {
		helper.ErrorResponse(c, "删除角色信息", err)
		return
	}
	// 判断用户是否为1
	if sysRole.ID == 1 {
		helper.ErrorResponse(c, "删除角色信息", fmt.Errorf("admin最高管理角色不能被删除"))
		return
	}
	// 删除角色
	if err = models.DB.Where("id = ?", id).Delete(new(models.SysRole)).Error; err != nil {
		helper.ErrorResponse(c, "删除角色信息", err)
		return
	}
	helper.SuccessResponse(c, "删除角色信息", nil)
}
