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
	CombineYn       bool   `gorm:"column:combine_yn;default:false;comment:'휴대폰 결합 여부'" json:"combineYn"`
	EstimatedPrice  int64  `gorm:"column:estimated_price;type:bigint;comment:'예상 금액'" json:"estimatedPrice"`
	Subsidy         int64  `gorm:"column:subsidy;type:bigint;comment:'지원금'" json:"subsidy"`
}

func (OrderInternet) TableName() string {
	return "orders_internet"
}
