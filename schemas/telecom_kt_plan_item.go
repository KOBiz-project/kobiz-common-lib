package schemas

import "gorm.io/gorm"

type TelecomKTPlanItem struct {
	gorm.Model
	PplGb                string  `gorm:"column:ppl_gb; type:varchar(10); not null" json:"pplGb"`
	CtgNameForMenu       *string `gorm:"column:ctg_name_for_menu; type:varchar(255)" json:"ctgNameForMenu"`
	PplGrpNm             *string `gorm:"column:ppl_grp_nm; type:varchar(255)" json:"pplGrpNm"`
	PplId                string  `gorm:"column:ppl_id; type:varchar(50); not null" json:"pplId"`
	PplNm                string  `gorm:"column:ppl_nm; type:varchar(255); not null" json:"pplNm"`
	PplGrpCd             string  `gorm:"column:ppl_grp_cd; type:varchar(50); not null" json:"pplGrpCd"`
	PplDivCd             *string `gorm:"column:ppl_div_cd; type:varchar(50)" json:"pplDivCd"`
	MonthUseChage        *string `gorm:"column:month_use_chage; type:varchar(50)" json:"monthUseChage"`
	EngtDcAmt            *string `gorm:"column:engt_dc_amt; type:varchar(50)" json:"engtDcAmt"`
	FreeTlkPrvQnt        *string `gorm:"column:free_tlk_prv_qnt; type:varchar(100)" json:"freeTlkPrvQnt"`
	FreeChrPrvQnt        *string `gorm:"column:free_chr_prv_qnt; type:varchar(100)" json:"freeChrPrvQnt"`
	FreeDataPrvQnt       *string `gorm:"column:free_data_prv_qnt; type:varchar(100)" json:"freeDataPrvQnt"`
	OllehMblFreeVcTlkQnt *string `gorm:"column:olleh_mbl_free_vc_tlk_qnt; type:varchar(100)" json:"ollehMblFreeVcTlkQnt"`
	SfznApdDataQnt       *string `gorm:"column:sfzn_apd_data_qnt; type:varchar(100)" json:"sfznApdDataQnt"`
	DataFreeOptnYn       *string `gorm:"column:data_free_optn_yn; type:char(1)" json:"dataFreeOptnYn"`
	ChrTrmOptnYn         *string `gorm:"column:chr_trm_optn_yn; type:char(1)" json:"chrTrmOptnYn"`
	DataCfwdPrvYn        *string `gorm:"column:data_cfwd_prv_yn; type:char(1)" json:"dataCfwdPrvYn"`
	RcmdPplYn            *string `gorm:"column:rcmd_ppl_yn; type:char(1)" json:"rcmdPplYn"`
	DtlDescSbst          *string `gorm:"column:dtl_desc_sbst; type:text" json:"dtlDescSbst"`
	IndcOdrg             *string `gorm:"column:indc_odrg; type:varchar(10)" json:"indcOdrg"`
	UseYn                *string `gorm:"column:use_yn; type:char(1)" json:"useYn"`
	UseDivCd             *string `gorm:"column:use_div_cd; type:varchar(50)" json:"useDivCd"`
	Expnsn1StrVal        *string `gorm:"column:expnsn1_str_val; type:varchar(255)" json:"expnsn1StrVal"`
	OnfrmCd              string  `gorm:"column:onfrm_cd; type:varchar(50); not null" json:"onfrmCd"`
	PplSectCd            *string `gorm:"column:ppl_sect_cd; type:varchar(50)" json:"pplSectCd"`
	PplSectNm            *string `gorm:"column:ppl_sect_nm; type:varchar(255)" json:"pplSectNm"`
	SpnsrTypeCd          *string `gorm:"column:spnsr_type_cd; type:varchar(50)" json:"spnsrTypeCd"`
	TngrPplYn            *string `gorm:"column:tngr_ppl_yn; type:char(1)" json:"tngrPplYn"`
	ChgUseYn             *string `gorm:"column:chg_use_yn; type:char(1)" json:"chgUseYn"`
	SysRegrId            *string `gorm:"column:sys_regr_id; type:varchar(50)" json:"sysRegrId"`
	SysRegDt             *string `gorm:"column:sys_reg_dt; type:varchar(20)" json:"sysRegDt"`
	SysAmdrId            *string `gorm:"column:sys_amdr_id; type:varchar(50)" json:"sysAmdrId"`
	SysAmdDt             *string `gorm:"column:sys_amd_dt; type:varchar(20)" json:"sysAmdDt"`
	CombDcPplCd          *string `gorm:"column:comb_dc_ppl_cd; type:varchar(50)" json:"combDcPplCd"`
	ChageDispNo          string  `gorm:"column:chage_disp_no; type:varchar(10); not null" json:"chageDispNo"`
	DataBasic            string  `gorm:"column:data_basic; type:varchar(255); not null" json:"dataBasic"`
	DataBenefit          *string `gorm:"column:data_benefit; type:varchar(255)" json:"dataBenefit"`
	TlkBasic             string  `gorm:"column:tlk_basic; type:varchar(255); not null" json:"tlkBasic"`
	CharBasic            string  `gorm:"column:char_basic; type:varchar(255); not null" json:"charBasic"`
	TlkBenefit           *string `gorm:"column:tlk_benefit; type:varchar(255)" json:"tlkBenefit"`
	ChageNote            string  `gorm:"column:chage_note; type:text; not null" json:"chageNote"`
	PplDcUseYn           *string `gorm:"column:ppl_dc_use_yn; type:char(1)" json:"pplDcUseYn"`
	OnlineFrmUseYn       *string `gorm:"column:online_frm_use_yn; type:char(1)" json:"onlineFrmUseYn"`
	LimitPplApyYn        *string `gorm:"column:limit_ppl_apy_yn; type:char(1)" json:"limitPplApyYn"`
	PrmtnDcAmt           *string `gorm:"column:prmtn_dc_amt; type:varchar(50)" json:"prmtnDcAmt"`
	UnityKeyword         *string `gorm:"column:unity_keyword; type:varchar(255)" json:"unityKeyword"`
	BnfBdgTxt            *string `gorm:"column:bnf_bdg_txt; type:varchar(255)" json:"bnfBdgTxt"`
	BnfBdgClr            *string `gorm:"column:bnf_bdg_clr; type:varchar(20)" json:"bnfBdgClr"`
	TxtBnfDesc           string  `gorm:"column:txt_bnf_desc; type:text; not null" json:"txtBnfDesc"`
	TxtBnfTltpDispYn     *string `gorm:"column:txt_bnf_tltp_disp_yn; type:char(1)" json:"txtBnfTltpDispYn"`
	TxtBnfTltpDesc       *string `gorm:"column:txt_bnf_tltp_desc; type:text" json:"txtBnfTltpDesc"`
	PplGroupSortCd       *string `gorm:"column:ppl_group_sort_cd; type:varchar(50)" json:"pplGroupSortCd"`
	ProdAdtnDefVal       *string `gorm:"column:prod_adtn_def_val; type:varchar(255)" json:"prodAdtnDefVal"`
	SeniorYn             *string `gorm:"column:senior_yn; type:char(1)" json:"seniorYn"`
	DailyReductionCost   *string `gorm:"column:daily_reduction_cost; type:varchar(50)" json:"dailyReductionCost"`
	VoiceCallCost        *string `gorm:"column:voice_call_cost; type:varchar(50)" json:"voiceCallCost"`
	VideoCallCost        *string `gorm:"column:video_call_cost; type:varchar(50)" json:"videoCallCost"`
	ShortTextCost        *string `gorm:"column:short_text_cost; type:varchar(50)" json:"shortTextCost"`
	LongTextCost         *string `gorm:"column:long_text_cost; type:varchar(50)" json:"longTextCost"`
	MmsTextCost          *string `gorm:"column:mms_text_cost; type:varchar(50)" json:"mmsTextCost"`
	CostPerPacket        *string `gorm:"column:cost_per_packet; type:varchar(50)" json:"costPerPacket"`
	CostPerPacketNormal  *string `gorm:"column:cost_per_packet_normal; type:varchar(50)" json:"costPerPacketNormal"`
	MTxtBnfDesc          string  `gorm:"column:m_txt_bnf_desc; type:text; not null" json:"mTxtBnfDesc"`
	PrsnlShare           string  `gorm:"column:prsnl_share; type:varchar(255); not null" json:"prsnlShare"`
	ROAMING              string  `gorm:"column:roaming; type:varchar(255); not null" json:"rOAMING"` // 수정: 12550 -> 255
	PrmBnMImgPath        *string `gorm:"column:prm_bn_m_img_path; type:varchar(255)" json:"prmBnMImgPath"`
	PrmBnMImgName        *string `gorm:"column:prm_bn_m_img_name; type:varchar(255)" json:"prmBnMImgName"`
	PunoMonthUseChage    string  `gorm:"column:puno_month_use_chage; type:varchar(50); not null" json:"punoMonthUseChage"`
	PunoMonthUseDcChage  string  `gorm:"column:puno_month_use_dc_chage; type:varchar(50); not null" json:"punoMonthUseDcChage"`
	IsLatest             bool    `gorm:"column:is_latest; type:bool; default:true; not null" json:"isLatest"`
}

func (TelecomKTPlanItem) TableName() string {
	return "telecom_kt_plan_items"
}
