package schemas

import "gorm.io/gorm"

// KTPlanItem 구조체 정의 (MySQL에 맞게 작성)
// GORM 모델을 사용해 MySQL 데이터베이스의 테이블과 매핑
type TelecomKTPlanItem struct {
	gorm.Model // GORM 기본 필드: ID, CreatedAt, UpdatedAt, DeletedAt
	// PplGb: 요금제 구분
	PplGb string `gorm:"column:ppl_gb; type:varchar(10); not null; comment:'요금제 구분'" json:"pplGb"`
	// CtgNameForMenu: 메뉴용 카테고리 이름 (NULL 허용)
	CtgNameForMenu *string `gorm:"column:ctg_name_for_menu; type:varchar(255); comment:'메뉴용 카테고리 이름'" json:"ctgNameForMenu"`
	// PplGrpNm: 요금제 그룹 이름 (NULL 허용)
	PplGrpNm *string `gorm:"column:ppl_grp_nm; type:varchar(255); comment:'요금제 그룹 이름'" json:"pplGrpNm"`
	// PplId: 요금제 ID
	PplId string `gorm:"column:ppl_id; type:varchar(50); not null; comment:'요금제 ID'" json:"pplId"`
	// PplNm: 요금제 이름
	PplNm string `gorm:"column:ppl_nm; type:varchar(255); not null; comment:'요금제 이름'" json:"pplNm"`
	// PplGrpCd: 요금제 그룹 코드
	PplGrpCd string `gorm:"column:ppl_grp_cd; type:varchar(50); not null; comment:'요금제 그룹 코드'" json:"pplGrpCd"`
	// PplDivCd: 요금제 구분 코드 (NULL 허용)
	PplDivCd *string `gorm:"column:ppl_div_cd; type:varchar(50); comment:'요금제 구분 코드'" json:"pplDivCd"`
	// MonthUseChage: 월 사용 요금 (NULL 허용)
	MonthUseChage *string `gorm:"column:month_use_chage; type:varchar(50); comment:'월 사용 요금'" json:"monthUseChage"`
	// EngtDcAmt: 약정 할인 금액 (NULL 허용)
	EngtDcAmt *string `gorm:"column:engt_dc_amt; type:varchar(50); comment:'약정 할인 금액'" json:"engtDcAmt"`
	// FreeTlkPrvQnt: 무료 통화 제공량 (NULL 허용)
	FreeTlkPrvQnt *string `gorm:"column:free_tlk_prv_qnt; type:varchar(100); comment:'무료 통화 제공량'" json:"freeTlkPrvQnt"`
	// FreeChrPrvQnt: 무료 문자 제공량 (NULL 허용)
	FreeChrPrvQnt *string `gorm:"column:free_chr_prv_qnt; type:varchar(100); comment:'무료 문자 제공량'" json:"freeChrPrvQnt"`
	// FreeDataPrvQnt: 무료 데이터 제공량 (NULL 허용)
	FreeDataPrvQnt *string `gorm:"column:free_data_prv_qnt; type:varchar(100); comment:'무료 데이터 제공량'" json:"freeDataPrvQnt"`
	// OllehMblFreeVcTlkQnt: 올레 모바일 무료 음성 통화량 (NULL 허용)
	OllehMblFreeVcTlkQnt *string `gorm:"column:olleh_mbl_free_vc_tlk_qnt; type:varchar(100); comment:'올레 모바일 무료 음성 통화량'" json:"ollehMblFreeVcTlkQnt"`
	// SfznApdDataQnt: 안심 데이터 제공량 (NULL 허용)
	SfznApdDataQnt *string `gorm:"column:sfzn_apd_data_qnt; type:varchar(100); comment:'안심 데이터 제공량'" json:"sfznApdDataQnt"`
	// DataFreeOptnYn: 데이터 무료 옵션 여부 (Y/N, NULL 허용)
	DataFreeOptnYn *string `gorm:"column:data_free_optn_yn; type:char(1); comment:'데이터 무료 옵션 여부 (Y/N)'" json:"dataFreeOptnYn"`
	// ChrTrmOptnYn: 요금 기간 옵션 여부 (Y/N, NULL 허용)
	ChrTrmOptnYn *string `gorm:"column:chr_trm_optn_yn; type:char(1); comment:'요금 기간 옵션 여부 (Y/N)'" json:"chrTrmOptnYn"`
	// DataCfwdPrvYn: 데이터 이월 제공 여부 (Y/N, NULL 허용)
	DataCfwdPrvYn *string `gorm:"column:data_cfwd_prv_yn; type:char(1); comment:'데이터 이월 제공 여부 (Y/N)'" json:"dataCfwdPrvYn"`
	// RcmdPplYn: 추천 요금제 여부 (Y/N, NULL 허용)
	RcmdPplYn *string `gorm:"column:rcmd_ppl_yn; type:char(1); comment:'추천 요금제 여부 (Y/N)'" json:"rcmdPplYn"`
	// DtlDescSbst: 상세 설명 내용 (NULL 허용)
	DtlDescSbst *string `gorm:"column:dtl_desc_sbst; type:text; comment:'상세 설명 내용'" json:"dtlDescSbst"`
	// IndcOdrg: 표시 순서 (NULL 허용)
	IndcOdrg *string `gorm:"column:indc_odrg; type:varchar(10); comment:'표시 순서'" json:"indcOdrg"`
	// UseYn: 사용 여부 (Y/N, NULL 허용)
	UseYn *string `gorm:"column:use_yn; type:char(1); comment:'사용 여부 (Y/N)'" json:"useYn"`
	// UseDivCd: 사용 구분 코드 (NULL 허용)
	UseDivCd *string `gorm:"column:use_div_cd; type:varchar(50); comment:'사용 구분 코드'" json:"useDivCd"`
	// Expnsn1StrVal: 확장1 문자열 값 (NULL 허용)
	Expnsn1StrVal *string `gorm:"column:expnsn1_str_val; type:varchar(255); comment:'확장1 문자열 값'" json:"expnsn1StrVal"`
	// OnfrmCd: 확인 코드
	OnfrmCd string `gorm:"column:onfrm_cd; type:varchar(50); not null; comment:'확인 코드'" json:"onfrmCd"`
	// PplSectCd: 요금제 섹션 코드 (NULL 허용)
	PplSectCd *string `gorm:"column:ppl_sect_cd; type:varchar(50); comment:'요금제 섹션 코드'" json:"pplSectCd"`
	// PplSectNm: 요금제 섹션 이름 (NULL 허용)
	PplSectNm *string `gorm:"column:ppl_sect_nm; type:varchar(255); comment:'요금제 섹션 이름'" json:"pplSectNm"`
	// SpnsrTypeCd: 스폰서 유형 코드 (NULL 허용)
	SpnsrTypeCd *string `gorm:"column:spnsr_type_cd; type:varchar(50); comment:'스폰서 유형 코드'" json:"spnsrTypeCd"`
	// TngrPplYn: 청소년 요금제 여부 (Y/N, NULL 허용)
	TngrPplYn *string `gorm:"column:tngr_ppl_yn; type:char(1); comment:'청소년 요금제 여부 (Y/N)'" json:"tngrPplYn"`
	// ChgUseYn: 변경 사용 여부 (Y/N, NULL 허용)
	ChgUseYn *string `gorm:"column:chg_use_yn; type:char(1); comment:'변경 사용 여부 (Y/N)'" json:"chgUseYn"`
	// SysRegrId: 시스템 등록자 ID (NULL 허용)
	SysRegrId *string `gorm:"column:sys_regr_id; type:varchar(50); comment:'시스템 등록자 ID'" json:"sysRegrId"`
	// SysRegDt: 시스템 등록 일시 (NULL 허용)
	SysRegDt *string `gorm:"column:sys_reg_dt; type:varchar(20); comment:'시스템 등록 일시'" json:"sysRegDt"`
	// SysAmdrId: 시스템 수정자 ID (NULL 허용)
	SysAmdrId *string `gorm:"column:sys_amdr_id; type:varchar(50); comment:'시스템 수정자 ID'" json:"sysAmdrId"`
	// SysAmdDt: 시스템 수정 일시 (NULL 허용)
	SysAmdDt *string `gorm:"column:sys_amd_dt; type:varchar(20); comment:'시스템 수정 일시'" json:"sysAmdDt"`
	// CombDcPplCd: 결합 할인 요금제 코드 (NULL 허용)
	CombDcPplCd *string `gorm:"column:comb_dc_ppl_cd; type:varchar(50); comment:'결합 할인 요금제 코드'" json:"combDcPplCd"`
	// ChageDispNo: 요금 표시 번호
	ChageDispNo string `gorm:"column:chage_disp_no; type:varchar(10); not null; comment:'요금 표시 번호'" json:"chageDispNo"`
	// DataBasic: 데이터 기본 제공량
	DataBasic string `gorm:"column:data_basic; type:varchar(255); not null; comment:'데이터 기본 제공량'" json:"dataBasic"`
	// DataBenefit: 데이터 혜택 (NULL 허용)
	DataBenefit *string `gorm:"column:data_benefit; type:varchar(255); comment:'데이터 혜택'" json:"dataBenefit"`
	// TlkBasic: 통화 기본 제공량
	TlkBasic string `gorm:"column:tlk_basic; type:varchar(255); not null; comment:'통화 기본 제공량'" json:"tlkBasic"`
	// CharBasic: 문자 기본 제공량
	CharBasic string `gorm:"column:char_basic; type:varchar(255); not null; comment:'문자 기본 제공량'" json:"charBasic"`
	// TlkBenefit: 통화 혜택 (NULL 허용)
	TlkBenefit *string `gorm:"column:tlk_benefit; type:varchar(255); comment:'통화 혜택'" json:"tlkBenefit"`
	// ChageNote: 요금 노트
	ChageNote string `gorm:"column:chage_note; type:text; not null; comment:'요금 노트'" json:"chageNote"`
	// PplDcUseYn: 요금제 할인 사용 여부 (Y/N, NULL 허용)
	PplDcUseYn *string `gorm:"column:ppl_dc_use_yn; type:char(1); comment:'요금제 할인 사용 여부 (Y/N)'" json:"pplDcUseYn"`
	// OnlineFrmUseYn: 온라인 양식 사용 여부 (Y/N, NULL 허용)
	OnlineFrmUseYn *string `gorm:"column:online_frm_use_yn; type:char(1); comment:'온라인 양식 사용 여부 (Y/N)'" json:"onlineFrmUseYn"`
	// LimitPplApyYn: 제한 요금제 적용 여부 (Y/N, NULL 허용)
	LimitPplApyYn *string `gorm:"column:limit_ppl_apy_yn; type:char(1); comment:'제한 요금제 적용 여부 (Y/N)'" json:"limitPplApyYn"`
	// PrmtnDcAmt: 프로모션 할인 금액 (NULL 허용)
	PrmtnDcAmt *string `gorm:"column:prmtn_dc_amt; type:varchar(50); comment:'프로모션 할인 금액'" json:"prmtnDcAmt"`
	// UnityKeyword: 통합 키워드 (NULL 허용)
	UnityKeyword *string `gorm:"column:unity_keyword; type:varchar(255); comment:'통합 키워드'" json:"unityKeyword"`
	// BnfBdgTxt: 혜택 배지 텍스트 (NULL 허용)
	BnfBdgTxt *string `gorm:"column:bnf_bdg_txt; type:varchar(255); comment:'혜택 배지 텍스트'" json:"bnfBdgTxt"`
	// BnfBdgClr: 혜택 배지 색상 (NULL 허용)
	BnfBdgClr *string `gorm:"column:bnf_bdg_clr; type:varchar(20); comment:'혜택 배지 색상'" json:"bnfBdgClr"`
	// TxtBnfDesc: 텍스트 혜택 설명
	TxtBnfDesc string `gorm:"column:txt_bnf_desc; type:text; not null; comment:'텍스트 혜택 설명'" json:"txtBnfDesc"`
	// TxtBnfTltpDispYn: 텍스트 혜택 툴팁 표시 여부 (Y/N, NULL 허용)
	TxtBnfTltpDispYn *string `gorm:"column:txt_bnf_tltp_disp_yn; type:char(1); comment:'텍스트 혜택 툴팁 표시 여부 (Y/N)'" json:"txtBnfTltpDispYn"`
	// TxtBnfTltpDesc: 텍스트 혜택 툴팁 설명 (NULL 허용)
	TxtBnfTltpDesc *string `gorm:"column:txt_bnf_tltp_desc; type:text; comment:'텍스트 혜택 툴팁 설명'" json:"txtBnfTltpDesc"`
	// PplGroupSortCd: 요금제 그룹 정렬 코드 (NULL 허용)
	PplGroupSortCd *string `gorm:"column:ppl_group_sort_cd; type:varchar(50); comment:'요금제 그룹 정렬 코드'" json:"pplGroupSortCd"`
	// ProdAdtnDefVal: 상품 추가 기본값 (NULL 허용)
	ProdAdtnDefVal *string `gorm:"column:prod_adtn_def_val; type:varchar(255); comment:'상품 추가 기본값'" json:"prodAdtnDefVal"`
	// SeniorYn: 시니어 여부 (Y/N, NULL 허용)
	SeniorYn *string `gorm:"column:senior_yn; type:char(1); comment:'시니어 여부 (Y/N)'" json:"seniorYn"`
	// DailyReductionCost: 일일 감소 비용 (NULL 허용)
	DailyReductionCost *string `gorm:"column:daily_reduction_cost; type:varchar(50); comment:'일일 감소 비용'" json:"dailyReductionCost"`
	// VoiceCallCost: 음성 통화 비용 (NULL 허용)
	VoiceCallCost *string `gorm:"column:voice_call_cost; type:varchar(50); comment:'음성 통화 비용'" json:"voiceCallCost"`
	// VideoCallCost: 영상 통화 비용 (NULL 허용)
	VideoCallCost *string `gorm:"column:video_call_cost; type:varchar(50); comment:'영상 통화 비용'" json:"videoCallCost"`
	// ShortTextCost: 단문 메시지 비용 (NULL 허용)
	ShortTextCost *string `gorm:"column:short_text_cost; type:varchar(50); comment:'단문 메시지 비용'" json:"shortTextCost"`
	// LongTextCost: 장문 메시지 비용 (NULL 허용)
	LongTextCost *string `gorm:"column:long_text_cost; type:varchar(50); comment:'장문 메시지 비용'" json:"longTextCost"`
	// MmsTextCost: MMS 메시지 비용 (NULL 허용)
	MmsTextCost *string `gorm:"column:mms_text_cost; type:varchar(50); comment:'MMS 메시지 비용'" json:"mmsTextCost"`
	// CostPerPacket: 패킷당 비용 (NULL 허용)
	CostPerPacket *string `gorm:"column:cost_per_packet; type:varchar(50); comment:'패킷당 비용'" json:"costPerPacket"`
	// CostPerPacketNormal: 일반 패킷당 비용 (NULL 허용)
	CostPerPacketNormal *string `gorm:"column:cost_per_packet_normal; type:varchar(50); comment:'일반 패킷당 비용'" json:"costPerPacketNormal"`
	// MTxtBnfDesc: 모바일 텍스트 혜택 설명
	MTxtBnfDesc string `gorm:"column:m_txt_bnf_desc; type:text; not null; comment:'모바일 텍스트 혜택 설명'" json:"mTxtBnfDesc"`
	// PrsnlShare: 개인 공유
	PrsnlShare string `gorm:"column:prsnl_share; type:varchar(255); not null; comment:'개인 공유'" json:"prsnlShare"`
	// ROAMING: 로밍
	ROAMING string `gorm:"column:roaming; type:varchar(12550); not null; comment:'로밍'" json:"rOAMING"`
	// PrmBnMImgPath: 프로모션 배너 이미지 경로 (NULL 허용)
	PrmBnMImgPath *string `gorm:"column:prm_bn_m_img_path; type:varchar(255); comment:'프로모션 배너 이미지 경로'" json:"prmBnMImgPath"`
	// PrmBnMImgName: 프로모션 배너 이미지 이름 (NULL 허용)
	PrmBnMImgName *string `gorm:"column:prm_bn_m_img_name; type:varchar(255); comment:'프로모션 배너 이미지 이름'" json:"prmBnMImgName"`
	// PunoMonthUseChage: 순수 월 사용 요금
	PunoMonthUseChage string `gorm:"column:puno_month_use_chage; type:varchar(50); not null; comment:'순수 월 사용 요금'" json:"punoMonthUseChage"`
	// PunoMonthUseDcChage: 순수 월 사용 할인 요금
	PunoMonthUseDcChage string `gorm:"column:puno_month_use_dc_chage; type:varchar(50); not null; comment:'순수 월 사용 할인 요금'" json:"punoMonthUseDcChage"`
	// IsLatest: 최신 여부 (기본값 true)
	IsLatest bool `gorm:"column:is_latest; type:bool; default:true; not null; comment:'최신 여부'" json:"isLatest"`
}

// TableName 커스터마이징 (MySQL 테이블 이름 지정)
func (TelecomKTPlanItem) TableName() string {
	return "telecom_kt_plan_items"
}
