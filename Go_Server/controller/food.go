package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetFoodList
// @Summary 获取菜品列表
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param FoodListRequest query define.GetFoodListRequest true "获取菜品列表参数"
// @Router /food/get [get]
func GetFoodList(c *gin.Context) {
	in := &define.GetFoodListRequest{QueryRequest: NewQueryRequest()}
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
		list = make([]*define.GetFoodListReply, 0)
	)
	err = models.GetFoodList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
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

// AddFood
// @Summary 添加菜品
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param AddFoodRequest body define.AddFoodRequest true "接收添加菜品表单数据"
// @Router /food/add [post]
func AddFood(c *gin.Context) {
	in := new(define.AddFoodRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常",
		})
		return
	}
	// 1.判断菜品是否存在
	var cnt int64
	err = models.DB.Model(new(models.SysFood)).Where("foodname='" + in.Foodname + "'").Count(&cnt).Error
	// 大于0说明存在菜品
	if cnt > 0 && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "添加失败，菜品已经存在",
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
	// 获取用户名的基本信息
	getuser, err := models.GetUserDetail(uinfo.UId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常：" + err.Error(),
		})
		return
	}
	if in.Price < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "价格不能为负数",
		})
		return
	}
	if _, err := os.Stat(in.FoodIcon); in.FoodIcon != "" && os.IsNotExist(err) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": in.FoodIcon + " 文件不存在",
		})
		return
	}
	if _, err := os.Stat(in.Video); in.Video != "" && os.IsNotExist(err) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": in.Video + " 文件不存在",
		})
		return
	}
	// 2.保存数据
	err = models.DB.Create(&models.SysFood{
		FoodName:      in.Foodname,
		User_id:       getuser.ID,
		FoodIcon:      in.FoodIcon,
		FoodProcedure: in.FoodProcedure,
		Remarks:       in.Remarks,
		Video:         in.Video,
		Price:         in.Price,
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
		"message": "添加成功",
	})
}

// GetFoodDetail
// @Summary 获取菜品详情信息
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param id query int true "获取菜品ID"
// @Router /food/detail [get]
func GetFoodDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "ID不能为空",
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
	data := new(define.GetFoodDetailReply)
	// 1.获取菜品信息
	sysFood, err := models.GetFoodDetail(uint(Fid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "没有该菜品信息",
		})
		return
	}
	// 赋值
	data.ID = sysFood.ID
	data.Foodname = sysFood.FoodName
	data.FoodIcon = sysFood.FoodIcon
	data.Video = sysFood.Video
	data.User_id = sysFood.User_id
	data.FoodProcedure = sysFood.FoodProcedure
	data.Remarks = sysFood.Remarks
	data.Price = sysFood.Price
	// 返回角色信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取数据成功",
		"result":  data,
	})
}

// UpdateFood
// @Summary 更新菜品信息
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param UpdateFoodRequest body define.UpdateFoodRequest true "更新菜品信息"
// @Router /food/update [put]
func UpdateFood(c *gin.Context) {
	in := new(define.UpdateFoodRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常",
		})
		return
	}
	// 验证ID是否存在
	_, err = models.GetFoodDetail(in.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "该食物不存在",
		})
		return
	}
	// 1.判断菜品是否已存在
	var cnt int64
	err = models.DB.Model(new(models.SysFood)).Where("id != ? AND foodname = ?", in.ID, in.Foodname).Count(&cnt).Error
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
			"message": "更新失败，菜品已经存在",
		})
		return
	}
	if _, err := os.Stat(in.FoodIcon); in.FoodIcon != "" && os.IsNotExist(err) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": in.FoodIcon + " 文件不存在",
		})
		return
	}
	if _, err := os.Stat(in.Video); in.Video != "" && os.IsNotExist(err) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": in.Video + " 文件不存在",
		})
		return
	}
	// 2.修改数据
	err = models.DB.Model(new(models.SysFood)).Where("id = ?", in.ID).Updates(map[string]any{
		"foodname":      in.Foodname,
		"foodicon":      in.FoodIcon,
		"foodprocedure": in.FoodProcedure,
		"remarks":       in.Remarks,
		"video":         in.Video,
		"price":         in.Price,
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

// DeleteFood
// @Summary 删除菜品
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param id path int true "菜品ID"
// @Router /food/delete/{id} [delete]
func DeleteFood(c *gin.Context) {
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
	// 获取菜品信息
	sysFood, err := models.GetFoodDetail(uint(Fid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常" + err.Error(),
		})
		return
	}
	// 删除图片
	if sysFood.FoodIcon != "" {
		if strings.Contains(sysFood.FoodIcon, "..") {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "图片文件删除失败",
			})
			return
		}
		if err = helper.DeleteFile(sysFood.FoodIcon); err != nil {
			fmt.Println("删除图片失败", err)
		}
	}
	// 删除视频
	if sysFood.Video != "" {
		if strings.Contains(sysFood.FoodIcon, "..") {
			c.JSON(http.StatusOK, gin.H{
				"code":    -1,
				"message": "视频文件删除失败",
			})
			return
		}
		if err = helper.DeleteFile(sysFood.Video); err != nil {
			fmt.Println("删除视频失败", err)
		}
	}
	err = models.DB.Unscoped().Where("id = ?", id).Delete(new(models.SysFood)).Error
	if err != nil {
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

// UploadFoodIcon
// @Summary 更新菜品ICON
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param file formData file true "更新菜品ICON"
// @Router /food/upfoodicon [post]
func UploadFoodIcon(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "获取文件信息失败",
		})
		return
	}
	// 识别图片
	filepath, err := helper.UploadFile(fh, "food/", fh.Filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": filepath,
	})
}

// UploadFoodVideo
// @Summary 更新菜品Video
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param file formData file true "更新菜品视频"
// @Router /food/upfoodvideo [post]
func UploadFoodVideo(c *gin.Context) {
	//	Context.FormFile的参数为前端上传文件的参数
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "获取文件信息失败",
		})
		return
	}
	// 识别视频
	filepath, err := helper.UploadFile(fh, "food/", fh.Filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": filepath,
	})
}
