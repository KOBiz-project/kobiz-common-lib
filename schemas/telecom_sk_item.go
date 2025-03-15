package schemas

import "gorm.io/gorm"

type TelecomSKItem struct {
	gorm.Model // GORM 기본 필드: ID, CreatedAt, UpdatedAt, DeletedAt

	// Num: 순번
	Num int `gorm:"column:num; type:int; not null; comment:'순번'" json:"num"`
	// PhoneImg: 휴대폰 이미지 경로
	PhoneImg string `gorm:"column:phone_img; type:varchar(255); not null; comment:'휴대폰 이미지 경로'" json:"phoneImg"`
	// ModelCd: 모델 코드
	ModelCd string `gorm:"column:model_cd; type:varchar(50); not null; comment:'모델 코드'" json:"modelCd"`
	// CompanyNm: 제조사명
	CompanyNm string `gorm:"column:company_nm; type:varchar(255); not null; comment:'제조사명'" json:"companyNm"`
	// ProductNm: 제품명
	ProductNm string `gorm:"column:product_nm; type:varchar(255); not null; comment:'제품명'" json:"productNm"`
	// ProductGrpID: 제품 그룹 ID
	ProductGrpID string `gorm:"column:product_grp_id; type:varchar(50); not null; comment:'제품 그룹 ID'" json:"productGrpId"`
	// ProductMem: 제품 용량
	ProductMem string `gorm:"column:product_mem; type:varchar(50); not null; comment:'제품 용량'" json:"productMem"`
	// ProductRentAmt: 제품 렌탈 금액
	ProductRentAmt int `gorm:"column:product_rent_amt; type:int; not null; comment:'제품 렌탈 금액'" json:"productRentAmt"`
	// ProdID: 요금제 ID
	ProdID string `gorm:"column:prod_id; type:varchar(50); not null; comment:'요금제 ID'" json:"prodId"`
	// ProdNm: 요금제명
	ProdNm string `gorm:"column:prod_nm; type:varchar(255); not null; comment:'요금제명'" json:"prodNm"`
	// FactoryPrice: 출고가
	FactoryPrice int `gorm:"column:factory_price; type:int; not null; comment:'출고가'" json:"factoryPrice"`
	// FactorySaleAmt: 제조사 할인금액
	FactorySaleAmt int `gorm:"column:factory_sale_amt; type:int; not null; comment:'제조사 할인금액'" json:"factorySaleAmt"`
	// TelecomSaleAmt: 통신사 공시지원금
	TelecomSaleAmt int `gorm:"column:telecom_sale_amt; type:int; not null; comment:'통신사 공시지원금'" json:"telecomSaleAmt"`
	// TwdSaleAmt: 추가지원금
	TwdSaleAmt int `gorm:"column:twd_sale_amt; type:int; not null; comment:'추가지원금'" json:"twdSaleAmt"`
	// SumSaleAmt: 총 할인금액
	SumSaleAmt int `gorm:"column:sum_sale_amt; type:int; not null; comment:'총 할인금액'" json:"sumSaleAmt"`
	// TwdSumSaleAmt: 추가지원금 포함 총 할인금액
	TwdSumSaleAmt int `gorm:"column:twd_sum_sale_amt; type:int; not null; comment:'추가지원금 포함 총 할인금액'" json:"twdSumSaleAmt"`
	// Price: 할인 후 구매가
	Price int `gorm:"column:price; type:int; not null; comment:'할인 후 구매가'" json:"price"`
	// TwdPrice: 추가지원금 적용 후 구매가
	TwdPrice int `gorm:"column:twd_price; type:int; not null; comment:'추가지원금 적용 후 구매가'" json:"twdPrice"`
	// SaleAmtGrpID: 할인금액 그룹 ID
	SaleAmtGrpID string `gorm:"column:sale_amt_grp_id; type:varchar(50); not null; comment:'할인금액 그룹 ID'" json:"saleAmtGrpId"`
	// SaleYn: 판매여부 Y/N
	SaleYn string `gorm:"column:sale_yn; type:char(1); not null; comment:'판매여부 Y/N'" json:"saleYn"`
	// EffStaDt: 적용 시작일자 YYYYMMDD
	EffStaDt string `gorm:"column:eff_sta_dt; type:varchar(8); not null; comment:'적용 시작일자 YYYYMMDD'" json:"effStaDt"`
	// FactoryDt: 출고일자 YYYYMMDD
	FactoryDt string `gorm:"column:factory_dt; type:varchar(8); not null; comment:'출고일자 YYYYMMDD'" json:"factoryDt"`
	// FeeSaleAmt: 요금 할인액
	FeeSaleAmt int `gorm:"column:fee_sale_amt; type:int; not null; comment:'요금 할인액'" json:"feeSaleAmt"`
	// DiffDiscount: 차액 할인금액
	DiffDiscount int `gorm:"column:diff_discount; type:int; not null; comment:'차액 할인금액'" json:"diffDiscount"`
	// SprateSupmAmt: 별도 지원금액
	SprateSupmAmt int `gorm:"column:sprate_supm_amt; type:int; not null; comment:'별도 지원금액'" json:"sprateSupmAmt"`
	// ScrbTypCd: 가입유형 코드
	ScrbTypCd string `gorm:"column:scrb_typ_cd; type:varchar(50); not null; comment:'가입유형 코드'" json:"scrbTypCd"`
	// DcMthdCd: 할인방법 코드
	DcMthdCd string `gorm:"column:dc_mthd_cd; type:varchar(50); not null; comment:'할인방법 코드'" json:"dcMthdCd"`
	// Gbn: 구분값
	Gbn int `gorm:"column:gbn; type:int; not null; comment:'구분값'" json:"gbn"`
	// Rowspan: 행 병합 수
	Rowspan int `gorm:"column:rowspan; type:int; not null; comment:'행 병합 수'" json:"rowspan"`
	// BtrMSaleAmt: 더 나은 월 할인금액
	BtrMSaleAmt int `gorm:"column:btr_m_sale_amt; type:int; not null; comment:'더 나은 월 할인금액'" json:"btrMSaleAmt"`
	// BtrTwdSaleAmt: 더 나은 추가지원금
	BtrTwdSaleAmt int `gorm:"column:btr_twd_sale_amt; type:int; not null; comment:'더 나은 추가지원금'" json:"btrTwdSaleAmt"`
	// BtrSprateSupmAmt: 더 나은 별도 지원금액
	BtrSprateSupmAmt int `gorm:"column:btr_sprate_supm_amt; type:int; not null; comment:'더 나은 별도 지원금액'" json:"btrSprateSupmAmt"`
	// BtrSumSaleAmt: 더 나은 총 할인금액
	BtrSumSaleAmt int `gorm:"column:btr_sum_sale_amt; type:int; not null; comment:'더 나은 총 할인금액'" json:"btrSumSaleAmt"`
	// BtrTwdSumSaleAmt: 더 나은 추가지원금 포함 총 할인금액
	BtrTwdSumSaleAmt int `gorm:"column:btr_twd_sum_sale_amt; type:int; not null; comment:'더 나은 추가지원금 포함 총 할인금액'" json:"btrTwdSumSaleAmt"`
	// BtrPrice: 더 나은 할인 후 가격
	BtrPrice int `gorm:"column:btr_price; type:int; not null; comment:'더 나은 할인 후 가격'" json:"btrPrice"`
	// BtrTwdPrice: 더 나은 추가지원금 적용 후 가격
	BtrTwdPrice int `gorm:"column:btr_twd_price; type:int; not null; comment:'더 나은 추가지원금 적용 후 가격'" json:"btrTwdPrice"`
	// IsLatest: 최신 여부 (기본값 true)
	IsLatest bool `gorm:"column:is_latest; type:bool; default:true; not null; comment:'최신 여부'" json:"isLatest"`
}

// TableName GORM 테이블 이름 지정
func (TelecomSKItem) TableName() string {
	return "telecom_sk_items"
}
