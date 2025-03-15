package schemas

import "time"

type Product struct {
	ID                int       `gorm:"primaryKey;column:id;autoIncrement" json:"id"`
	ServiceProvider   int8      `gorm:"column:service_provider;not null;type:smallint;comment:'통신사분류 1/2/3/4 [SKT/KT/LGU+/INTERNET]';uniqueIndex:idx_provider_model" json:"service_provider"`
	NameModel         string    `gorm:"column:name_model;not null;type:varchar(256);comment:'모델 명';uniqueIndex:idx_provider_model" json:"name_model"`
	NameProduct       string    `gorm:"column:name_product;not null;type:varchar(256);comment:'상품 명'" json:"name_product"`
	Desc              string    `gorm:"column:desc;not null;type:varchar(2000);comment:'상품 설명'" json:"desc"`
	ChangeNumber      string    `gorm:"column:change_number;type:varchar(256);comment:'번호이동 내용'" json:"change_number"`
	ChangeDevice      string    `gorm:"column:change_device;type:varchar(256);comment:'기기변경 내용'" json:"change_device"`
	ChangeNumberPrice int       `gorm:"column:change_number_price;comment:'번호이동 가격'" json:"change_number_price"`
	ChangeDevicePrice int       `gorm:"column:change_device_price;comment:'기기변경 가격'" json:"change_device_price"`
	DevicePhoto       string    `gorm:"column:device_photo;type:varchar(1000);comment:'디바이스 이미지 경로'" json:"device_photo"`
	ActiveYn          bool      `gorm:"column:active_yn;default:false" json:"active_yn"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at"`
}
