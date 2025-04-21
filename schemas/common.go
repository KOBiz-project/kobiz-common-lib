package schemas

import "time"

/**
SKT	KT	LGU
모델코드	model_cd	hndset_model_nm	urc_trm_mdl_cd
모델명	items.product_nm	items.pet_nm	items.urc_trm_mdl_nm
용량	items.product_mem	-	-
요금제코드	prod_id	prod_no	urc_mbl_pp_cd
출고가	factory_price	ofw_amt	dlvr_prc
공시지원금	telecom_sale_amt	td_kt_suprt_amt	six_plan_puan_supt_amt (basic_plan_supt_tamt 둘 다 같은 값이라고 하셨음)
*/
type CommonItem struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	CodeModel          string    `gorm:"column:code_model;not null;type:varchar(50);comment:'모델 코드';uniqueIndex:idx_code_model"`
	NameModel          string    `gorm:"column:name_model;not null;type:varchar(256);comment:'모델 명'"`
	Capacity           string    `gorm:"column:capacity;not null;type:varchar(50);comment:'용량'"`
	CodePlan           string    `gorm:"column:code_plan;not null;type:varchar(50);comment:'요금제 코드';uniqueIndex:idx_code_plan"`
	PriceFactory       int       `gorm:"column:price_factory;not null;type:int;comment:'제조사 가격'"`
	PricePublicSupport int       `gorm:"column:price_public_support;not null;type:int;comment:'공시지원금'"`
	CreatedAt          time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP"`
}

func (CommonItem) TableName() string {
	return "common_items"
}

/**
SKT	KT	LGU
요금제명	groups.subscription_nm	items.ppl_nm	items.urc_mbl_pp_nm
요금제코드	groups.subscription_id	items.ppl_id	items.urc_mbl_pp_cd
기본요금	groups.basic_charge	items.puno_month_use_chage	items.urc_pp_basf_amt
*/
type CommonPlan struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	NamePlan   string    `gorm:"column:name_plan;not null;type:varchar(256);comment:'요금제 명'"`
	CodePlan   string    `gorm:"column:code_plan;not null;type:varchar(256);comment:'요금제 코드';uniqueIndex:idx_code_plan"`
	PriceBasic int       `gorm:"column:price_basic;not null;type:int;comment:'기본 요금'"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP"`
}

func (CommonPlan) TableName() string {
	return "common_plans"
}
