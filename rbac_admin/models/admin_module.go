package models

// AdminModule table
type AdminModule struct {
	ID         uint32
	ParentID   uint8       `gorm:"not null"`
	Describe   string      `gorm:"size:32;not null;comment:'功能描述'"`
	Action     string      `gorm:"size:32;not null;comment:'操作'"`
	AdminRoles []AdminRole `gorm:"many2many:admin_role_module"`
}
