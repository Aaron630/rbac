package models

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 连接
var DB *gorm.DB

func init() {
	timezone := "'+08:00'"
	db, err := gorm.Open(
		mysql.Open("root:123456@(192.168.137.10:3306)/rbac?charset=utf8mb4&parseTime=True&loc=Local&time_zone="+url.QueryEscape(timezone)),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database:")
	}
	DB = db
	// defer DB.Close()

	// Migrate the schema
	DB.AutoMigrate(&Admin{}, &Auth{}, &Module{}, &Role{})

	// 初始化数据
	var adminModel Admin
	db.First(&adminModel)
	if adminModel.ID == 0 {
		// *Qq12345678
		DB.Debug().Create([]Admin{
			{
				Account: "admin",
				Passwd:  "5D44EDB067F702C21E95D46D9D787426",
				Phone:   15387349343,
				Email:   "1796419581@qq.com",
				AdminRoles: []Role{
					{
						RoleName: "超级管理员",
						Modules: []Module{
							{
								Action:   "ping",
								Describe: "允许ping",
							},
						},
					},
				},
			},
			{
				Account: "test",
				Passwd:  "5D44EDB067F702C21E95D46D9D787426",
				Phone:   15387349344,
				Email:   "1796419582@qq.com",
			},
		})
	}
}
