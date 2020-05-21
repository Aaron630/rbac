package models

// AdminRole table
type AdminRole struct {
	ID       uint32
	RoleName string `gorm:"size:32;not null;comment:'角色名称'"`
}
