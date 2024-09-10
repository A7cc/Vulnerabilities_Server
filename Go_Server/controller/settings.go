package controller

import (
	"Go_server/config"
	"Go_server/helper"
	"database/sql"
	"fmt"
	"os/exec"
	"path"
	"runtime"
	"time"

	"github.com/JamesStewy/go-mysqldump"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// Ping
// @Summary 测试连通性
// @Tags 鉴权接口-系统设置相关方法
// @accept application/x-www-form-urlencoded
// @Param Authorization header string true "Authorization"
// @Param addre body string true "IP地址"
// @Router /settings/ping [post]
func Ping(c *gin.Context) {
	ipaddr := c.PostForm("addre")
	// ping对应主机
	var cmd *exec.Cmd
	// 判断操作系统
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "ping "+ipaddr)
	case "linux":
		cmd = exec.Command("bash", "-c", "ping "+ipaddr)
	case "darwin":
		cmd = exec.Command("bash", "-c", "ping "+ipaddr)
	}
	// 执行cmd的命令
	out, err := cmd.CombinedOutput()
	if err != nil {
		helper.ErrorResponse(c, "测试连通性", err)
		return
	}
	output, err := simplifiedchinese.GB18030.NewDecoder().String(string(out)) //转换字符集，解决UTF-8乱码
	if err != nil {
		helper.ErrorResponse(c, "测试连通性", err)
		return
	}
	helper.SuccessResponse(c, "执行", output)
}

// GetBackupsDb
// @Summary 获取备份数据库列表
// @Tags 鉴权接口-系统设置相关方法
// @accept application/x-www-form-urlencoded
// @Param Authorization header string true "Authorization"
// @Param dir body string true "数据库备份目录"
// @Router /settings/getdb [post]
func GetBackupsDb(c *gin.Context) {
	dir := c.PostForm("dir")
	dbNames, err := helper.GetAllFile(dir)
	if err != nil {
		helper.ErrorResponse(c, "获取备份数据库列表", err)
		return
	}
	helper.SuccessResponse(c, "获取备份数据库列表", dbNames)
}

// BackupsDb
// @Summary 备份数据库
// @Tags 鉴权接口-系统设置相关方法
// @Param Authorization header string true "Authorization"
// @Router /settings/backupsdb [get]
func BackupsDb(c *gin.Context) {
	dbConfig := config.Config.Db
	// 打开数据库
	dbbak, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Db, dbConfig.Charset))
	if err != nil {
		helper.ErrorResponse(c, "备份数据库", err)
		return
	}
	dumpdir := config.Config.BackupsdbDir
	if err := helper.IsDirExists(dumpdir); err != nil {
		helper.ErrorResponse(c, "备份数据库", err)
		return
	}
	// 设置备份数据库名字
	bakdbfile := fmt.Sprintf("backup_%v", time.Now().Unix())
	// 备份数据库
	dumper, err := mysqldump.Register(dbbak, dumpdir, bakdbfile)
	if err != nil {
		helper.ErrorResponse(c, "备份数据库", err)
		return
	}
	defer dumper.Close()
	// 备份数据库
	if _, err = dumper.Dump(); err != nil {
		helper.ErrorResponse(c, "备份数据库", err)
		return
	}
	helper.SuccessResponse(c, "备份数据库", nil)
}

// DeleteDb
// @Summary 删除备份数据库
// @Tags 鉴权接口-系统设置相关方法
// @accept application/x-www-form-urlencoded
// @Param Authorization header string true "Authorization"
// @Param dbfile body string true "删除数据库备份文件"
// @Router /settings/deletedb [post]
func DeleteDb(c *gin.Context) {
	dbfile := c.PostForm("dbfile")
	dumpdir := config.Config.BackupsdbDir
	if err := helper.IsDirExists(dumpdir); err != nil {
		helper.ErrorResponse(c, "删除备份数据库", err)
		return
	}
	if err := helper.DeleteFile(path.Join(dumpdir, dbfile)); err != nil {
		helper.ErrorResponse(c, "删除备份数据库", err)
		return
	}
	helper.SuccessResponse(c, "删除"+dbfile+"文件数据库", nil)
}

// DownDb
// @Summary 数据库下载
// @Tags 鉴权接口-系统设置相关方法
// @accept application/x-www-form-urlencoded
// @Param Authorization header string true "Authorization"
// @Param dbfile body string true "下载数据库备份文件"
// @Router /settings/downdb [post]
func DownDb(c *gin.Context) {
	dbfile := c.PostForm("dbfile")
	// 下载文件其实和响应文件差不多
	dumpdir := config.Config.BackupsdbDir
	if err := helper.IsDirExists(dumpdir); err != nil {
		helper.ErrorResponse(c, "下载数据库", err)
		return
	}
	tmppath := path.Join(dumpdir, dbfile)
	if err := helper.IsFileExists(tmppath); err != nil {
		helper.ErrorResponse(c, "下载数据库", err)
		return
	}
	// 类型是文件流，唤起浏览器下载，设置该参数就要设置文件名，如果不设置的话，就会下载默认路由名字
	c.Header("Content-Type", "application/octet-stream")
	// 用来设置下载下来的文件名,
	c.Header("Content-Disposition", "attachment; filename="+dbfile)
	// 下载传输过程中的编码形式，这里是设置为字节
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Transfer-Encoding", "true")
	c.File(tmppath)
}
