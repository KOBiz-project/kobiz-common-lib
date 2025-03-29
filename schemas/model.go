package schemas

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	NameModel     string `gorm:"column:name_model;not null;type:varchar(256);comment:'모델 명'" json:"nameModel"`
	Description   string `gorm:"column:description;type:text;comment:'모델 설명'" json:"description"`
	IdCompany     int    `gorm:"column:id_company;not null;comment:'제조사 고유 아이디'" json:"idCompany"`
	NameCompany   string `gorm:"column:name_company;not null;type:varchar(256);comment:'제조사 명'" json:"nameCompany"`
	Size64G       bool   `gorm:"column:size_64g;default:false;comment:'용량 64GB 여부'" json:"size64G"`
	Size64GCDSKT  string `gorm:"column:size_64g_cd_skt;type:varchar(256);comment:'용량 64GB SKT 모델 고유코드 (model_cd)'" json:"size64GCDSKT"`
	Size64GCDKT   string `gorm:"column:size_64g_cd_kt;type:varchar(256);comment:'용량 64GB KT 모델 고유코드 (prod_no)'" json:"size64GCDKT"`
	Size64GCDLGT  string `gorm:"column:size_64g_cd_lgt;type:varchar(256);comment:'용량 64GB LGT 모델 고유코드 (urc_trm_mdl_cd)'" json:"size64GCDLGT"`
	Size128G      bool   `gorm:"column:size_128g;default:false;comment:'용량 128GB 여부'" json:"size128G"`
	Size128GCDSKT string `gorm:"column:size_128g_cd_skt;type:varchar(256);comment:'용량 128GB SKT 모델 고유코드'" json:"size128GCDSKT"`
	Size128GCDKT  string `gorm:"column:size_128g_cd_kt;type:varchar(256);comment:'용량 128GB KT 모델 고유코드'" json:"size128GCDKT"`
	Size128GCDLGT string `gorm:"column:size_128g_cd_lgt;type:varchar(256);comment:'용량 128GB LGT 모델 고유코드'" json:"size128GCDLGT"`
	Size256G      bool   `gorm:"column:size_256g;default:false;comment:'용량 256GB 여부'" json:"size256G"`
	Size256GCDSKT string `gorm:"column:size_256g_cd_skt;type:varchar(256);comment:'용량 256GB SKT 모델 고유코드'" json:"size256GCDSKT"`
	Size256GCDKT  string `gorm:"column:size_256g_cd_kt;type:varchar(256);comment:'용량 256GB KT 모델 고유코드'" json:"size256GCDKT"`
	Size256GCDLGT string `gorm:"column:size_256g_cd_lgt;type:varchar(256);comment:'용량 256GB LGT 모델 고유코드'" json:"size256GCDLGT"`
	Size512G      bool   `gorm:"column:size_512g;default:false;comment:'용량 512GB 여부'" json:"size512G"`
	Size512GCDSKT string `gorm:"column:size_512g_cd_skt;type:varchar(256);comment:'용량 512GB SKT 모델 고유코드'" json:"size512GCDSKT"`
	Size512GCDKT  string `gorm:"column:size_512g_cd_kt;type:varchar(256);comment:'용량 512GB KT 모델 고유코드'" json:"size512GCDKT"`
	Size512GCDLGT string `gorm:"column:size_512g_cd_lgt;type:varchar(256);comment:'용량 512GB LGT 모델 고유코드'" json:"size512GCDLGT"`
	Size1T        bool   `gorm:"column:size_1t;default:false;comment:'용량 1TB 여부'" json:"size1T"`
	Size1TCDSKT   string `gorm:"column:size_1t_cd_skt;type:varchar(256);comment:'용량 1TB SKT 모델 고유코드'" json:"size1TCDSKT"`
	Size1TCDKT    string `gorm:"column:size_1t_cd_kt;type:varchar(256);comment:'용량 1TB KT 모델 고유코드'" json:"size1TCDKT"`
	Size1TCDLGT   string `gorm:"column:size_1t_cd_lgt;type:varchar(256);comment:'용량 1TB LGT 모델 고유코드'" json:"size1TCDLGT"`
	ActiveYn      bool   `gorm:"column:active_yn;default:false" json:"activeYn"`
}

func (Model) TableName() string {
	return "models"
}
