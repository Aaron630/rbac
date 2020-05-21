package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MysqlConn 连接
var MysqlConn *gorm.DB

func init() {
	MysqlConn, err := gorm.Open("mysql", "root:123456@(192.168.218.81:3306)/rbac?charset=utf8mb4&parseTime=True&loc=Local&serverTimezone=UTC")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database:")
	}
	// defer MysqlConn.Close()

	// Migrate the schema
	MysqlConn.AutoMigrate(&AdminModule{}, &AdminRole{}, &AdminUser{})

	// 创建
	// MysqlConn.Create(&AdminUser{Account: "test", Passwd: "qqqqq555555"})

	// 读取
	var userModel AdminUser

	MysqlConn.Where("Phone = ?", 13136214586).First(&userModel)
	fmt.Println("userModel:", userModel)

	// MysqlConn.Model(&userModel).Update("Phone", 13136214586)

	// // 删除 - 删除product
	// db.Delete(&product)
}
