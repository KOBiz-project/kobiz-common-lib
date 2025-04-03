package schemas

import (
	"gorm.io/gorm"
)

type PlanGroup struct {
	gorm.Model
	Telecom       string `gorm:"column:telecom;not null;type:varchar(30);comment:'통신사'" json:"telecom"`
	NamePlanGroup string `gorm:"column:name_plan_group;not null;type:varchar(256);comment:'요금제 그룹명'" json:"namePlanGroup"`
	CodePlanGroup string `gorm:"column:code_plan_group;not null;type:varchar(20);comment:'요금제 그룹 코드 (Raw Data)'" json:"codePlanGroup"`
}

func (PlanGroup) TableName() string {
	return "plan_groups"
}
