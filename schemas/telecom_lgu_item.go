package schemas

import "gorm.io/gorm"

// LGUItem 구조체 정의 (MySQL에 맞게 작성)
// GORM 모델을 사용해 MySQL 데이터베이스의 테이블과 매핑
type TelecomLGUItem struct {
	gorm.Model // GORM 기본 필드: ID, CreatedAt, UpdatedAt, DeletedAt
	// UrcTrmMdlCd: 단말기 모델 코드
	UrcTrmMdlCd string `gorm:"column:urc_trm_mdl_cd; type:varchar(50); not null; comment:'단말기 모델 코드'" json:"urcTrmMdlCd"`
	// UrcTrmMdlNm: 단말기 모델 이름
	UrcTrmMdlNm string `gorm:"column:urc_trm_mdl_nm; type:varchar(255); not null; comment:'단말기 모델 이름'" json:"urcTrmMdlNm"`
	// PcUsgListImgeUrlAddr: 이미지 URL 주소
	PcUsgListImgeUrlAddr string `gorm:"column:pc_usg_list_imge_url_addr; type:varchar(1000); not null; comment:'이미지 URL 주소'" json:"pcUsgListImgeUrlAddr"`
	// RlCoutDttm: 출시 날짜
	RlCoutDttm string `gorm:"column:rl_cout_dttm; type:varchar(20); not null; comment:'출시 날짜'" json:"rlCoutDttm"`
	// DlvrPrc: 출고가
	DlvrPrc int `gorm:"column:dlvr_prc; type:int; not null; comment:'출고가'" json:"dlvrPrc"`
	// SixPlanPuanSuptAmt: 6플랜 공시지원금
	SixPlanPuanSuptAmt int `gorm:"column:six_plan_puan_supt_amt; type:int; not null; comment:'6플랜 공시지원금'" json:"sixPlanPuanSuptAmt"`
	// SixPlanAddSuptAmt: 6플랜 추가지원금
	SixPlanAddSuptAmt int `gorm:"column:six_plan_add_supt_amt; type:int; not null; comment:'6플랜 추가지원금'" json:"sixPlanAddSuptAmt"`
	// SixPlanCvrtSuptAmt: 6플랜 전환지원금
	SixPlanCvrtSuptAmt int `gorm:"column:six_plan_cvrt_supt_amt; type:int; not null; comment:'6플랜 전환지원금'" json:"sixPlanCvrtSuptAmt"`
	// SixPlanAddCvrtSuptAmt: 6플랜 추가전환지원금
	SixPlanAddCvrtSuptAmt int `gorm:"column:six_plan_add_cvrt_supt_amt; type:int; not null; comment:'6플랜 추가전환지원금'" json:"sixPlanAddCvrtSuptAmt"`
	// SixPlanSuptTamt: 6플랜 총 지원금
	SixPlanSuptTamt int `gorm:"column:six_plan_supt_tamt; type:int; not null; comment:'6플랜 총 지원금'" json:"sixPlanSuptTamt"`
	// SixpPanBuyPrc: 6플랜 구매가격
	SixpPanBuyPrc int `gorm:"column:sixp_pan_buy_prc; type:int; not null; comment:'6플랜 구매가격'" json:"sixpPanBuyPrc"`
	// BasicPlanPuanSuptAmt: 기본 플랜 공시지원금
	BasicPlanPuanSuptAmt int `gorm:"column:basic_plan_puan_supt_amt; type:int; not null; comment:'기본 플랜 공시지원금'" json:"basicPlanPuanSuptAmt"`
	// BasicPlanAddSuptAmt: 기본 플랜 추가지원금
	BasicPlanAddSuptAmt int `gorm:"column:basic_plan_add_supt_amt; type:int; not null; comment:'기본 플랜 추가지원금'" json:"basicPlanAddSuptAmt"`
	// BasicPlanCvrtSuptAmt: 기본 플랜 전환지원금
	BasicPlanCvrtSuptAmt int `gorm:"column:basic_plan_cvrt_supt_amt; type:int; not null; comment:'기본 플랜 전환지원금'" json:"basicPlanCvrtSuptAmt"`
	// BasicPlanAddCvrtSuptAmt: 기본 플랜 추가전환지원금
	BasicPlanAddCvrtSuptAmt int `gorm:"column:basic_plan_add_cvrt_supt_amt; type:int; not null; comment:'기본 플랜 추가전환지원금'" json:"basicPlanAddCvrtSuptAmt"`
	// BasicPlanSuptTamt: 기본 플랜 총 지원금
	BasicPlanSuptTamt int `gorm:"column:basic_plan_supt_tamt; type:int; not null; comment:'기본 플랜 총 지원금'" json:"basicPlanSuptTamt"`
	// BasicPlanBuyPrc: 기본 플랜 구매가격
	BasicPlanBuyPrc int `gorm:"column:basic_plan_buy_prc; type:int; not null; comment:'기본 플랜 구매가격'" json:"basicPlanBuyPrc"`
	// DvicManfEngNm: 제조사 영문명
	DvicManfEngNm string `gorm:"column:dvic_manf_eng_nm; type:varchar(50); not null; comment:'제조사 영문명'" json:"dvicManfEngNm"`
	// MdlbSufuGuidCntn: 모델 안내 내용
	MdlbSufuGuidCntn string `gorm:"column:mdlb_sufu_guid_cntn; type:text; comment:'모델 안내 내용'" json:"mdlbSufuGuidCntn"`
	// OnlnOrdrPsblYn: 온라인 주문 가능 여부
	OnlnOrdrPsblYn string `gorm:"column:onln_ordr_psbl_yn; type:char(1); not null; comment:'온라인 주문 가능 여부'" json:"onlnOrdrPsblYn"`
	// TrmMdlEposEngNm: 단말기 모델 영문 노출명
	TrmMdlEposEngNm string `gorm:"column:trm_mdl_epos_eng_nm; type:varchar(255); not null; comment:'단말기 모델 영문 노출명'" json:"trmMdlEposEngNm"`
	// UrcTrmKndEngNm: 단말기 종류 영문명
	UrcTrmKndEngNm *string `gorm:"column:urc_trm_knd_eng_nm; type:varchar(50); comment:'단말기 종류 영문명'" json:"urcTrmKndEngNm"`
	// IsLatest: 최신 여부 (기본값 true)
	IsLatest bool `gorm:"column:is_latest; type:bool; default:true; not null; comment:'최신 여부'" json:"isLatest"`
}

// TableName 커스터마이징 (MySQL 테이블 이름 지정)
func (TelecomLGUItem) TableName() string {
	return "telecom_lgu_items"
}
