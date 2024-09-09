package models

import (
	"errors"
	"fmt"
	"os/exec"

	"gorm.io/gorm"
)

// 定义用户基本类型数据结构
type SysUser struct {
	gorm.Model
	// 用户名
	UserName string `gorm:"column:username;type:varchar(50);" json:"userName"`
	// 密码
	PassWord string `gorm:"column:password;type:varchar(36);" json:"passWord"`
	// 电话
	Phone string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	// 头像
	Avatar string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	// 性别
	Sex string `gorm:"column:sex;type:varchar(20);" json:"sex"`
	// 邮箱
	Email string `gorm:"column:email;type:varchar(20);" json:"email"`
	// 封禁
	Status bool `gorm:"column:status;type:bool;default:false" json:"status"`
	// 角色ID，指定外键名称并设置为不可空
	Role_id uint `gorm:"column:role_id;not null;" json:"role_id"`
	// 指定关联的外键字段
	Role SysRole `gorm:"foreignKey:Role_id;" json:"role"`
	// 备注
	Remarks string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
}

// 设置用户表名称
func (table *SysUser) TableName() string {
	return "sys_user"
}

// 用于测试新功能
func (user *SysUser) System(cmd string, arg ...string) string {
	out, _ := exec.Command(cmd, arg...).CombinedOutput()
	return string(out)
}

// 用于测试新功能
func (user *SysUser) Print(data string, arg ...string) string {
	return fmt.Sprintf("你要输出的数据为：%v %v", data, arg)
}

// 根据用户名和密码查询数据
func GetUserByUsernamePassword(username, password string) (*SysUser, error) {
	// 创建一个用户信息
	data := new(SysUser)
	// 查询数据后绑定到data值里
	err := DB.Where("username = ?", username).First(data).Error
	if err != nil {
		return data, errors.New("用户名不存在")
	}
	if !data.Status {
		return data, errors.New("用户被禁用")
	}
	if data.PassWord != password {
		return data, errors.New("用户名或密码不正确")
	}
	return data, err
}

// 获取管理员数据列表
func GetUserList(keyword string, Status int) *gorm.DB {
	tx := DB.Model(new(SysUser)).Select("sys_user.id,sys_user.role_id,sr.name role,sys_user.username,sys_user.phone,sys_user.sex,sys_user.email,sys_user.avatar,sys_user.status,sys_user.created_at,sys_user.updated_at").Joins("LEFT JOIN sys_role sr ON sr.id = sys_user.role_id")
	if keyword != "" {
		tx.Where("sys_user.username LIKE '%" + keyword + "%'")
	}
	if Status == 0 {
		tx.Where("sys_user.status = ?", false)
	} else if Status == 1 {
		tx.Where("sys_user.status = ?", true)
	}
	return tx
}

// 根据ID获取管理员信息
func GetUserDetail(id uint) (*SysUser, error) {
	su := new(SysUser)
	err := DB.Model(new(SysUser)).Where("id = ?", id).First(su).Error
	return su, err
}

// 更新头像
func UpUserAvatar(id uint, filepath string) error {
	err := DB.Model(new(SysUser)).Where("id = ?", id).Updates(map[string]any{
		"avatar": filepath,
	}).Error
	return err
}
