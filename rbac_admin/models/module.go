package models

// Module table
type Module struct {
	ID       uint32
	ParentID uint8  `gorm:"comment:'父ID'"`
	Action   string `gorm:"size:32;not null;comment:'操作'"`
	Describe string `gorm:"size:32;not null;comment:'功能描述'"`
}
