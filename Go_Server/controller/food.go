package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"fmt"
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
	if err := c.ShouldBindQuery(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	// 定义存放数据列表的变量
	var (
		cnt  int64
		list = make([]*define.GetFoodListReply, 0)
	)
	// 获取食物数据列表
	if err := models.GetFoodList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error; err != nil {
		helper.ErrorResponse(c, "获取食物数据列表", err)
		return
	}
	helper.SuccessResponse(c, "获取食物数据列表", gin.H{
		"list":  list,
		"count": cnt,
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
	if err := c.ShouldBindJSON(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	// 定义变量
	var cnt int64
	// 判断菜品是否存在<大于0说明存在菜品
	if err := models.DB.Model(new(models.SysFood)).Where("foodname='" + in.Foodname + "'").Count(&cnt).Error; cnt > 0 && err == nil {
		helper.ErrorResponse(c, "添加菜品", fmt.Errorf("菜品已经存在"))
		return
	}
	// 解密密钥，获取用户信息
	uinfo, err := helper.GetAuthorizationUserInfo(c.Request.Header.Get("Authorization"))
	if err != nil {
		helper.ErrorResponse(c, "解密密钥", err)
		return
	}
	// 获取用户名的基本信息
	getuser, err := models.GetUserDetail(uinfo.UId)
	if err != nil {
		helper.ErrorResponse(c, "获取用户信息", err)
		return
	}
	if in.Price < 0 {
		helper.ErrorResponse(c, "添加菜品", fmt.Errorf("价格不能为负数"))
		return
	}
	if _, err := os.Stat(in.FoodIcon); in.FoodIcon != "" && os.IsNotExist(err) {
		helper.ErrorResponse(c, "添加菜品", err)
		return
	}
	if _, err := os.Stat(in.Video); in.Video != "" && os.IsNotExist(err) {
		helper.ErrorResponse(c, "添加菜品", err)
		return
	}
	// 2.保存数据
	if err := models.DB.Create(&models.SysFood{
		FoodName:      in.Foodname,
		User_id:       getuser.ID,
		FoodIcon:      in.FoodIcon,
		FoodProcedure: in.FoodProcedure,
		Remarks:       in.Remarks,
		Video:         in.Video,
		Price:         in.Price,
	}).Error; err != nil {
		helper.ErrorResponse(c, "添加菜品", err)
		return
	}
	helper.SuccessResponse(c, "添加菜品", nil)
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
		helper.ErrorResponse(c, "获取菜品详情信息", fmt.Errorf("ID不能为空"))
		return
	}
	Fid, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorResponse(c, "获取菜品详情信息", err)
		return
	}
	// 获取菜品信息
	sysFood, err := models.GetFoodDetail(uint(Fid))
	if err != nil {
		helper.ErrorResponse(c, "获取菜品详情信息", err)
		return
	}
	data := new(define.GetFoodDetailReply)
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
	helper.SuccessResponse(c, "获取菜品详情信息", data)
}

// UpdateFood
// @Summary 更新菜品信息
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param UpdateFoodRequest body define.UpdateFoodRequest true "更新菜品信息"
// @Router /food/update [put]
func UpdateFood(c *gin.Context) {
	in := new(define.UpdateFoodRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		helper.ErrorResponse(c, "参数绑定", err)
		return
	}
	// 验证ID是否存在
	if _, err := models.GetFoodDetail(in.ID); err != nil {
		helper.ErrorResponse(c, "更新菜品信息", err)
		return
	}
	// 1.判断菜品是否已存在
	var cnt int64
	if err := models.DB.Model(new(models.SysFood)).Where("id != ? AND foodname = ?", in.ID, in.Foodname).Count(&cnt).Error; err != nil {
		helper.ErrorResponse(c, "更新菜品信息", err)
		return
	}
	if cnt > 0 {
		helper.ErrorResponse(c, "更新菜品信息", fmt.Errorf("菜品已存在"))
		return
	}
	if _, err := os.Stat(in.FoodIcon); in.FoodIcon != "" && os.IsNotExist(err) {
		helper.ErrorResponse(c, "更新菜品信息", err)
		return
	}
	if _, err := os.Stat(in.Video); in.Video != "" && os.IsNotExist(err) {
		helper.ErrorResponse(c, "更新菜品信息", err)
		return
	}
	// 修改数据
	if err := models.DB.Model(new(models.SysFood)).Where("id = ?", in.ID).Updates(map[string]any{
		"foodname":      in.Foodname,
		"foodicon":      in.FoodIcon,
		"foodprocedure": in.FoodProcedure,
		"remarks":       in.Remarks,
		"video":         in.Video,
		"price":         in.Price,
	}).Error; err != nil {
		helper.ErrorResponse(c, "更新菜品信息", err)
		return
	}
	helper.SuccessResponse(c, "更新菜品信息", nil)
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
		helper.ErrorResponse(c, "删除菜品", fmt.Errorf("ID不能为空"))
		return
	}
	Fid, err := strconv.Atoi(id)
	if err != nil {
		helper.ErrorResponse(c, "删除菜品", err)
		return
	}
	// 获取菜品信息
	sysFood, err := models.GetFoodDetail(uint(Fid))
	if err != nil {
		helper.ErrorResponse(c, "删除菜品", err)
		return
	}
	// 删除图片
	if sysFood.FoodIcon != "" {
		if strings.Contains(sysFood.FoodIcon, "..") {
			helper.ErrorResponse(c, "删除菜品", fmt.Errorf("图片文件删除失败"))
			return
		}
		if err = helper.DeleteFile(sysFood.FoodIcon); err != nil {
			helper.ErrorResponse(c, "删除菜品", err)
			return
		}
	}
	// 删除视频
	if sysFood.Video != "" {
		if strings.Contains(sysFood.FoodIcon, "..") {
			helper.ErrorResponse(c, "删除菜品", fmt.Errorf("视频文件删除失败"))
			return
		}
		if err = helper.DeleteFile(sysFood.Video); err != nil {
			helper.ErrorResponse(c, "删除菜品", err)
			return
		}
	}
	if err := models.DB.Unscoped().Where("id = ?", id).Delete(new(models.SysFood)).Error; err != nil {
		helper.ErrorResponse(c, "删除菜品", err)
		return
	}
	helper.SuccessResponse(c, "删除菜品", nil)
}

// UploadFoodIcon
// @Summary 上传菜品ICON
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param file formData file true "上传菜品ICON"
// @Router /food/upfoodicon [post]
func UploadFoodIcon(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		helper.ErrorResponse(c, "上传菜品ICON", err)
		return
	}
	// 识别图片
	filepath, err := helper.UploadFile(fh, "food/", fh.Filename)
	if err != nil {
		helper.ErrorResponse(c, "上传菜品ICON", err)
		return
	}
	helper.SuccessResponse(c, "上传菜品ICON", filepath)
}

// UploadFoodVideo
// @Summary 上传菜品视频
// @Tags 鉴权接口-菜品相关方法
// @Param Authorization header string true "Authorization"
// @Param file formData file true "上传菜品视频"
// @Router /food/upfoodvideo [post]
func UploadFoodVideo(c *gin.Context) {
	//	Context.FormFile的参数为前端上传文件的参数
	fh, err := c.FormFile("file")
	if err != nil {
		helper.ErrorResponse(c, "上传菜品视频", err)
		return
	}
	// 识别视频
	filepath, err := helper.UploadFile(fh, "food/", fh.Filename)
	if err != nil {
		helper.ErrorResponse(c, "上传菜品视频", err)
		return
	}
	helper.SuccessResponse(c, "上传菜品视频", filepath)
}
