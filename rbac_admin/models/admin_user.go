package models

import (
	"time"
)

// AdminUser table
type AdminUser struct {
	ID       uint32
	Account  string     `gorm:"not null;size:32;comment:'账号'"`
	Passwd   string     `gorm:"type:char(32);not null;comment:'密码'"`
	Phone    uint64     `gorm:"comment:'手机号码'"`
	Avatar   string     `gorm:"size:1024;comment:'头像'"`
	Gender   uint8      `gorm:"not null;default:0;comment:'性别:0 未知 1 男 2 女'"`
	Mail     string     `gorm:"size:64;comment:'邮箱'"`
	Status   uint8      `gorm:"not null;default:0;comment:'状态:0启用1关闭'"`
	CreateAt *time.Time `gorm:"type:timestamp;default:current_timestamp;comment:'创建时间'"`
	UpdateAt *time.Time `gorm:"type:timestamp;default:current_timestamp on update current_timestamp;comment:'更新时间'"`
}
