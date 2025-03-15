package schemas

import "gorm.io/gorm"

type TelecomLGUPlanItem struct {
	gorm.Model
	PlanGroupId          uint    `gorm:"column:plan_group_id;comment:'요금제 그룹 ID'" json:"planGroupId"`
	UrcMblPpCd           string  `gorm:"column:urc_mbl_pp_cd;type:varchar(10);comment:'요금제 코드'" json:"urcMblPpCd"`
	UrcMblPpNm           string  `gorm:"column:urc_mbl_pp_nm;type:varchar(50);comment:'요금제 이름'" json:"urcMblPpNm"`
	Mm24ChocAgmtDcntAmt  string  `gorm:"column:mm_24_choc_agmt_dcnt_amt;comment:'24개월 선택 약정 할인 금액'" json:"mm24ChocAgmtDcntAmt"`
	Mm24ChocAgmtDcntTamt string  `gorm:"column:mm_24_choc_agmt_dcnt_tamt;comment:'24개월 선택 약정 총 할인 금액'" json:"mm24ChocAgmtDcntTamt"`
	UrcPpCatgCd          *string `gorm:"column:urc_pp_catg_cd;type:varchar(2);comment:'요금제 카테고리 코드'" json:"urcPpCatgCd,omitempty"`
	UrcPpBasfAmt         string  `gorm:"column:urc_pp_basf_amt;comment:'기본 요금'" json:"urcPpBasfAmt"`
	LastBasfAmt          string  `gorm:"column:last_basf_amt;comment:'최종 요금 (할인 적용 후)'" json:"lastBasfAmt"`
	NagmPpYn             bool    `gorm:"column:nagm_pp_yn;comment:'장기 약정 여부'" json:"nagmPpYn"`
	PpDirtDcntAplyPsblYn bool    `gorm:"column:pp_dirt_dcnt_aply_psbl_yn;comment:'직접 할인 적용 가능 여부'" json:"ppDirtDcntAplyPsblYn"`
	// IsLatest: 최신 여부 (기본값 true)
	IsLatest bool `gorm:"column:is_latest; type:bool; default:true; not null; comment:'최신 여부'" json:"isLatest"`
}
