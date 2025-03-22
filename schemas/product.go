package schemas

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ServiceProvider   int8   `gorm:"column:service_provider;not null;type:smallint;comment:'통신사분류 1/2/3/4 [SKT/KT/LGU+/INTERNET]';uniqueIndex:idx_provider_model" json:"serviceProvider"`
	NameModel         string `gorm:"column:name_model;not null;type:varchar(256);comment:'모델 명';uniqueIndex:idx_provider_model" json:"nameModel"`
	NameProduct       string `gorm:"column:name_product;not null;type:varchar(256);comment:'상품 명'" json:"nameProduct"`
	Desc              string `gorm:"column:desc;not null;type:varchar(2000);comment:'상품 설명'" json:"desc"`
	ChangeNumber      string `gorm:"column:change_number;type:varchar(256);comment:'번호이동 내용'" json:"changeNumber"`
	ChangeDevice      string `gorm:"column:change_device;type:varchar(256);comment:'기기변경 내용'" json:"changeDevice"`
	ChangeNumberPrice int    `gorm:"column:change_number_price;comment:'번호이동 가격'" json:"changeNumberPrice"`
	ChangeDevicePrice int    `gorm:"column:change_device_price;comment:'기기변경 가격'" json:"changeDevicePrice"`
	DevicePhoto       string `gorm:"column:device_photo;type:varchar(1000);comment:'디바이스 이미지 경로'" json:"devicePhoto"`
	ActiveYn          bool   `gorm:"column:active_yn;default:false" json:"activeYn"`
}
