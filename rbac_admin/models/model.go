package models

import (
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 连接
var DB *gorm.DB

func init() {
	timezone := "'+08:00'"
	DB, err := gorm.Open("mysql", "root:123456@(192.168.43.126:3306)/rbac?charset=utf8mb4&parseTime=True&loc=Local&time_zone="+url.QueryEscape(timezone))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database:")
	}
	// defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&AdminModule{}, &AdminRole{}, &AdminUser{})

	// 创建
	// DB.Create(&AdminUser{Account: "test1", Passwd: "qqqqq555555", Phone: 13587349346})

	// 读取
	// var userModel AdminUser

	// DB.Where("Phone = ?", 13587349346).First(&userModel)
	// fmt.Println("userModel:", userModel)

	// DB.Model(&userModel).Update("Phone", 13136214586)

	// // 删除 - 删除product
	// db.Delete(&product)
}
