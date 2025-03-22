package schemas

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	NameCompany string `gorm:"column:name_company;not null;type:varchar(256);comment:'제조사 명';uniqueIndex:idx_company_name" json:"nameCompany"`
	ActiveYn    bool   `gorm:"column:active_yn;default:false" json:"activeYn"`
}

func (Company) TableName() string {
	return "companies"
}
