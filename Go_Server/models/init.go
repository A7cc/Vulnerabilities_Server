package models

// models/init.go
import (
	"Go_server/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func NewGormDB() {
	dbConfig := config.Config.Db
	// 连接数据库基本信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Db, dbConfig.Charset)
	// 打开数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 在打开连接时设置日志级别为Info，打印所有sql语句
		Logger: logger.Default.LogMode(logger.Info),
		// 是否禁止自动创建外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	// 自动建表
	if err := db.AutoMigrate(&SysRole{}, &SysUser{}, &SysFood{}, &SysOrder{}); err != nil {
		panic(err)
	}

	// 初始化最原始的角色和用户
	// 判断角色是否存在
	var cnt int64
	if err := db.Model(new(SysRole)).Where("name = ?", "root").Count(&cnt).Error; err != nil {
		panic(err)
	}
	if cnt <= 0 {
		// 创建最高管理员角色
		if err := db.Create(&SysRole{
			Name:    "root",
			Level:   1,
			Remarks: "最高管理员权限",
		}).Error; err != nil {
			panic(err)
		}
	}
	if err := db.Model(new(SysUser)).Where("id = ?", "1").Count(&cnt).Error; err != nil {
		panic(err)
	}
	if cnt <= 0 {
		// 创建初始用户
		err = db.Create(&SysUser{
			UserName: "admin",
			PassWord: "D19e534b_com",
			Phone:    "18888888888",
			Status:   true,
			Role_id:  1,
			Sex:      "男",
			Email:    "123@qq.com",
			Remarks:  "初始管理员",
		}).Error
		if err != nil {
			panic(err)
		}
	}

	if err := db.Model(new(SysRole)).Where("name = ?", "test").Count(&cnt).Error; err != nil {
		panic(err)
	}
	if cnt <= 0 {
		// 创建最高管理员角色
		if err := db.Create(&SysRole{
			Name:    "test",
			Level:   2,
			Remarks: "test权限",
		}).Error; err != nil {
			panic(err)
		}
	}
	if err := db.Model(new(SysUser)).Where("id = ?", "2").Count(&cnt).Error; err != nil {
		panic(err)
	}
	if cnt <= 0 {
		// 创建初始用户
		err = db.Create(&SysUser{
			UserName: "test",
			PassWord: "123456",
			Phone:    "18888888888",
			Status:   true,
			Role_id:  2,
			Sex:      "男",
			Email:    "123@qq.com",
			Remarks:  "测试用户",
		}).Error
		if err != nil {
			panic(err)
		}
	}
	DB = db
}
