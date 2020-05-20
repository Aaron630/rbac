package models

// AdminModule table
type AdminModule struct {
	ID       uint32
	ParentID uint8  `gorm:"type:tinyint;NOT NULL"`
	Describe string `gorm:"size:32;NOT NULL"`
	Action   string `gorm:"size:32;comment:'操作';NOT NULL"`
}
