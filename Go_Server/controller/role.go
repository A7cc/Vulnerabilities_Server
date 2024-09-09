package controller

import (
	"Go_server/define"
	"Go_server/models"
	"net/http"
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
	err := c.ShouldBindQuery(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常",
		})
		return
	}
	var (
		cnt  int64
		list = make([]*define.GetRoleListReply, 0)
	)
	err = models.GetRoleList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "加载数据成功",
		"result": gin.H{
			"list":  list,
			"count": cnt,
		},
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
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常" + err.Error(),
		})
		return
	}
	// 1.判断角色是否存在
	var cnt int64
	err = models.DB.Model(new(models.SysRole)).Where("name = ?", in.Name).Count(&cnt).Error
	// 大于0说明存在用户
	if cnt > 1 && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败，用户名已经存在",
		})
		return
	}
	// 2.保存数据
	err = models.DB.Create(&models.SysRole{
		Name:    in.Name,
		Level:   in.Level,
		Remarks: in.Remarks,
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败，数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "保存成功",
	})
}

// GetRoleDetail
// @Summary 根据ID获取角色信息
// @Tags 鉴权接口-角色相关方法
// @Param Authorization header string true "Authorization"
// @Param id query int true "获取角色ID"
// @Router /role/detail [get]
func GetRoleDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "ID不能为空",
		})
		return
	}
	Uid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据转换异常",
		})
		return
	}
	data := new(define.GetRoleDetailReply)
	// 1.获取角色信息
	sysRole, err := models.GetRoleDetail(uint(Uid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常" + err.Error(),
		})
		return
	}
	// 赋值
	data.ID = sysRole.ID
	data.Name = sysRole.Name
	data.Level = sysRole.Level
	data.Remarks = sysRole.Remarks
	// 返回角色信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取数据成功",
		"result":  data,
	})
}

// UpdateRole
// @Summary 修改角色信息
// @Tags 鉴权接口-角色相关方法
// @Param Authorization header string true "Authorization"
// @Param UpdateRoleRequest body define.UpdateRoleRequest true "更新角色信息参数"
// @Router /role/update [put]
func UpdateRole(c *gin.Context) {
	in := new(define.UpdateRoleRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常",
		})
		return
	}
	// 验证ID是否存在
	_, err = models.GetRoleDetail(in.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "该角色不存在",
		})
		return
	}
	// 1.判断角色是否已存在
	var cnt int64
	err = models.DB.Model(new(models.SysRole)).Where("id != ? AND name = ?", in.ID, in.Name).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "更新失败，角色名已经存在",
		})
		return
	}
	// 获取角色信息
	sysRole, err := models.GetRoleDetail(in.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常" + err.Error(),
		})
		return
	}
	// 判断用户是否为1
	if sysRole.ID == 1 && in.Level != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "修改失败，admin最高管理角色等级不能被修改",
		})
		return
	}
	// 2.修改数据
	err = models.DB.Model(new(models.SysRole)).Where("id = ?", in.ID).Updates(map[string]any{
		"name":    in.Name,
		"level":   in.Level,
		"remarks": in.Remarks,
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "更新失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "修改成功",
	})
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
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败，ID不能为空",
		})
		return
	}
	Fid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据转换异常",
		})
		return
	}
	// 获取角色信息
	sysRole, err := models.GetRoleDetail(uint(Fid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常" + err.Error(),
		})
		return
	}
	// 判断用户是否为1
	if sysRole.ID == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败，admin最高管理角色不能被删除",
		})
		return
	}

	if err = models.DB.Where("id = ?", id).Delete(new(models.SysRole)).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败，数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
