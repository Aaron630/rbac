package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MysqlConn 连接
var MysqlConn *gorm.DB

func init() {
	MysqlConn, err := gorm.Open("mysql", "root:123456@(192.168.218.144)/rbac?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database:")
	}
	// defer MysqlConn.Close()

	// Migrate the schema
	MysqlConn.AutoMigrate(&AdminModule{})

	// 创建
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	// var product Product
	// db.First(&product, 1) // 查询id为1的product
	// db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	// fmt.Println(product)
	// // 更新 - 更新product的price为2000
	// db.Model(&product).Update("Price", 2000)

	// // 删除 - 删除product
	// db.Delete(&product)
}
