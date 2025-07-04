package schemas

import (
	"gorm.io/gorm"
)

type MemberGroup struct {
	gorm.Model
	NameGroup            string `gorm:"column:name_group;not null;type:varchar(255);comment:'그룹명';uniqueIndex:idx_name_group" json:"nameGroup"`
	CustomAddSubsidyName string `gorm:"column:custom_add_subsidy_name;type:varchar(255);comment:'추가 지원금 별칭';" json:"customAddSubsidyName"`
	InquiryTel           string `gorm:"column:inquiry_tel;type:varchar(255);comment:'문의 번호';" json:"inquiryTel"`
	Memo                 string `gorm:"column:memo;type:text;comment:'메모'" json:"memo"`
}

func (MemberGroup) TableName() string {
	return "member_group"
}
