package schemas

import "gorm.io/gorm"

type TelecomSKItem struct {
	gorm.Model

	Num              int    `gorm:"column:num; type:int; not null" json:"num"`
	PhoneImg         string `gorm:"column:phone_img; type:varchar(255); not null" json:"phoneImg"`
	ModelCd          string `gorm:"column:model_cd; type:varchar(50); not null" json:"modelCd"`
	CompanyNm        string `gorm:"column:company_nm; type:varchar(255); not null" json:"companyNm"`
	ProductNm        string `gorm:"column:product_nm; type:varchar(255); not null" json:"productNm"`
	ProductGrpID     string `gorm:"column:product_grp_id; type:varchar(50); not null" json:"productGrpId"`
	ProductMem       string `gorm:"column:product_mem; type:varchar(50); not null" json:"productMem"`
	ProductRentAmt   int    `gorm:"column:product_rent_amt; type:int; not null" json:"productRentAmt"`
	ProdID           string `gorm:"column:prod_id; type:varchar(50); not null" json:"prodId"`
	ProdNm           string `gorm:"column:prod_nm; type:varchar(255); not null" json:"prodNm"`
	FactoryPrice     int    `gorm:"column:factory_price; type:int; not null" json:"factoryPrice"`
	FactorySaleAmt   int    `gorm:"column:factory_sale_amt; type:int; not null" json:"factorySaleAmt"`
	TelecomSaleAmt   int    `gorm:"column:telecom_sale_amt; type:int; not null" json:"telecomSaleAmt"`
	TwdSaleAmt       int    `gorm:"column:twd_sale_amt; type:int; not null" json:"twdSaleAmt"`
	SumSaleAmt       int    `gorm:"column:sum_sale_amt; type:int; not null" json:"sumSaleAmt"`
	TwdSumSaleAmt    int    `gorm:"column:twd_sum_sale_amt; type:int; not null" json:"twdSumSaleAmt"`
	Price            int    `gorm:"column:price; type:int; not null" json:"price"`
	TwdPrice         int    `gorm:"column:twd_price; type:int; not null" json:"twdPrice"`
	SaleAmtGrpID     string `gorm:"column:sale_amt_grp_id; type:varchar(50); not null" json:"saleAmtGrpId"`
	SaleYn           string `gorm:"column:sale_yn; type:char(1); not null" json:"saleYn"`
	EffStaDt         string `gorm:"column:eff_sta_dt; type:varchar(8); not null" json:"effStaDt"`
	FactoryDt        string `gorm:"column:factory_dt; type:varchar(8); not null" json:"factoryDt"`
	FeeSaleAmt       int    `gorm:"column:fee_sale_amt; type:int; not null" json:"feeSaleAmt"`
	DiffDiscount     int    `gorm:"column:diff_discount; type:int; not null" json:"diffDiscount"`
	SprateSupmAmt    int    `gorm:"column:sprate_supm_amt; type:int; not null" json:"sprateSupmAmt"`
	ScrbTypCd        string `gorm:"column:scrb_typ_cd; type:varchar(50); not null" json:"scrbTypCd"`
	DcMthdCd         string `gorm:"column:dc_mthd_cd; type:varchar(50); not null" json:"dcMthdCd"`
	Gbn              int    `gorm:"column:gbn; type:int; not null" json:"gbn"`
	Rowspan          int    `gorm:"column:rowspan; type:int; not null" json:"rowspan"`
	BtrMSaleAmt      int    `gorm:"column:btr_m_sale_amt; type:int; not null" json:"btrMSaleAmt"`
	BtrTwdSaleAmt    int    `gorm:"column:btr_twd_sale_amt; type:int; not null" json:"btrTwdSaleAmt"`
	BtrSprateSupmAmt int    `gorm:"column:btr_sprate_supm_amt; type:int; not null" json:"btrSprateSupmAmt"`
	BtrSumSaleAmt    int    `gorm:"column:btr_sum_sale_amt; type:int; not null" json:"btrSumSaleAmt"`
	BtrTwdSumSaleAmt int    `gorm:"column:btr_twd_sum_sale_amt; type:int; not null" json:"btrTwdSumSaleAmt"`
	BtrPrice         int    `gorm:"column:btr_price; type:int; not null" json:"btrPrice"`
	BtrTwdPrice      int    `gorm:"column:btr_twd_price; type:int; not null" json:"btrTwdPrice"`
	IsLatest         bool   `gorm:"column:is_latest; type:bool; default:true; not null" json:"isLatest"`
}

func (TelecomSKItem) TableName() string {
	return "telecom_sk_items"
}
