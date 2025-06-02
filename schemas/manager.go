package schemas

import (
	"gorm.io/gorm"
)

type Manager struct {
	gorm.Model
	NameManager string `gorm:"column:name_manager;not null;type:varchar(256);comment:'담당자명';uniqueIndex:idx_manager_name" json:"nameManager"`
}

func (Manager) TableName() string {
	return "managers"
}
