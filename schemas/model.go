package schemas

import (
	"gorm.io/gorm"
)

// 자체 모델 스키마
type Model struct {
	gorm.Model
	NameModel     string `gorm:"column:name_model;not null;type:varchar(256);comment:'모델 명';uniqueIndex:idx_model_name" json:"nameModel"`
	Description   string `gorm:"column:description;type:text;comment:'모델 설명'" json:"description"`
	IdCompany     int    `gorm:"column:id_company;not null;comment:'제조사 고유 아이디'" json:"idCompany"`
	NameCompany   string `gorm:"column:name_company;not null;type:varchar(256);comment:'제조사 명'" json:"nameCompany"`
	Size64G       bool   `gorm:"column:size_64g;default:false;comment:'용량 64GB 여부'" json:"size64G"`
	Size64GCDSKT  string `gorm:"column:size_64g_cd_skt;type:varchar(256);comment:'용량 64GB SKT 모델 고유코드 (model_cd)'" json:"size64GCDSKT"`
	Size64GCDKT   string `gorm:"column:size_64g_cd_kt;type:varchar(256);comment:'용량 64GB KT 모델 고유코드 (prod_no)'" json:"size64GCDKT"`
	Size64GCDLGU  string `gorm:"column:size_64g_cd_lgu;type:varchar(256);comment:'용량 64GB LGU 모델 고유코드 (urc_trm_mdl_cd)'" json:"size64GCDLGU"`
	Size128G      bool   `gorm:"column:size_128g;default:false;comment:'용량 128GB 여부'" json:"size128G"`
	Size128GCDSKT string `gorm:"column:size_128g_cd_skt;type:varchar(256);comment:'용량 128GB SKT 모델 고유코드'" json:"size128GCDSKT"`
	Size128GCDKT  string `gorm:"column:size_128g_cd_kt;type:varchar(256);comment:'용량 128GB KT 모델 고유코드'" json:"size128GCDKT"`
	Size128GCDLGU string `gorm:"column:size_128g_cd_lgu;type:varchar(256);comment:'용량 128GB LGU 모델 고유코드'" json:"size128GCDLGU"`
	Size256G      bool   `gorm:"column:size_256g;default:false;comment:'용량 256GB 여부'" json:"size256G"`
	Size256GCDSKT string `gorm:"column:size_256g_cd_skt;type:varchar(256);comment:'용량 256GB SKT 모델 고유코드'" json:"size256GCDSKT"`
	Size256GCDKT  string `gorm:"column:size_256g_cd_kt;type:varchar(256);comment:'용량 256GB KT 모델 고유코드'" json:"size256GCDKT"`
	Size256GCDLGU string `gorm:"column:size_256g_cd_lgu;type:varchar(256);comment:'용량 256GB LGU 모델 고유코드'" json:"size256GCDLGU"`
	Size512G      bool   `gorm:"column:size_512g;default:false;comment:'용량 512GB 여부'" json:"size512G"`
	Size512GCDSKT string `gorm:"column:size_512g_cd_skt;type:varchar(256);comment:'용량 512GB SKT 모델 고유코드'" json:"size512GCDSKT"`
	Size512GCDKT  string `gorm:"column:size_512g_cd_kt;type:varchar(256);comment:'용량 512GB KT 모델 고유코드'" json:"size512GCDKT"`
	Size512GCDLGU string `gorm:"column:size_512g_cd_lgu;type:varchar(256);comment:'용량 512GB LGU 모델 고유코드'" json:"size512GCDLGU"`
	Size1T        bool   `gorm:"column:size_1t;default:false;comment:'용량 1TB 여부'" json:"size1T"`
	Size1TCDSKT   string `gorm:"column:size_1t_cd_skt;type:varchar(256);comment:'용량 1TB SKT 모델 고유코드'" json:"size1TCDSKT"`
	Size1TCDKT    string `gorm:"column:size_1t_cd_kt;type:varchar(256);comment:'용량 1TB KT 모델 고유코드'" json:"size1TCDKT"`
	Size1TCDLGU   string `gorm:"column:size_1t_cd_lgu;type:varchar(256);comment:'용량 1TB LGU 모델 고유코드'" json:"size1TCDLGU"`
	ActiveYn      bool   `gorm:"column:active_yn;default:false" json:"activeYn"`
}

func (Model) TableName() string {
	return "models"
}
