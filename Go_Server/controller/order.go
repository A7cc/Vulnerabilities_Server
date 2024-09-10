package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetOrderList
// @Summary 获取订单列表
// @Tags 鉴权接口-订单相关方法
// @Param Authorization header string true "Authorization"
// @Param GetOrderListRequest body define.GetOrderListRequest true "获取订单查询参数"
// @Router /order/get [get]
func GetOrderList(c *gin.Context) {
	in := &define.GetOrderListRequest{QueryRequest: NewQueryRequest()}
	if err := c.ShouldBindQuery(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	var (
		cnt  int64
		list = make([]*define.GetOrderListReply, 0)
	)

	if err := models.GetOrderList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error; err != nil {
		helper.ErrorResponse(c, "获取订单列表", err)
		return
	}
	helper.SuccessResponse(c, "获取订单列表", gin.H{
		"list":  list,
		"count": cnt,
	})
}

// AddOrder
// @Summary 新增订单信息
// @Tags 鉴权接口-订单相关方法
// @Param Authorization header string true "Authorization"
// @Param AddOrderRequest body define.AddOrderRequest true "添加订单信息"
// @Router /order/add [post]
func AddOrder(c *gin.Context) {
	in := new(define.AddOrderRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	// 1.判断食物是否存在
	var cnt int64
	// 大于0说明存在食物
	if err := models.DB.Model(new(models.SysFood)).Where("foodname = ?", in.Food).Count(&cnt).Error; cnt <= 0 || err != nil {
		helper.ErrorResponse(c, "新增订单信息", fmt.Errorf("可能不存在该食物"))
		return
	}
	// 解密密钥
	// 获取用户名的基本信息
	uinfo, err := helper.GetAuthorizationUserInfo(c.Request.Header.Get("Authorization"))
	if err != nil {
		helper.ErrorResponse(c, "新增订单信息", err)
		return
	}

	// 保存数据
	if err := models.DB.Create(&models.SysOrder{
		User:    uinfo.Name,
		Food:    in.Food,
		Num:     in.Num,
		Remarks: in.Remarks,
	}).Error; err != nil {
		helper.ErrorResponse(c, "新增订单信息", err)
		return
	}
	helper.SuccessResponse(c, "新增订单信息", nil)
}

// GetOrderDetail
// @Summary 根据ID获取订单信息
// @Tags 鉴权接口-订单相关方法
// @Param Authorization header string true "Authorization"
// @Param id query string true "获取订单ID"
// @Router /order/detail [get]
func GetOrderDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		helper.ErrorResponse(c, "获取订单信息", fmt.Errorf("ID不能为空"))
		return
	}
	data := new(define.GetOrderDetailReply)
	// 1.获取订单信息
	sysOrder, err := models.GetOrderDetail(id)
	if err != nil {
		helper.ErrorResponse(c, "获取订单信息", err)
		return
	}
	// 赋值
	data.ID = sysOrder.ID
	data.User = sysOrder.User
	data.Food = sysOrder.Food
	data.Num = sysOrder.Num
	data.Remarks = sysOrder.Remarks
	// 返回订单信息
	helper.SuccessResponse(c, "获取订单信息", data)
}

// DeleteOrder
// @Summary 删除订单信息
// @Tags 鉴权接口-订单相关方法
// @Param Authorization header string true "Authorization"
// @Param id path int true "删除订单ID"
// @Router /order/delete/{id} [delete]
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.ErrorResponse(c, "删除订单信息", fmt.Errorf("ID不能为空"))
		return
	}

	if err := models.DB.Where("id = ?", id).Delete(new(models.SysOrder)).Error; err != nil {
		helper.ErrorResponse(c, "删除订单信息", err)
		return
	}
	helper.SuccessResponse(c, "删除订单信息", nil)
}
