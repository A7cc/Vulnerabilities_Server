package controller

import (
	"Go_server/config"
	"Go_server/helper"
	"Go_server/models"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/gin-gonic/gin"
)

// 测试Golang的原生模板
func CeshiTemplate(c *gin.Context) {
	query := c.Query("query")
	user := &models.SysUser{
		UserName: "admin",
	}
	var text = fmt.Sprintf(`
  <html>
  <head>
  <title>测试Golang原生模板</title>
  </head>
  <body>
    <h2>Hello {{ .UserName }}</h2>
    <p>可以通过{ { .xxxx } }去获取SysUser的一些属性或者是方法，你可以测试一下，下面是搜索的结果：</p>
    <p>%s</p>
  </body></html>
  `, query)
	tmpl := template.New("hello")
	t, err := tmpl.Parse(text)
	if err != nil {
		helper.ErrorResponse(c, "解析模板", err)
		return
	}
	t.Execute(c.Writer, &user)
}

// 测试上传ZIP并解压功能
func UploadZip(c *gin.Context) {
	// 解析表单，获取zip文件
	fh, err := c.FormFile("file")
	if err != nil {
		helper.ErrorResponse(c, "上传ZIP文件", err)
		return
	}
	// 检查文件类型是否为zip
	if ext := filepath.Ext(fh.Filename); ext != ".zip" || fh.Header.Get("Content-Type") != "application/zip" {
		helper.ErrorResponse(c, "上传ZIP文件", fmt.Errorf("文件类型不合规，请上传zip文件"))
		return
	}

	// 上传zip文件
	filepath, err := helper.UploadFile(fh, "zip/", fh.Filename)
	if err != nil {
		helper.ErrorResponse(c, "上传ZIP文件", err)
		return
	}
	// 解压文件到指定目录
	target_dir, err := helper.Unzip(filepath, config.Config.ZipDir)
	if err != nil {
		helper.ErrorResponse(c, "解压ZIP文件", err)
		return
	}
	helper.SuccessResponse(c, "上传并解压", "解压目录为: "+target_dir)
}
