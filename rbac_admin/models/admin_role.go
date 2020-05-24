package models

// AdminRole table
type AdminRole struct {
	ID           uint32
	RoleName     string        `gorm:"size:32;not null;comment:'角色名称'"`
	AdminUsers   []AdminUser   `gorm:"many2many:admin_user_role"`
	AdminModuleS []AdminModule `gorm:"many2many:admin_role_module"`
}
