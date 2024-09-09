package controller

import (
	"Go_server/define"
	"Go_server/helper"
	"Go_server/models"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"math/rand"

	"github.com/gin-gonic/gin"
)

// GetHomeInfo
// @Summary 获取系统数据信息
// @Param Authorization header string true "Authorization"
// @Tags 鉴权接口-首页设置方法
// @Router /home/get [get]
func GetHomeInfo(c *gin.Context) {
	// 初始化变量
	var (
		usernum   int64
		foodnum   int64
		ordernum  int64
		foodinfos []*define.GetFoodListReply
	)
	err := models.GetFoodList("").Count(&foodnum).Find(&foodinfos).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常",
		})
		return
	}
	if err = models.GetUserList("", -1).Count(&usernum).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常",
		})
		return
	}
	if err = models.GetOrderList("").Count(&ordernum).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "成功",
		"result": struct {
			Usernum   int64                      `json:"usernum"`
			Foodnum   int64                      `json:"foodnum"`
			Ordernum  int64                      `json:"ordernum"`
			FoodInfos []*define.GetFoodListReply `json:"foodinfos"`
		}{
			Usernum:   usernum,
			Foodnum:   foodnum,
			Ordernum:  ordernum,
			FoodInfos: foodinfos,
		},
	})
}

// UpdateInfo
// @Summary 用户自己更新个人信息
// @Tags 鉴权接口-首页设置方法
// @Param Authorization header string true "Authorization"
// @Param UpdateInfoType body define.UpdateInfoType true "用户自己更新个人信息参数"
// @Router /home/updateInfo [put]
func UpdateInfo(c *gin.Context) {
	in := new(define.UpdateInfoType)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常" + err.Error(),
		})
		return
	}
	if in.Username == "" || in.Sex == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名性别不能为空",
		})
		return
	}
	// 获取用户名的基本信息
	getuser, err := models.GetUserDetail(in.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常：" + err.Error(),
		})
		return
	}
	// 判断角色是否已存在
	var cnt int64
	err = models.DB.Model(new(models.SysUser)).Where("id != ? AND username = ?", getuser.ID, in.Username).Count(&cnt).Error
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
	if _, err := os.Stat(in.Avatar); in.Avatar != "" && os.IsNotExist(err) {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": in.Avatar + " 文件不存在",
		})
		return
	}
	// 2.修改数据
	err = models.DB.Model(new(models.SysUser)).Where("id = ?", getuser.ID).Updates(map[string]any{
		"username": in.Username,
		"sex":      in.Sex,
		"avatar":   in.Avatar,
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

// UploadUserAvatar
// @Summary 更新头像
// @Tags 鉴权接口-首页设置方法
// @Param Authorization header string true "Authorization"
// @Param file formData file true "更新头像"
// @Router /home/upuseravatar [post]
func UploadUserAvatar(c *gin.Context) {
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	// 识别图片
	filepath, err := helper.UploadFile(fh, "user/", fh.Filename)
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

// UpdatePwd
// @Summary 更改个人密码
// @Tags 鉴权接口-首页设置方法
// @accept application/x-www-form-urlencoded
// @Param Authorization header string true "Authorization"
// @Param newpwdinfo body string true "用户自己更新密码和UID"
// @Router /home/updatePwd [put]
func UpdatePwd(c *gin.Context) {
	NewPass := c.PostForm("newpass")
	Uid := c.PostForm("uid")
	if NewPass == "" || Uid == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "ID或新密码不能为空",
		})
		return
	}
	Fid, err := strconv.Atoi(Uid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "数据转换异常",
		})
		return
	}
	// 获取用户名的基本信息
	getuser, err := models.GetUserDetail(uint(Fid))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "参数异常：" + err.Error(),
		})
		return
	}
	err = models.DB.Model(new(models.SysUser)).Where("id = ?", getuser.ID).Updates(map[string]any{
		"password": NewPass,
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
		"message": "修改密码成功",
	})
}

// GetSentence
// @Summary 获取每日金句
// @Tags 鉴权接口-首页设置方法
// @Param Authorization header string true "Authorization"
// @Param url query string true "获取每日金句的url" default "http://127.0.0.1:8081/static/sentence/sentence.txt"
// @Router /home/getsentence [get]
func GetSentence(c *gin.Context) {
	url := c.Query("url")
	res, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	strtmp := strings.Replace(string(body), "\r", "", -1)
	stlist := strings.SplitN(strtmp, "\n", -1)
	rand.Seed(time.Now().UnixNano())
	stsrt := stlist[rand.Intn(len(stlist))]
	if stsrt == "" {
		stsrt = "好好吃饭，天天健康。"
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": stsrt,
	})
}
