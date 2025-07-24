package schemas

import (
	"gorm.io/gorm"
)

type OrderInternet struct {
	gorm.Model
	IdMember        int    `gorm:"column:id_member;not null;comment:'회원 고유 아이디'" json:"idMember"`
	IdGoodsInternet int    `gorm:"column:id_goods_internet;not null;comment:'인터넷 상품 고유 아이디'" json:"idGoodsInternet"`
	Telecom         string `gorm:"column:telecom;not null;type:varchar(30);comment:'통신사'" json:"telecom"`
	SpeedType       int    `gorm:"column:speed_type;not null;comment:'인터넷 속도 유형 (1: 100 Mbps, 2: 500 Mbps, 3: 1 Gbps)'" json:"speedType"`
	TvType          string `gorm:"column:tv_type;not null;type:varchar(30);comment:'TV 유형'" json:"tvType"`
	OrderType       int    `gorm:"column:order_type;not null;comment:'가입 유형 (1: 통신사 변경, 2: 신규 가입)'" json:"orderType"`
	ProductType     int    `gorm:"column:product_type;not null;comment:'상품 유형 (1: 인터넷만, 2: 인터넷 + TV)'" json:"productType"`
	CombineYn       bool   `gorm:"column:combine_yn;default:false;comment:'휴대폰 결합 여부'" json:"combineYn"`
	EstimatedPrice  int64  `gorm:"column:estimated_price;type:bigint;comment:'예상 금액'" json:"estimatedPrice"`
	Subsidy         int64  `gorm:"column:subsidy;type:bigint;comment:'지원금'" json:"subsidy"`
	NameCustomer    string `gorm:"column:name_customer;not null;comment:'고객명'" json:"nameCustomer"`
	BirthDate       string `gorm:"column:birth_date;type:char(8);not null;comment:'생년월일'" json:"birthDate"`
	NumPhone        string `gorm:"column:num_phone;not null;comment:'연락처'" json:"numPhone"`
	Address         string `gorm:"column:address;comment:'설치 주소'" json:"address"`
	Memo            string `gorm:"column:memo;type:text;comment:'메모'" json:"memo"`
}

func (OrderInternet) TableName() string {
	return "orders_internet"
}
