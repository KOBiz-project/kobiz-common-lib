package schemas

import "gorm.io/gorm"

// ListDataKT 구조체 정의 (MySQL에 맞게 작성)
// GORM 모델을 사용해 MySQL 데이터베이스의 테이블과 매핑
type TelecomKTItem struct {
	gorm.Model // GORM 기본 필드: ID, CreatedAt, UpdatedAt, DeletedAt
	// ProdNo: 제품 번호
	ProdNo string `gorm:"column:prod_no; type:varchar(50); not null; comment:'제품 번호'" json:"prodNo"`
	// PetNm: 제품 이름
	PetNm string `gorm:"column:pet_nm; type:varchar(255); not null; comment:'제품 이름'" json:"petNm"`
	// HndsetModelId: 핸드셋 모델 ID
	HndsetModelId string `gorm:"column:hndset_model_id; type:varchar(50); not null; comment:'핸드셋 모델 ID'" json:"hndsetModelId"`
	// HndsetModelNm: 핸드셋 모델 이름
	HndsetModelNm string `gorm:"column:hndset_model_nm; type:varchar(255); not null; comment:'핸드셋 모델 이름'" json:"hndsetModelNm"`
	// MakrCd: 제조사 코드
	MakrCd string `gorm:"column:makr_cd; type:varchar(10); not null; comment:'제조사 코드'" json:"makrCd"`
	// ShowOdrg: 표시 순서
	ShowOdrg string `gorm:"column:show_odrg; type:varchar(10); not null; comment:'표시 순서'" json:"showOdrg"`
	// OfwAmt: 공시지원금 금액
	OfwAmt string `gorm:"column:ofw_amt; type:varchar(20); not null; comment:'공시지원금 금액'" json:"ofwAmt"`
	// StorSuprtAmt: 매장 지원금 금액
	StorSuprtAmt string `gorm:"column:stor_suprt_amt; type:varchar(20); not null; comment:'매장 지원금 금액'" json:"storSuprtAmt"`
	// MonthUseChageDcAmt: 월 사용 요금 할인 금액
	MonthUseChageDcAmt string `gorm:"column:month_use_chage_dc_amt; type:varchar(20); not null; comment:'월 사용 요금 할인 금액'" json:"monthUseChageDcAmt"`
	// RealAmt: 실제 금액
	RealAmt string `gorm:"column:real_amt; type:varchar(20); not null; comment:'실제 금액'" json:"realAmt"`
	// PplId: 요금제 ID
	PplId string `gorm:"column:ppl_id; type:varchar(50); not null; comment:'요금제 ID'" json:"pplId"`
	// PplNm: 요금제 이름
	PplNm string `gorm:"column:ppl_nm; type:varchar(255); not null; comment:'요금제 이름'" json:"pplNm"`
	// PplGroupDivCd: 요금제 그룹 코드
	PplGroupDivCd string `gorm:"column:ppl_group_div_cd; type:varchar(50); not null; comment:'요금제 그룹 코드'" json:"pplGroupDivCd"`
	// HndSetImgNm: 핸드셋 이미지 이름
	HndSetImgNm string `gorm:"column:hnd_set_img_nm; type:varchar(255); not null; comment:'핸드셋 이미지 이름'" json:"hndSetImgNm"`
	// MshopHndSetImgNm: 매장 핸드셋 이미지 이름
	MshopHndSetImgNm string `gorm:"column:mshop_hnd_set_img_nm; type:varchar(255); not null; comment:'매장 핸드셋 이미지 이름'" json:"mshopHndSetImgNm"`
	// SortProd: 정렬 기준 (NULL 허용)
	SortProd *string `gorm:"column:sort_prod; type:varchar(50); comment:'정렬 기준'" json:"sortProd"`
	// DscnOptnCd: 할인 옵션 코드 (NULL 허용)
	DscnOptnCd *string `gorm:"column:dscn_optn_cd; type:varchar(50); comment:'할인 옵션 코드'" json:"dscnOptnCd"`
	// DispCtgCd: 디스플레이 카테고리 코드 (NULL 허용)
	DispCtgCd *string `gorm:"column:disp_ctg_cd; type:varchar(50); comment:'디스플레이 카테고리 코드'" json:"dispCtgCd"`
	// ConvSupotProdYn: 전환 지원 여부 (Y/N, NULL 허용)
	ConvSupotProdYn *string `gorm:"column:conv_supot_prod_yn; type:char(1); comment:'전환 지원 여부 (Y/N)'" json:"convSupotProdYn"`
	// TdKtSuprtAmt: KT 지원금 금액
	TdKtSuprtAmt string `gorm:"column:td_kt_suprt_amt; type:varchar(20); not null; comment:'KT 지원금 금액'" json:"tdKtSuprtAmt"`
	// TcKtSuprtAmt: 추가 KT 지원금 금액 (NULL 허용)
	TcKtSuprtAmt *string `gorm:"column:tc_kt_suprt_amt; type:varchar(20); comment:'추가 KT 지원금 금액'" json:"tcKtSuprtAmt"`
	// KdKtSuprtAmt: KD 지원금 금액
	KdKtSuprtAmt string `gorm:"column:kd_kt_suprt_amt; type:varchar(20); not null; comment:'KD 지원금 금액'" json:"kdKtSuprtAmt"`
	// KcKtSuprtAmt: KC 지원금 금액 (NULL 허용)
	KcKtSuprtAmt *string `gorm:"column:kc_kt_suprt_amt; type:varchar(20); comment:'KC 지원금 금액'" json:"kcKtSuprtAmt"`
	// PrdcCd: 제품 코드 (NULL 허용)
	PrdcCd *string `gorm:"column:prdc_cd; type:varchar(50); comment:'제품 코드'" json:"prdcCd"`
	// SpnsGrpCd: 스폰서 그룹 코드 (NULL 허용)
	SpnsGrpCd *string `gorm:"column:spns_grp_cd; type:varchar(50); comment:'스폰서 그룹 코드'" json:"spnsGrpCd"`
	// SpnsMonsType: 스폰서 월 유형 (NULL 허용)
	SpnsMonsType *string `gorm:"column:spns_mons_type; type:varchar(50); comment:'스폰서 월 유형'" json:"spnsMonsType"`
	// StrRow: 시작 행 (NULL 허용)
	StrRow *string `gorm:"column:str_row; type:varchar(50); comment:'시작 행'" json:"strRow"`
	// ToRow: 종료 행 (NULL 허용)
	ToRow *string `gorm:"column:to_row; type:varchar(50); comment:'종료 행'" json:"toRow"`
	// DeviceType: 디바이스 유형 (NULL 허용)
	DeviceType *string `gorm:"column:device_type; type:varchar(50); comment:'디바이스 유형'" json:"deviceType"`
	// SpnsrPunoDate: 스폰서 지원 날짜
	SpnsrPunoDate string `gorm:"column:spnsr_puno_date; type:varchar(20); not null; comment:'스폰서 지원 날짜'" json:"spnsrPunoDate"`
	// SaleSttusCd: 판매 상태 코드
	SaleSttusCd string `gorm:"column:sale_sttus_cd; type:varchar(10); not null; comment:'판매 상태 코드'" json:"saleSttusCd"`
	// DispYn: 표시 여부 (Y/N)
	DispYn string `gorm:"column:disp_yn; type:char(1); not null; comment:'표시 여부 (Y/N)'" json:"dispYn"`
	// IsLatest: 최신 여부 (기본값 true)
	IsLatest bool `gorm:"column:is_latest; type:bool; default:true; not null; comment:'최신 여부'" json:"isLatest"`
}

// TableName 커스터마이징 (MySQL 테이블 이름 지정)
func (TelecomKTItem) TableName() string {
	return "telecom_kt_items"
}
