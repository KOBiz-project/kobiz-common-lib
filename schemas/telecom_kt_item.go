package schemas

import "gorm.io/gorm"

// type TelecomKTItem struct {
// 	gorm.Model
// 	ProdNo             string  `gorm:"column:prod_no; type:varchar(50); not null" json:"prodNo"`
// 	PetNm              string  `gorm:"column:pet_nm; type:varchar(255); not null" json:"petNm"`
// 	HndsetModelId      string  `gorm:"column:hndset_model_id; type:varchar(50); not null" json:"hndsetModelId"`
// 	HndsetModelNm      string  `gorm:"column:hndset_model_nm; type:varchar(255); not null" json:"hndsetModelNm"`
// 	MakrCd             string  `gorm:"column:makr_cd; type:varchar(10); not null" json:"makrCd"`
// 	ShowOdrg           string  `gorm:"column:show_odrg; type:varchar(10); not null" json:"showOdrg"`
// 	OfwAmt             string  `gorm:"column:ofw_amt; type:varchar(20); not null" json:"ofwAmt"`
// 	StorSuprtAmt       string  `gorm:"column:stor_suprt_amt; type:varchar(20); not null" json:"storSuprtAmt"`
// 	MonthUseChageDcAmt string  `gorm:"column:month_use_chage_dc_amt; type:varchar(20); not null" json:"monthUseChageDcAmt"`
// 	RealAmt            string  `gorm:"column:real_amt; type:varchar(20); not null" json:"realAmt"`
// 	PplId              string  `gorm:"column:ppl_id; type:varchar(50); not null" json:"pplId"`
// 	PplNm              string  `gorm:"column:ppl_nm; type:varchar(255); not null" json:"pplNm"`
// 	PplGroupDivCd      string  `gorm:"column:ppl_group_div_cd; type:varchar(50); not null" json:"pplGroupDivCd"`
// 	HndSetImgNm        string  `gorm:"column:hnd_set_img_nm; type:varchar(255); not null" json:"hndSetImgNm"`
// 	MshopHndSetImgNm   string  `gorm:"column:mshop_hnd_set_img_nm; type:varchar(255); not null" json:"mshopHndSetImgNm"`
// 	SortProd           *string `gorm:"column:sort_prod; type:varchar(50)" json:"sortProd"`
// 	DscnOptnCd         *string `gorm:"column:dscn_optn_cd; type:varchar(50)" json:"dscnOptnCd"`
// 	DispCtgCd          *string `gorm:"column:disp_ctg_cd; type:varchar(50)" json:"dispCtgCd"`
// 	ConvSupotProdYn    *string `gorm:"column:conv_supot_prod_yn; type:char(1)" json:"convSupotProdYn"`
// 	TdKtSuprtAmt       string  `gorm:"column:td_kt_suprt_amt; type:varchar(20); not null" json:"tdKtSuprtAmt"`
// 	TcKtSuprtAmt       *string `gorm:"column:tc_kt_suprt_amt; type:varchar(20)" json:"tcKtSuprtAmt"`
// 	KdKtSuprtAmt       string  `gorm:"column:kd_kt_suprt_amt; type:varchar(20); not null" json:"kdKtSuprtAmt"`
// 	KcKtSuprtAmt       *string `gorm:"column:kc_kt_suprt_amt; type:varchar(20)" json:"kcKtSuprtAmt"`
// 	PrdcCd             *string `gorm:"column:prdc_cd; type:varchar(50)" json:"prdcCd"`
// 	SpnsGrpCd          *string `gorm:"column:spns_grp_cd; type:varchar(50)" json:"spnsGrpCd"`
// 	SpnsMonsType       *string `gorm:"column:spns_mons_type; type:varchar(50)" json:"spnsMonsType"`
// 	StrRow             *string `gorm:"column:str_row; type:varchar(50)" json:"strRow"`
// 	ToRow              *string `gorm:"column:to_row; type:varchar(50)" json:"toRow"`
// 	DeviceType         *string `gorm:"column:device_type; type:varchar(50)" json:"deviceType"`
// 	SpnsrPunoDate      string  `gorm:"column:spnsr_puno_date; type:varchar(20); not null" json:"spnsrPunoDate"`
// 	SaleSttusCd        string  `gorm:"column:sale_sttus_cd; type:varchar(10); not null" json:"saleSttusCd"`
// 	DispYn             string  `gorm:"column:disp_yn; type:char(1); not null" json:"dispYn"`
// 	IsLatest           bool    `gorm:"column:is_latest; type:bool; default:true; not null" json:"isLatest"`
// }

type TelecomKTItem struct {
	gorm.Model
	ProdNo             string  `gorm:"column:prod_no; type:varchar(50); not null" json:"prodNo"`
	PetNm              string  `gorm:"column:pet_nm; type:varchar(255); not null" json:"petNm"`
	HndsetModelId      string  `gorm:"column:hndset_model_id; type:varchar(50); not null" json:"hndsetModelId"`
	HndsetModelNm      string  `gorm:"column:hndset_model_nm; type:varchar(255); not null" json:"hndsetModelNm"`
	MakrCd             string  `gorm:"column:makr_cd; type:varchar(10); not null" json:"makrCd"`
	ShowOdrg           string  `gorm:"column:show_odrg; type:varchar(10); not null" json:"showOdrg"`
	OfwAmt             string  `gorm:"column:ofw_amt; type:varchar(20); not null" json:"ofwAmt"`
	StorSuprtAmt       *string `gorm:"column:stor_suprt_amt; type:varchar(20)" json:"storSuprtAmt"`
	MonthUseChageDcAmt string  `gorm:"column:month_use_chage_dc_amt; type:varchar(20); not null" json:"monthUseChageDcAmt"`
	RealAmt            string  `gorm:"column:real_amt; type:varchar(20); not null" json:"realAmt"`
	PplId              string  `gorm:"column:ppl_id; type:varchar(50); not null" json:"pplId"`
	PplNm              string  `gorm:"column:ppl_nm; type:varchar(255); not null" json:"pplNm"`
	PplGroupDivCd      string  `gorm:"column:ppl_group_div_cd; type:varchar(50); not null" json:"pplGroupDivCd"`
	HndSetImgNm        string  `gorm:"column:hnd_set_img_nm; type:varchar(255); not null" json:"hndSetImgNm"`
	MshopHndSetImgNm   string  `gorm:"column:mshop_hnd_set_img_nm; type:varchar(255); not null" json:"mshopHndSetImgNm"`
	SortProd           *string `gorm:"column:sort_prod; type:varchar(50)" json:"sortProd"`
	DscnOptnCd         *string `gorm:"column:dscn_optn_cd; type:varchar(50)" json:"dscnOptnCd"`
	DispCtgCd          *string `gorm:"column:disp_ctg_cd; type:varchar(50)" json:"dispCtgCd"`
	ConvSupotProdYn    *string `gorm:"column:conv_supot_prod_yn; type:char(1)" json:"convSupotProdYn"`
	KtSuprtAmt         *string `gorm:"column:kt_suprt_amt; type:varchar(20)" json:"ktSuprtAmt"`
	TdKtSuprtAmt       *string `gorm:"column:td_kt_suprt_amt; type:varchar(20)" json:"tdKtSuprtAmt"`
	TcKtSuprtAmt       *string `gorm:"column:tc_kt_suprt_amt; type:varchar(20)" json:"tcKtSuprtAmt"`
	KdKtSuprtAmt       *string `gorm:"column:kd_kt_suprt_amt; type:varchar(20)" json:"kdKtSuprtAmt"`
	KcKtSuprtAmt       *string `gorm:"column:kc_kt_suprt_amt; type:varchar(20)" json:"kcKtSuprtAmt"`
	PrdcCd             *string `gorm:"column:prdc_cd; type:varchar(50)" json:"prdcCd"`
	SpnsGrpCd          *string `gorm:"column:spns_grp_cd; type:varchar(50)" json:"spnsGrpCd"`
	SpnsMonsType       *string `gorm:"column:spns_mons_type; type:varchar(50)" json:"spnsMonsType"`
	StrRow             *string `gorm:"column:str_row; type:varchar(50)" json:"strRow"`
	ToRow              *string `gorm:"column:to_row; type:varchar(50)" json:"toRow"`
	DeviceType         *string `gorm:"column:device_type; type:varchar(50)" json:"deviceType"`
	SpnsrPunoDate      string  `gorm:"column:spnsr_puno_date; type:varchar(20); not null" json:"spnsrPunoDate"`
	SaleSttusCd        string  `gorm:"column:sale_sttus_cd; type:varchar(10); not null" json:"saleSttusCd"`
	DispYn             string  `gorm:"column:disp_yn; type:char(1); not null" json:"dispYn"`
}

func (TelecomKTItem) TableName() string {
	return "telecom_kt_items"
}
