package schemas

import (
	"gorm.io/gorm"
)

type Reception struct {
	gorm.Model
	NameReception string `gorm:"column:name_reception;not null;type:varchar(256);comment:'개통처명';uniqueIndex:idx_reception_name" json:"nameReception"`
}

func (Reception) TableName() string {
	return "receptions"
}
