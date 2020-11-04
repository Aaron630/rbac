package models

// Role table
type Role struct {
	ID       uint32
	RoleName string `gorm:"size:32;not null;comment:'角色名称'"`
	// Admins   []Admin  `gorm:"many2many:admin_role"`
	Modules []Module `gorm:"many2many:role_module"`
}
