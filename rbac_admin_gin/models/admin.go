package models

import (
	"time"
)

// Admin table
type Admin struct {
	ID         uint32
	Account    string     `gorm:"not null;size:32;unique_index;comment:'账号'"`
	Passwd     string     `gorm:"type:char(32);not null;comment:'密码'"`
	Phone      uint64     `gorm:"unique_index;comment:'手机号码'"`
	Avatar     string     `gorm:"size:1024;comment:'头像'"`
	Gender     uint8      `gorm:"not null;default:0;comment:'性别:0 未知 1 男 2 女'"`
	Mail       string     `gorm:"index;size:64;comment:'邮箱'"`
	Status     uint8      `gorm:"not null;default:0;comment:'状态:0启用1关闭'"`
	Token      string     `gorm:"type:char(255);comment:'用户token更新后72小时过期'"`
	CreateAt   *time.Time `gorm:"type:timestamp;default:current_timestamp;comment:'创建时间'"`
	UpdateAt   *time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;comment:'更新时间'"`
	AdminRoles []Role     `gorm:"many2many:admin_role;"`
	Auths      []Auth     `gorm:"foreignkey:AdminID"`
}

// GetUserInfoByAccount func
func (c *Admin) GetUserInfoByAccount(account string) {
	DB.Where("account = ?", account).First(&c)
}

// GetUserInfoByID func
func (c *Admin) GetUserInfoByID(id int) {
	DB.Where("id = ?", id).First(&c)
}
