package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	IdMember            int        `gorm:"column:id_member;not null;comment:'회원 고유 아이디'" json:"idMember"`
	IdGoods             int        `gorm:"column:id_goods;not null;comment:'상품 고유 아이디'" json:"idGoods"`
	IdModel             int        `gorm:"column:id_model;not null;comment:'모델 고유 아이디'" json:"idModel"`
	Telecom             string     `gorm:"column:telecom;not null;type:varchar(30);comment:'통신사'" json:"telecom"`
	NameModel           string     `gorm:"column:name_model;not null;type:varchar(256);comment:'모델명'" json:"nameModel"`
	CustomNameModel     string     `gorm:"column:custom_name_model;type:varchar(256);comment:'커스텀 모델명'" json:"customNameModel"`
	Size                string     `gorm:"column:size;not null;type:varchar(10);comment:'기기용량 64GB/128GB/256GB/512GB/1TB'" json:"size"`
	IdColor             int        `gorm:"column:id_color;not null;comment:'색상 고유 아이디'" json:"idColor"`
	ActivationType      int        `gorm:"column:activation_type;not null;comment:'가입 유형 (1: 번이, 2: 기변, 3: 신규)'" json:"activationType"`
	SubsidyYn           bool       `gorm:"column:subsidy_yn;default:false;comment:'공시지원 여부'" json:"subsidyYn"`
	PaymentType         int        `gorm:"column:payment_type;not null;comment:'결제 유형 (1: 24개월 할부, 2: 현금 완납)'" json:"paymentType"`
	IdPlan              int        `gorm:"column:id_plan;not null;comment:'요금제 고유 아이디'" json:"idPlan"`
	NamePlan            string     `gorm:"column:name_plan;type:varchar(256);comment:'요금제명';" json:"namePlan"`
	CustomNamePlan      string     `gorm:"column:custom_name_plan;type:varchar(256);comment:'커스텀 요금제명';" json:"customNamePlan"`
	BasicPrice          int64      `gorm:"column:basic_price;not null;type:bigint;comment:'요금제 기본요금'" json:"basicPrice"`
	PriceFactory        int64      `gorm:"column:price_factory;not null;type:bigint;comment:'출고가'" json:"priceFactory"`
	Subsidy             int64      `gorm:"column:subsidy;not null;type:bigint;comment:'공시지원금'" json:"subsidy"`
	SwitchSubsidy       int64      `gorm:"column:switch_subsidy;not null;type:bigint;comment:'전환지원금'" json:"switchSubsidy"`
	AddSubsidy          int64      `gorm:"column:add_subsidy;not null;type:bigint;comment:'추가지원금'" json:"addSubsidy"`
	FinalPrice          int64      `gorm:"column:final_price;not null;type:bigint;comment:'할부원금 (출고가에서 각종 지원금을 제한 금액)'" json:"finalPrice"`
	InstallmentFee      int64      `gorm:"column:installment_fee;not null;type:bigint;comment:'할부 수수료 합계 (24개월)'" json:"installmentFee"`
	MonthlyInstallment  int64      `gorm:"column:monthly_installment;not null;type:bigint;comment:'월 할부금(이자포함)'" json:"monthlyInstallment"`
	TotalDiscountAmount int64      `gorm:"column:total_discount_amount;not null;type:bigint;comment:'총 할인 금액 (선택 약정)'" json:"totalDiscountAmount"`
	FinalMonthlyPayment int64      `gorm:"column:final_monthly_payment;not null;type:bigint;comment:'최종 월 납부금'" json:"finalMonthlyPayment"`
	Status              int        `gorm:"column:status;not null;default:1;comment:'주문 상태 (1: 접수 완료, 2: 접수 확인, 3: URL 완료, 4: URL 및 발송 완료, 5: 개통 완료, 6: 반품 처리 중, 7: 반품 완료)'" json:"status"`
	ActivationDt        *time.Time `gorm:"column:activation_dt;default:null;type:date;comment:'개통일자'" json:"activationDt"`
	ReturnsReason       string     `gorm:"column:returns_reason;type:text;comment:'반품 사유'" json:"returnsReason"`
	CustomerType        int        `gorm:"column:customer_type;not null;comment:'가입자 유형 (1: 개인, 2: 가족)'" json:"customerType"`
	NameEmployee        string     `gorm:"column:name_employee;comment:'임직원 성함'" json:"nameEmployee"`
	NumPhoneEmployee    string     `gorm:"column:num_phone_employee;comment:'임직원 연락처'" json:"numPhoneEmployee"`
	NameCustomer        string     `gorm:"column:name_customer;not null;comment:'고객명'" json:"nameCustomer"`
	NumPhone            string     `gorm:"column:num_phone;not null;comment:'개통할 번호'" json:"numPhone"`
	NumPhoneExt         string     `gorm:"column:num_phone_ext;comment:'다른 연락처'" json:"numPhoneExt"`
	Address             string     `gorm:"column:address;comment:'택배 주소'" json:"address"`
	BirthDate           string     `gorm:"column:birth_date;type:char(8);not null;comment:'생년월일'" json:"birthDate"`
	Memo                string     `gorm:"column:memo;type:text;comment:'메모'" json:"memo"`
	MemoAdmin           string     `gorm:"column:memo_admin;type:text;comment:'관리자 메모'" json:"memoAdmin"`
	IdManager           int        `gorm:"column:id_manager;comment:'담당자 고유 아이디'" json:"idManager"`
	IdReception         int        `gorm:"column:id_reception;comment:'개통처 고유 아이디'" json:"idReception"`
}

func (Order) TableName() string {
	return "orders"
}
