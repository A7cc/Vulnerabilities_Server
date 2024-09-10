package models

import "gorm.io/gorm"

// 定义角色基本类型数据结构
type SysRole struct {
	gorm.Model
	// 角色名称
	Name string `gorm:"column:name;type:varchar(100);" json:"name"`
	// 角色等级
	Level uint `gorm:"column:level;type:int(11);default:0" json:"level"`
	// 备注
	Remarks string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
}

// 设置角色表名称
func (table *SysRole) TableName() string {
	return "sys_role"
}

// 获取角色数据列表
func GetRoleList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysRole)).Select("id,name,level,created_at,updated_at")
	if keyword != "" {
		tx.Where("name LIKE '%" + keyword + "%'")
	}
	return tx
}

// 根据ID获取角色信息
func GetRoleDetail(id uint) (*SysRole, error) {
	sr := new(SysRole)
	err := DB.Model(new(SysRole)).Where("id = ?", id).First(sr).Error
	return sr, err
}
