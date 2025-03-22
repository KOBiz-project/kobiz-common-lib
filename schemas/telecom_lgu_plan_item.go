package schemas

import "gorm.io/gorm"

type TelecomLGUPlanItem struct {
	gorm.Model
	PlanGroupId          uint    `gorm:"column:plan_group_id" json:"planGroupId"`
	UrcMblPpCd           string  `gorm:"column:urc_mbl_pp_cd;type:varchar(10)" json:"urcMblPpCd"`
	UrcMblPpNm           string  `gorm:"column:urc_mbl_pp_nm;type:varchar(50)" json:"urcMblPpNm"`
	Mm24ChocAgmtDcntAmt  string  `gorm:"column:mm_24_choc_agmt_dcnt_amt" json:"mm24ChocAgmtDcntAmt"`
	Mm24ChocAgmtDcntTamt string  `gorm:"column:mm_24_choc_agmt_dcnt_tamt" json:"mm24ChocAgmtDcntTamt"`
	UrcPpCatgCd          *string `gorm:"column:urc_pp_catg_cd;type:varchar(2)" json:"urcPpCatgCd,omitempty"`
	UrcPpBasfAmt         string  `gorm:"column:urc_pp_basf_amt" json:"urcPpBasfAmt"`
	LastBasfAmt          string  `gorm:"column:last_basf_amt" json:"lastBasfAmt"`
	NagmPpYn             bool    `gorm:"column:nagm_pp_yn" json:"nagmPpYn"`
	PpDirtDcntAplyPsblYn bool    `gorm:"column:pp_dirt_dcnt_aply_psbl_yn" json:"ppDirtDcntAplyPsblYn"`
	IsLatest             bool    `gorm:"column:is_latest; type:bool; default:true; not null" json:"isLatest"`
}

func (TelecomLGUPlanItem) TableName() string {
	return "telecom_lgu_plan_items"
}
