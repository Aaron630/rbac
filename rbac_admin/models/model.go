package models

import (
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MysqlConn 连接
var MysqlConn *gorm.DB

type User struct {
	gorm.Model
	MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignkey:UserMemberNumber;association_foreignkey:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number           string
	UserMemberNumber string
}

func init() {
	timezone := "'+08:00'"
	fmt.Println(url.QueryEscape(timezone))
	MysqlConn, err := gorm.Open("mysql", "root:123456@(192.168.218.81:3306)/rbac?charset=utf8mb4&parseTime=True&loc=Local&time_zone="+url.QueryEscape(timezone))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database:")
	}
	// defer MysqlConn.Close()

	// Migrate the schema
	MysqlConn.AutoMigrate(&AdminModule{}, &AdminRole{}, &AdminUser{}, &User{}, &CreditCard{})

	// 创建
	// MysqlConn.Create(&AdminUser{Account: "test1", Passwd: "qqqqq555555", Phone: 13587349346})

	// 读取
	var userModel AdminUser

	MysqlConn.Where("Phone = ?", 13587349346).First(&userModel)
	fmt.Println("userModel:", userModel)

	// MysqlConn.Model(&userModel).Update("Phone", 13136214586)

	// // 删除 - 删除product
	// db.Delete(&product)
}
