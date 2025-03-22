package schemas

import "gorm.io/gorm"

type TelecomSKPlanGroup struct {
	gorm.Model
	CategoryID       string `gorm:"column:category_id;type:varchar(20)" json:"categoryId"`
	CategoryNm       string `gorm:"column:category_nm;type:varchar(100)" json:"categoryNm"`
	BasicChargeMin   string `gorm:"column:basic_charge_min;type:varchar(20)" json:"basicChargeMin"`
	BasicChargeMax   string `gorm:"column:basic_charge_max;type:varchar(20)" json:"basicChargeMax"`
	BasicCharge      string `gorm:"column:basic_charge;type:varchar(20)" json:"basicCharge"`
	SubscriptionNm   string `gorm:"column:subscription_nm;type:varchar(100)" json:"subscriptionNm"`
	SubscriptionID   string `gorm:"column:subscription_id;type:varchar(20);primary_key" json:"subscriptionId"`
	FlagCheck        string `gorm:"column:flag_check;type:varchar(10)" json:"flagCheck"`
	SubLimit         string `gorm:"column:sub_limit;type:varchar(50)" json:"subLimit"`
	DataInfo         string `gorm:"column:data_info;type:varchar(100)" json:"dataInfo"`
	VoiceInfo        string `gorm:"column:voice_info;type:varchar(100)" json:"voiceInfo"`
	SMSInfo          string `gorm:"column:sms_info;type:varchar(100)" json:"smsInfo"`
	NetTyp           string `gorm:"column:net_typ;type:varchar(10)" json:"netTyp"`
	SubOption        string `gorm:"column:sub_option;type:varchar(200)" json:"subOption"`
	ModelNm          string `gorm:"column:model_nm;type:varchar(100)" json:"modelNm"`
	RcProdID         string `gorm:"column:rc_prod_id;type:varchar(20)" json:"rcProdId"`
	ZsvcFeeProdID    string `gorm:"column:zsvc_fee_prod_id;type:varchar(20)" json:"zsvcFeeProdId"`
	BeqpEqpMdlNm     string `gorm:"column:beqp_eqp_mdl_nm;type:varchar(100)" json:"beqpEqpMdlNm"`
	UserSesnID       string `gorm:"column:user_sesn_id;type:varchar(50)" json:"userSesnId"`
	BuyBrwsClCd      string `gorm:"column:buy_brws_cl_cd;type:varchar(20)" json:"buyBrwsClCd"`
	BuyPprocStepCd   string `gorm:"column:buy_pproc_step_cd;type:varchar(20)" json:"buyPprocStepCd"`
	BuyBrwsProdCtt   string `gorm:"column:buy_brws_prod_ctt;type:varchar(200)" json:"buyBrwsProdCtt"`
	ProductGrpNm     string `gorm:"column:product_grp_nm;type:varchar(100)" json:"productGrpNm"`
	ProductGrpID     string `gorm:"column:product_grp_id;type:varchar(20)" json:"productGrpId"`
	ChnlMbrID        string `gorm:"column:chnl_mbr_id;type:varchar(50)" json:"chnlMbrId"`
	UsedFlag         string `gorm:"column:used_flag;type:varchar(10)" json:"usedFlag"`
	OptSelMndtCnt    string `gorm:"column:opt_sel_mndt_cnt;type:varchar(10)" json:"optSelMndtCnt"`
	DisplayYn        string `gorm:"column:display_yn;type:varchar(2)" json:"displayYn"`
	IcasProdID       string `gorm:"column:icas_prod_id;type:varchar(20)" json:"icasProdId"`
	DplanRecommendID string `gorm:"column:dplan_recommend_id;type:varchar(20)" json:"dplanRecommendId"`
	SubcommNm        string `gorm:"column:subcomm_nm;type:varchar(100)" json:"subcommNm"`
	SubcommDisamt    string `gorm:"column:subcomm_disamt;type:varchar(20)" json:"subcommDisamt"`
	SelSubcommYn     string `gorm:"column:sel_subcomm_yn;type:varchar(2)" json:"selSubcommYn"`
	SelSubcommDisamt string `gorm:"column:sel_subcomm_disamt;type:varchar(20)" json:"selSubcommDisamt"`
	FmlyNoCtgryFl    string `gorm:"column:fmly_no_ctgry_fl;type:varchar(2)" json:"fmlyNoCtgryFl"`
	DataDtlPhrs      string `gorm:"column:data_dtl_phrs;type:varchar(200)" json:"dataDtlPhrs"`
	TcDefaultQty     string `gorm:"column:tc_default_qty;type:varchar(100)" json:"tcDefaultQty"`
	SmsDefaultQty    string `gorm:"column:sms_default_qty;type:varchar(100)" json:"smsDefaultQty"`
	DataDefaultQty   string `gorm:"column:data_default_qty;type:varchar(100)" json:"dataDefaultQty"`
	SubcategoryID    string `gorm:"column:subcategory_id;type:varchar(10)" json:"subcategoryId"`
	BenefitInfo      string `gorm:"column:benefit_info;type:varchar(500)" json:"benefitInfo"`
	SubcommTerm      string `gorm:"column:subcomm_term;type:varchar(10)" json:"subcommTerm"`
	SubcommNoYn      string `gorm:"column:subcomm_no_yn;type:varchar(2)" json:"subcommNoYn"`
	IsLatest         bool   `gorm:"column:is_latest;type:bool;default:true;not null" json:"isLatest"`
}

func (TelecomSKPlanGroup) TableName() string {
	return "telecom_sk_plan_groups"
}

type TelecomSKPlanItem struct {
	gorm.Model
	SubOptID    string `gorm:"column:sub_opt_id;type:varchar(20);uniqueIndex" json:"subOptId"`
	SubOptNm    string `gorm:"column:sub_opt_nm;type:varchar(100)" json:"subOptNm"`
	SubOptTag   string `gorm:"column:sub_opt_tag;type:varchar(200)" json:"subOptTag"`
	BasSelYn    string `gorm:"column:bas_sel_yn;type:varchar(2)" json:"basSelYn"`
	OptYn       string `gorm:"column:opt_yn;type:varchar(2)" json:"optYn"`
	BenfDtlDesc string `gorm:"column:benf_dtl_desc;type:varchar(500)" json:"benfDtlDesc"`
	BenfImgPath string `gorm:"column:benf_img_path;type:varchar(200)" json:"benfImgPath"`
	BenfImgDesc string `gorm:"column:benf_img_desc;type:varchar(200)" json:"benfImgDesc"`
	IsLatest    bool   `gorm:"column:is_latest;type:bool;default:true;not null" json:"isLatest"`
}

func (TelecomSKPlanItem) TableName() string {
	return "telecom_sk_plan_items"
}

type TelecomSKPlanRelation struct {
	gorm.Model
	PlanGroupId uint `gorm:"column:plan_group_id;uniqueIndex:idx_plan_group_item" json:"planGroupId"`
	PlanItemId  uint `gorm:"column:plan_item_id;uniqueIndex:idx_plan_group_item" json:"planItemId"`
}

func (TelecomSKPlanRelation) TableName() string {
	return "telecom_sk_plan_relations"
}
