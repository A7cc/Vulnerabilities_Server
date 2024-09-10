package define

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	// jwt的key，密钥
	Jwtkey = []byte("")
	// token的有效期，7天
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	// 刷新token有效期，14天
	RefreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
	// 默认分页没有显示条数
	DefaultSize = 10
)

// 定义JWT token中所包含的信息
type UserClaim struct {
	// id
	UId uint
	// 角色ID
	RId uint
	// 用户名
	Name string
	// JWT 的标准声明，包含了 JWT 的一些基本信息
	jwt.StandardClaims
}
