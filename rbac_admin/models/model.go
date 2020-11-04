package models

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 连接
var DB *gorm.DB

func init() {
	timezone := "'+08:00'"
	db, err := gorm.Open(mysql.Open("root:123456@(192.168.137.10:3306)/rbac?charset=utf8mb4&parseTime=True&loc=Local&time_zone="+url.QueryEscape(timezone)), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database:")
	}
	DB = db
	// defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&Admin{}, &Auth{}, &Module{}, &Role{})

	// fmt.Println(languages)
	// 创建
	// DB.Create(&AdminUser{Account: "test1", Passwd: "qqqqq555555", Phone: 13587349346})

	// 读取
	// var userModel AdminUser
	// DB.First(&userModel, 1)
	// DB.Where("Phone = ?", 13587349346).First(&userModel)
	// fmt.Println("userModel:", userModel)

	// DB.Model(&userModel).Update("Phone", 13136214586)

	// // 删除 - 删除product
	// DB.Delete(&product)
}
