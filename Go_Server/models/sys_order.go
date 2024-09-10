package models

import "gorm.io/gorm"

// 定义订单基本类型数据结构
type SysOrder struct {
	gorm.Model
	// 指定关联的外键字段
	User string `gorm:"column:user;type:varchar(50);" json:"user"`
	// 用户ID
	Food string `gorm:"column:food;type:varchar(50);" json:"food"`
	// 数量
	Num uint8 `gorm:"column:num;type:int(11);" json:"num"`
	// 备注
	Remarks string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
}

// 设置订单表名称
func (table *SysOrder) TableName() string {
	return "sys_order"
}

// 获取订单数据列表
func GetOrderList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysOrder)).Select("id,food,user,num,created_at,updated_at")
	if keyword != "" {
		tx.Where("food LIKE '%" + keyword + "%'")
	}
	return tx
}

// 根据ID获取订单信息
func GetOrderDetail(id string) (*SysOrder, error) {
	sr := new(SysOrder)
	err := DB.Model(new(SysOrder)).Where("id = '" + id + "'").First(sr).Error
	return sr, err
}
