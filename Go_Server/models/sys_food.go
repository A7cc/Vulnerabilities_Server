package models

import "gorm.io/gorm"

// 定义食物基本类型数据结构
type SysFood struct {
	gorm.Model
	// 菜名
	FoodName string `gorm:"column:foodname;type:varchar(50);" json:"foodname"`
	// 用户ID，指定外键名称并设置为不可空
	User_id uint `gorm:"column:user_id;not null;" json:"user_id"`
	// 指定关联的外键字段
	User SysUser `gorm:"foreignKey:User_id;" json:"user"`
	// web的图标
	FoodIcon string `gorm:"column:foodicon;type:varchar(100);" json:"foodicon"`
	// 做菜步骤
	FoodProcedure string `gorm:"column:foodprocedure;type:longtext;" json:"foodprocedure"`
	// 视频
	Video string `gorm:"column:video;type:varchar(100);" json:"video"`
	// 价格
	Price float64 `gorm:"column:price;type:float;" json:"price"`
	// 备注
	Remarks string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
}

// 设置食物表名称
func (table *SysFood) TableName() string {
	return "sys_food"
}

// 获取食物数据列表
func GetFoodList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysFood)).Select("sys_food.id,sys_food.foodname,sys_food.price,su.username user,sys_food.user_id,sys_food.foodicon,sys_food.foodprocedure,sys_food.video,sys_food.remarks,sys_food.created_at,sys_food.updated_at").Joins("LEFT JOIN sys_user su ON su.id = sys_food.user_id")
	if keyword != "" {
		tx.Where("sys_food.foodname LIKE '%" + keyword + "%'")
	}
	return tx
}

// 根据ID获取食物信息
func GetFoodDetail(id uint) (*SysFood, error) {
	sf := new(SysFood)
	err := DB.Model(new(SysFood)).Where("id = ?", id).First(sf).Error
	return sf, err
}

// 更新头像
func UpFoodIcon(id uint, filepath string) error {
	err := DB.Model(new(SysFood)).Where("id = ?", id).Updates(map[string]any{
		"foodicon": filepath,
	}).Error
	return err
}

// 更新头像
func UpFoodVideo(id uint, filepath string) error {
	err := DB.Model(new(SysFood)).Where("id = ?", id).Updates(map[string]any{
		"video": filepath,
	}).Error
	return err
}
