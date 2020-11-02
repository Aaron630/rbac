package models

import "time"

// Auth table
type Auth struct {
	ID       uint32
	AdminID  uint32     `gorm:"comment:'关联 admin id'"`
	ip       string     `gorm:"size:16"`
	CreateAt *time.Time `gorm:"type:timestamp;default:current_timestamp;comment:'创建时间'"`
}
