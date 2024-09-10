package helper

import (
	"Go_server/config"
	"Go_server/define"
	"archive/zip"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/mail"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 用于处理正确的响应
func SuccessResponse(c *gin.Context, msg string, data interface{}) {
	if data != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": msg + "成功",
			"result":  data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": msg + "成功",
		})
	}
}

// 用于处理错误的响应
func ErrorResponse(c *gin.Context, msg string, err error) {
	if err != nil {
		msg = msg + "失败, err: " + err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": msg,
	})
}

// 生成token
func GenerateToken(uid, rid uint, name string, expireAt int64) (string, error) {
	uc := define.UserClaim{
		UId:  uid,
		RId:  rid,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString(define.Jwtkey)
}

// 解析token
func ValidateToken(tokenString string) (*define.UserClaim, error) {
	if tokenString == "" {
		return nil, errors.New("token不存在")
	}
	// 正常解密Parse，如果自定义就使用ParseWithClaims
	claims := define.UserClaim{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return define.Jwtkey, nil
	})
	if err != nil || token == nil {
		return nil, err
	}
	return &claims, nil
}

// 通过Authorization获取用户信息
func GetAuthorizationUserInfo(authHeader string) (*define.UserClaim, error) {
	if authHeader == "" {
		return nil, errors.New("未登录系统")
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("当前登录已失效请重新登录")
	}
	tokenClaims, err := ValidateToken(parts[1])
	if tokenClaims == nil || err != nil {
		return nil, err
	}
	return tokenClaims, nil
}

// 文件上传功能
func UploadFile(file *multipart.FileHeader, filepath, filename string) (string, error) {
	filepath = "./" + config.Config.UploadDir + filepath
	// 判断文件夹是否存在
	if err := IsDirExists(filepath); err != nil {
		return "", err
	}

	// 文件保存到本地
	fileRead, err := file.Open()
	if err != nil {
		return "", err
	}

	// 在服务器上创建一个文件
	out, err := os.Create(path.Join(filepath, filename))
	if err != nil {
		return "", err
	}
	defer out.Close()
	// 将上传的文件复制到创建的文件中
	io.Copy(out, fileRead)
	return path.Join(filepath, filename), nil
}

// 文件删除
func DeleteFile(filepath string) error {
	if err := os.Remove(filepath); err != nil {
		return err
	}
	return nil
}

// 判断文件夹是否存在
func IsDirExists(dirPath string) error {
	// 使用Stat来检查目录是否存在
	if err := os.MkdirAll(dirPath, os.ModePerm); os.IsNotExist(err) {
		// 目录不存在，所以创建目录，// 设置权限为0755
		if err = os.Mkdir(dirPath, 0755); err != nil {
			// 处理错误，目录无法创建
			return err
		}
	}
	return nil
}

// 判断文件夹是否存在
func IsFileExists(filename string) error {
	_, err := os.Stat(filename)
	return err
}

// 验证电话号码是否正确
func ValidatePhone(phone string) bool {
	// 正则表达式匹配电话号码
	// 这里的正则表达式可以根据实际需要进行调整
	r, _ := regexp.Compile("^1[0-9]{10}$")
	return r.MatchString(phone)
}

// 验证邮箱是否正确
func ValidateEmail(email string) bool {
	if email == "" {
		return true
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

// 获取当前目录下所有文件
func GetAllFile(path string) ([]string, error) {
	// 存放文件名
	filename := []string{}
	// 读取当前目录中的所有文件和子目录
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {
		fd, err := file.Info()
		if err != nil {
			continue
		}
		filename = append(filename, fd.Name())
	}
	return filename, nil
}

// 解压zip文件
func Unzip(srczip, target_dir string) (string, error) {
	zipReader, err := zip.OpenReader(srczip)
	if err != nil {
		return "", err
	}
	defer zipReader.Close()
	for _, file := range zipReader.Reader.File {
		zipped, err := file.Open()
		if err != nil {
			continue
		}
		extracted := filepath.Join(target_dir, file.Name)
		if err := os.MkdirAll(filepath.Dir(extracted), 0755); err != nil {
			continue
		}

		output, err := os.Create(extracted)
		if err != nil {
			continue
		}

		_, err = io.Copy(output, zipped)
		if err != nil {
			continue
		}
		output.Close()
		zipped.Close()
	}
	return target_dir, err
}
