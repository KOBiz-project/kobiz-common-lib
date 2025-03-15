package schemas

import "gorm.io/gorm"

type TelecomLGUPlanGroup struct {
	gorm.Model
	UrcTrmPpGrpKwrdCd string `gorm:"column:urc_trm_pp_grp_kwrd_cd;type:varchar(2);comment:'그룹 키워드 코드'" json:"urcTrmPpGrpKwrdCd"`
	UrcTrmPpGrpNo     string `gorm:"column:urc_trm_pp_grp_no;type:varchar(2);comment:'그룹 번호'" json:"urcTrmPpGrpNo"`
	TrmPpGrpNm        string `gorm:"column:trm_pp_grp_nm;type:varchar(50);comment:'요금제 그룹 이름'" json:"trmPpGrpNm"`
	// IsLatest: 최신 여부 (기본값 true)
	IsLatest bool `gorm:"column:is_latest; type:bool; default:true; not null; comment:'최신 여부'" json:"isLatest"`
}
