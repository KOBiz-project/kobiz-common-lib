package schemas

import "gorm.io/gorm"

// Package models contains the database models for the subscription service

// Subscription 통신사 구독 상품 정보를 담는 구조체
type TelecomSKPlanGroup struct {
	gorm.Model
	CategoryID       string `gorm:"column:category_id;type:varchar(20);comment:'카테고리 ID'" json:"categoryId"`                     // 상품 카테고리 구분 ID
	CategoryNm       string `gorm:"column:category_nm;type:varchar(100);comment:'카테고리명'" json:"categoryNm"`                      // 상품 카테고리 이름
	BasicChargeMin   string `gorm:"column:basic_charge_min;type:varchar(20);comment:'최소 기본요금'" json:"basicChargeMin"`            // 최소 기본 요금
	BasicChargeMax   string `gorm:"column:basic_charge_max;type:varchar(20);comment:'최대 기본요금'" json:"basicChargeMax"`            // 최대 기본 요금
	BasicCharge      string `gorm:"column:basic_charge;type:varchar(20);comment:'기본요금'" json:"basicCharge"`                      // 실제 기본 요금
	SubscriptionNm   string `gorm:"column:subscription_nm;type:varchar(100);comment:'구독상품명'" json:"subscriptionNm"`              // 구독 상품 이름
	SubscriptionID   string `gorm:"column:subscription_id;type:varchar(20);primary_key;comment:'구독상품 ID'" json:"subscriptionId"` // 구독 상품 고유 식별자
	FlagCheck        string `gorm:"column:flag_check;type:varchar(10);comment:'플래그 체크'" json:"flagCheck"`                        // 상품 상태 체크 플래그
	SubLimit         string `gorm:"column:sub_limit;type:varchar(50);comment:'구독 제한'" json:"subLimit"`                           // 구독 제한 사항
	DataInfo         string `gorm:"column:data_info;type:varchar(100);comment:'데이터 정보'" json:"dataInfo"`                         // 데이터 제공 정보
	VoiceInfo        string `gorm:"column:voice_info;type:varchar(100);comment:'통화 정보'" json:"voiceInfo"`                        // 음성 통화 제공 정보
	SMSInfo          string `gorm:"column:sms_info;type:varchar(100);comment:'문자 정보'" json:"smsInfo"`                            // 문자 서비스 제공 정보
	NetTyp           string `gorm:"column:net_typ;type:varchar(10);comment:'네트워크 유형'" json:"netTyp"`                             // 네트워크 유형(5G, LTE 등)
	SubOption        string `gorm:"column:sub_option;type:varchar(200);comment:'구독 옵션'" json:"subOption"`                        // 추가 구독 옵션
	ModelNm          string `gorm:"column:model_nm;type:varchar(100);comment:'모델명'" json:"modelNm"`                              // 대상 모델명
	RcProdID         string `gorm:"column:rc_prod_id;type:varchar(20);comment:'RC 상품 ID'" json:"rcProdId"`                       // RC 상품 식별자
	ZsvcFeeProdID    string `gorm:"column:zsvc_fee_prod_id;type:varchar(20);comment:'서비스 요금 상품 ID'" json:"zsvcFeeProdId"`        // 서비스 요금 상품 식별자
	BeqpEqpMdlNm     string `gorm:"column:beqp_eqp_mdl_nm;type:varchar(100);comment:'장비 모델명'" json:"beqpEqpMdlNm"`               // 장비 모델명
	UserSesnID       string `gorm:"column:user_sesn_id;type:varchar(50);comment:'사용자 세션 ID'" json:"userSesnId"`                  // 사용자 세션 식별자
	BuyBrwsClCd      string `gorm:"column:buy_brws_cl_cd;type:varchar(20);comment:'구매 브라우저 분류 코드'" json:"buyBrwsClCd"`           // 구매 브라우저 분류 코드
	BuyPprocStepCd   string `gorm:"column:buy_pproc_step_cd;type:varchar(20);comment:'구매 처리 단계 코드'" json:"buyPprocStepCd"`       // 구매 처리 단계 코드
	BuyBrwsProdCtt   string `gorm:"column:buy_brws_prod_ctt;type:varchar(200);comment:'구매 브라우저 상품 내용'" json:"buyBrwsProdCtt"`    // 구매 브라우저 상품 내용
	ProductGrpNm     string `gorm:"column:product_grp_nm;type:varchar(100);comment:'상품 그룹명'" json:"productGrpNm"`                // 상품 그룹명
	ProductGrpID     string `gorm:"column:product_grp_id;type:varchar(20);comment:'상품 그룹 ID'" json:"productGrpId"`               // 상품 그룹 식별자
	ChnlMbrID        string `gorm:"column:chnl_mbr_id;type:varchar(50);comment:'채널 회원 ID'" json:"chnlMbrId"`                     // 채널 회원 식별자
	UsedFlag         string `gorm:"column:used_flag;type:varchar(10);comment:'사용 플래그'" json:"usedFlag"`                          // 사용 여부 플래그
	OptSelMndtCnt    string `gorm:"column:opt_sel_mndt_cnt;type:varchar(10);comment:'옵션 선택 필수 개수'" json:"optSelMndtCnt"`         // 필수 선택 옵션 개수
	DisplayYn        string `gorm:"column:display_yn;type:varchar(2);comment:'표시 여부'" json:"displayYn"`                          // 화면 표시 여부
	IcasProdID       string `gorm:"column:icas_prod_id;type:varchar(20);comment:'ICAS 상품 ID'" json:"icasProdId"`                 // ICAS 상품 식별자
	DplanRecommendID string `gorm:"column:dplan_recommend_id;type:varchar(20);comment:'다이렉트 플랜 추천 ID'" json:"dplanRecommendId"`  // 다이렉트 플랜 추천 식별자
	SubcommNm        string `gorm:"column:subcomm_nm;type:varchar(100);comment:'서브 커미션명'" json:"subcommNm"`                      // 서브 커미션명
	SubcommDisamt    string `gorm:"column:subcomm_disamt;type:varchar(20);comment:'서브 커미션 할인금액'" json:"subcommDisamt"`           // 서브 커미션 할인 금액
	SelSubcommYn     string `gorm:"column:sel_subcomm_yn;type:varchar(2);comment:'선택 서브 커미션 여부'" json:"selSubcommYn"`            // 선택 서브 커미션 여부
	SelSubcommDisamt string `gorm:"column:sel_subcomm_disamt;type:varchar(20);comment:'선택 서브 커미션 할인금액'" json:"selSubcommDisamt"` // 선택 서브 커미션 할인 금액
	FmlyNoCtgryFl    string `gorm:"column:fmly_no_ctgry_fl;type:varchar(2);comment:'가족 번호 카테고리 플래그'" json:"fmlyNoCtgryFl"`       // 가족 번호 카테고리 플래그
	DataDtlPhrs      string `gorm:"column:data_dtl_phrs;type:varchar(200);comment:'데이터 상세 문구'" json:"dataDtlPhrs"`               // 데이터 상세 설명
	TcDefaultQty     string `gorm:"column:tc_default_qty;type:varchar(100);comment:'통화 기본 제공량'" json:"tcDefaultQty"`             // 기본 제공 통화량
	SmsDefaultQty    string `gorm:"column:sms_default_qty;type:varchar(100);comment:'문자 기본 제공량'" json:"smsDefaultQty"`           // 기본 제공 문자량
	DataDefaultQty   string `gorm:"column:data_default_qty;type:varchar(100);comment:'데이터 기본 제공량'" json:"dataDefaultQty"`        // 기본 제공 데이터량
	SubcategoryID    string `gorm:"column:subcategory_id;type:varchar(10);comment:'서브카테고리 ID'" json:"subcategoryId"`             // 서브 카테고리 식별자
	BenefitInfo      string `gorm:"column:benefit_info;type:varchar(500);comment:'혜택 정보'" json:"benefitInfo"`                    // 제공 혜택 정보
	SubcommTerm      string `gorm:"column:subcomm_term;type:varchar(10);comment:'서브 커미션 기간'" json:"subcommTerm"`                 // 서브 커미션 제공 기간
	SubcommNoYn      string `gorm:"column:subcomm_no_yn;type:varchar(2);comment:'서브 커미션 번호 여부'" json:"subcommNoYn"`              // 서브 커미션 번호 여부
	IsLatest         bool   `gorm:"column:is_latest; type:bool; default:true; not null; comment:'최신 여부'" json:"isLatest"`
}

func (TelecomSKPlanGroup) TableName() string {
	return "telecom_sk_plan_groups"
}

type TelecomSKPlanItem struct {
	gorm.Model
	SubOptID    string `gorm:"column:sub_opt_id;type:varchar(20);comment:'구독 옵션 ID';uniqueIndex" json:"subOptId"` // 구독 옵션 식별자
	SubOptNm    string `gorm:"column:sub_opt_nm;type:varchar(100);comment:'구독 옵션명'" json:"subOptNm"`              // 구독 옵션 이름
	SubOptTag   string `gorm:"column:sub_opt_tag;type:varchar(200);comment:'구독 옵션 태그'" json:"subOptTag"`          // 구독 옵션 태그 정보
	BasSelYn    string `gorm:"column:bas_sel_yn;type:varchar(2);comment:'기본 선택 여부'" json:"basSelYn"`              // 기본 선택 여부
	OptYn       string `gorm:"column:opt_yn;type:varchar(2);comment:'옵션 여부'" json:"optYn"`                        // 옵션 여부
	BenfDtlDesc string `gorm:"column:benf_dtl_desc;type:varchar(500);comment:'혜택 상세 설명'" json:"benfDtlDesc"`      // 혜택 상세 설명
	BenfImgPath string `gorm:"column:benf_img_path;type:varchar(200);comment:'혜택 이미지 경로'" json:"benfImgPath"`     // 혜택 이미지 경로
	BenfImgDesc string `gorm:"column:benf_img_desc;type:varchar(200);comment:'혜택 이미지 설명'" json:"benfImgDesc"`     // 혜택 이미지 설명
	IsLatest    bool   `gorm:"column:is_latest; type:bool; default:true; not null; comment:'최신 여부'" json:"isLatest"`
}

func (TelecomSKPlanItem) TableName() string {
	return "telecom_sk_plan_items"
}

type TelecomSKPlanRelation struct {
	gorm.Model
	PlanGroupId uint `gorm:"column:plan_group_id;comment:'요금제 그룹 ID';uniqueIndex:idx_plan_group_item" json:"planGroupId"`
	PlanItemId  uint `gorm:"column:plan_item_id;comment:'요금제 ID';uniqueIndex:idx_plan_group_item" json:"planItemId"`
}

func (TelecomSKPlanRelation) TableName() string {
	return "telecom_sk_plan_relations"
}
