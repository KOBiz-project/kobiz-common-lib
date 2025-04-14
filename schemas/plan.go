package schemas

import (
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	Telecom    string `gorm:"column:telecom;not null;type:varchar(30);comment:'통신사';uniqueIndex:idx_plan_name" json:"telecom"`
	NamePlan   string `gorm:"column:name_plan;not null;type:varchar(256);comment:'요금제명';uniqueIndex:idx_plan_name" json:"namePlan"`
	CodePlan   string `gorm:"column:code_plan;not null;type:varchar(20);comment:'요금제 코드 (Raw Data)'" json:"codePlan"`
	BasicPrice int64  `gorm:"column:basic_price;not null;type:bigint;comment:'기본요금 (Raw Data)'" json:"basicPrice"`
	ActiveYn   bool   `gorm:"column:active_yn;default:false" json:"activeYn"`
}

func (Plan) TableName() string {
	return "plans"
}
