package models

// Module table
type Module struct {
	ID       uint32
	ParentID uint8  `gorm:"not null"`
	Describe string `gorm:"size:32;not null;comment:'功能描述'"`
	Action   string `gorm:"size:32;not null;comment:'操作'"`
	Roles    []Role `gorm:"many2many:role_module"`
}
