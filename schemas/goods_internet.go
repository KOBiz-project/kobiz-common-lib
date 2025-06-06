package schemas

import (
	"gorm.io/gorm"
)

// 자체 인터넷 상품 스키마
type GoodsInternet struct {
	gorm.Model
	Telecom        string `gorm:"column:telecom;not null;type:varchar(30);comment:'통신사';uniqueIndex:idx_goods_internet" json:"telecom"`
	SpeedType      int    `gorm:"column:speed_type;not null;comment:'인터넷 속도 유형 (1: 100 Mbps, 2: 500 Mbps, 3: 1 Gbps)';uniqueIndex:idx_goods_internet" json:"speedType"`
	TvType         string `gorm:"column:tv_type;not null;type:varchar(30);comment:'TV 유형';uniqueIndex:idx_goods_internet" json:"tvType"`
	EstimatedPrice int64  `gorm:"column:estimated_price;type:bigint;comment:'예상 금액'" json:"estimatedPrice"`
	Subsidy        int64  `gorm:"column:subsidy;type:bigint;comment:'지원금'" json:"subsidy"`
}

func (GoodsInternet) TableName() string {
	return "goods_internet"
}
