package schemas

import (
	"gorm.io/gorm"
)

type Plan struct {
	gorm.Model
	IdPlanGroup int    `gorm:"column:id_plan_group;not null;comment:'요금제 그룹 ID'" json:"idPlanGroup"`
	NamePlan    string `gorm:"column:name_plan;not null;type:varchar(256);comment:'요금제 명'" json:"namePlan"`
	CodePlan    string `gorm:"column:code_plan;not null;type:varchar(20);comment:'요금제 코드 (Raw Data)'" json:"codePlan"`
}

func (Plan) TableName() string {
	return "plans"
}
