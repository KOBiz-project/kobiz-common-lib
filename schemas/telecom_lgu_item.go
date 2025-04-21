package schemas

import "gorm.io/gorm"

type TelecomLGUItem struct {
	gorm.Model
	UrcTrmMdlCd             string  `gorm:"column:urc_trm_mdl_cd; type:varchar(50); not null" json:"urcTrmMdlCd"`
	UrcTrmMdlNm             string  `gorm:"column:urc_trm_mdl_nm; type:varchar(255); not null" json:"urcTrmMdlNm"`
	PcUsgListImgeUrlAddr    string  `gorm:"column:pc_usg_list_imge_url_addr; type:varchar(1000); not null" json:"pcUsgListImgeUrlAddr"`
	RlCoutDttm              string  `gorm:"column:rl_cout_dttm; type:varchar(20); not null" json:"rlCoutDttm"`
	DlvrPrc                 int     `gorm:"column:dlvr_prc; type:int; not null" json:"dlvrPrc"`
	UrcMblPpCd              string  `gorm:"column:urc_mbl_pp_cd;type:varchar(10)" json:"urcMblPpCd"`
	SixPlanPuanSuptAmt      int     `gorm:"column:six_plan_puan_supt_amt; type:int; not null" json:"sixPlanPuanSuptAmt"`
	SixPlanAddSuptAmt       int     `gorm:"column:six_plan_add_supt_amt; type:int; not null" json:"sixPlanAddSuptAmt"`
	SixPlanCvrtSuptAmt      int     `gorm:"column:six_plan_cvrt_supt_amt; type:int; not null" json:"sixPlanCvrtSuptAmt"`
	SixPlanAddCvrtSuptAmt   int     `gorm:"column:six_plan_add_cvrt_supt_amt; type:int; not null" json:"sixPlanAddCvrtSuptAmt"`
	SixPlanSuptTamt         int     `gorm:"column:six_plan_supt_tamt; type:int; not null" json:"sixPlanSuptTamt"`
	SixpPanBuyPrc           int     `gorm:"column:sixp_pan_buy_prc; type:int; not null" json:"sixpPanBuyPrc"`
	BasicPlanPuanSuptAmt    int     `gorm:"column:basic_plan_puan_supt_amt; type:int; not null" json:"basicPlanPuanSuptAmt"`
	BasicPlanAddSuptAmt     int     `gorm:"column:basic_plan_add_supt_amt; type:int; not null" json:"basicPlanAddSuptAmt"`
	BasicPlanCvrtSuptAmt    int     `gorm:"column:basic_plan_cvrt_supt_amt; type:int; not null" json:"basicPlanCvrtSuptAmt"`
	BasicPlanAddCvrtSuptAmt int     `gorm:"column:basic_plan_add_cvrt_supt_amt; type:int; not null" json:"basicPlanAddCvrtSuptAmt"`
	BasicPlanSuptTamt       int     `gorm:"column:basic_plan_supt_tamt; type:int; not null" json:"basicPlanSuptTamt"`
	BasicPlanBuyPrc         int     `gorm:"column:basic_plan_buy_prc; type:int; not null" json:"basicPlanBuyPrc"`
	DvicManfEngNm           string  `gorm:"column:dvic_manf_eng_nm; type:varchar(50); not null" json:"dvicManfEngNm"`
	MdlbSufuGuidCntn        string  `gorm:"column:mdlb_sufu_guid_cntn; type:text" json:"mdlbSufuGuidCntn"`
	OnlnOrdrPsblYn          string  `gorm:"column:onln_ordr_psbl_yn; type:char(1); not null" json:"onlnOrdrPsblYn"`
	TrmMdlEposEngNm         string  `gorm:"column:trm_mdl_epos_eng_nm; type:varchar(255); not null" json:"trmMdlEposEngNm"`
	UrcTrmKndEngNm          *string `gorm:"column:urc_trm_knd_eng_nm; type:varchar(50)" json:"urcTrmKndEngNm"`
	IsLatest                bool    `gorm:"column:is_latest; type:bool; default:true; not null" json:"isLatest"`
}

func (TelecomLGUItem) TableName() string {
	return "telecom_lgu_items"
}
