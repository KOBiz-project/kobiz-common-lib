package schemas

import (
	"gorm.io/gorm"
)

type MemberGroup struct {
	gorm.Model
	NameGroup            string `gorm:"column:name_group;not null;type:varchar(255);comment:'그룹명';uniqueIndex:idx_name_group" json:"nameGroup"`
	CustomAddSubsidyName string `gorm:"column:custom_add_subsidy_name;type:varchar(255);comment:'추가 지원금 별칭';" json:"customAddSubsidyName"`
}

func (MemberGroup) TableName() string {
	return "member_group"
}
