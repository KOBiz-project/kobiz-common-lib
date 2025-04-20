package schemas

import (
	"gorm.io/gorm"
)

// 자체 상품 스키마
type Goods struct {
	gorm.Model
	Telecom                 string `gorm:"column:telecom;not null;type:varchar(30);comment:'통신사'" json:"telecom"`
	IdModel                 int    `gorm:"column:id_model;not null;comment:'모델 고유 ID'" json:"idModel"`
	Size                    string `gorm:"column:size;not null;type:varchar(10);comment:'기기용량 64GB/128GB/256GB/512GB/1TB'" json:"size"`
	NumChg                  bool   `gorm:"column:num_chg;default:false;comment:'번호 이동 여부'" json:"numChg"`
	NumChgIdPlan            int    `gorm:"column:num_chg_id_plan;comment:'번호 이동 요금제 ID'" json:"numChgIdPlan"`
	NumChgSwitchSubsidy     int64  `gorm:"column:num_chg_switch_subsidy;type:bigint;comment:'번호 이동 전환지원금'" json:"numChgSwitchSubsidy"`
	NumChgAddSubsidy        int64  `gorm:"column:num_chg_add_subsidy;type:bigint;comment:'번호 이동 추가지원금'" json:"numChgAddSubsidy"`
	DeviceChg               bool   `gorm:"column:device_chg;default:false;comment:'기기 변경 여부'" json:"deviceChg"`
	DeviceChgIdPlan         int    `gorm:"column:device_chg_id_plan;comment:'기기 변경 요금제 ID'" json:"deviceChgIdPlan"`
	DeviceChgAddSubsidy     int64  `gorm:"column:device_chg_add_subsidy;type:bigint;comment:'기기 변경 추가지원금'" json:"deviceChgAddSubsidy"`
	NewActivation           bool   `gorm:"column:new_activation;default:false;comment:'신규 개통 여부'" json:"newActivation"`
	NewActivationIdPlan     int    `gorm:"column:new_activation_id_plan;comment:'신규 개통 요금제 ID'" json:"newActivationIdPlan"`
	NewActivationAddSubsidy int64  `gorm:"column:new_activation_add_subsidy;type:bigint;comment:'신규 개통 추가지원금'" json:"newActivationAddSubsidy"`
	SubsidyYn               bool   `gorm:"column:subsidy_yn;default:false;comment:'공시지원 여부'" json:"subsidyYn"`
	ActiveYn                bool   `gorm:"column:active_yn;default:false" json:"activeYn"`
}

func (Goods) TableName() string {
	return "goods"
}
