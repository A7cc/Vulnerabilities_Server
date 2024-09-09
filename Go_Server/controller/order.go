package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"net/http"

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
	err := c.ShouldBindQuery(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "输入参数数据格式不正确",
		})
		return
	}
	var (
		cnt  int64
		list = make([]*define.GetOrderListReply, 0)
	)
	err = models.GetOrderList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常",
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

// AddOrder
// @Summary 新增订单信息
// @Tags 鉴权接口-订单相关方法
// @Param Authorization header string true "Authorization"
// @Param AddOrderRequest body define.AddOrderRequest true "添加订单信息"
// @Router /order/add [post]
func AddOrder(c *gin.Context) {
	in := new(define.AddOrderRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常" + err.Error(),
		})
		return
	}
	// 1.判断食物是否存在
	var cnt int64
	err = models.DB.Model(new(models.SysFood)).Where("foodname = ?", in.Food).Count(&cnt).Error
	// 大于0说明存在食物
	if cnt <= 0 || err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "订单添加失败，可能不存在该食物",
		})
		return
	}
	// 解密密钥
	// 获取用户名的基本信息
	uinfo, err := helper.GetAuthorizationUserInfo(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常：" + err.Error(),
		})
		return
	}

	// 保存数据
	err = models.DB.Create(&models.SysOrder{
		User:    uinfo.Name,
		Food:    in.Food,
		Num:     in.Num,
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

// GetOrderDetail
// @Summary 根据ID获取订单信息
// @Tags 鉴权接口-订单相关方法
// @Param Authorization header string true "Authorization"
// @Param id query string true "获取订单ID"
// @Router /order/detail [get]
func GetOrderDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "ID不能为空",
		})
		return
	}
	data := new(define.GetOrderDetailReply)
	// 1.获取订单信息
	sysOrder, err := models.GetOrderDetail(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常" + err.Error(),
		})
		return
	}
	// 赋值
	data.ID = sysOrder.ID
	data.User = sysOrder.User
	data.Food = sysOrder.Food
	data.Num = sysOrder.Num
	data.Remarks = sysOrder.Remarks
	// 返回订单信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取数据成功",
		"result":  data,
	})
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
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败，ID不能为空",
		})
		return
	}

	if err := models.DB.Where("id = ?", id).Delete(new(models.SysOrder)).Error; err != nil {
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
