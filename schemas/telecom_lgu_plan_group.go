package schemas

import "gorm.io/gorm"

type TelecomLGUPlanGroup struct {
	gorm.Model
	UrcTrmPpGrpKwrdCd string `gorm:"column:urc_trm_pp_grp_kwrd_cd;type:varchar(2)" json:"urcTrmPpGrpKwrdCd"`
	UrcTrmPpGrpNo     string `gorm:"column:urc_trm_pp_grp_no;type:varchar(2)" json:"urcTrmPpGrpNo"`
	TrmPpGrpNm        string `gorm:"column:trm_pp_grp_nm;type:varchar(50)" json:"trmPpGrpNm"`
	IsLatest          bool   `gorm:"column:is_latest; type:bool; default:true; not null" json:"isLatest"`
}

func (TelecomLGUPlanGroup) TableName() string {
	return "telecom_lgu_plan_groups"
}
