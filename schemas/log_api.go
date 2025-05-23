package schemas

import "time"

type LogApi struct {
	ID                     int64     `gorm:"primaryKey;autoIncrement;comment:'아이디'" json:"id"`
	CreatedAt              time.Time `gorm:"autoCreateTime;comment:'생성 날짜'" json:"created_at"`
	Method                 string    `gorm:"column:method;type:varchar(50);comment:'HTTP 메소드'" json:"method"`
	MethodUrl              string    `gorm:"column:method_url;type:varchar(255);comment:'요청 URL'" json:"methodUrl"`
	ClientIP               string    `gorm:"column:client_ip;type:varchar(50);comment:'클라이언트 IP 주소'" json:"clientIP"`
	IDUser                 int       `gorm:"column:id_user;comment:'사용자 ID'" json:"id_user"`
	ReqBody                string    `gorm:"column:req_body;type:varchar(1000);comment:'요청 본문'" json:"reqBody"`
	ResResultCode          int       `gorm:"column:res_result_code;comment:'응답 결과 코드'" json:"resResultCode"`
	ResResultDesc          string    `gorm:"column:res_result_desc;type:varchar(255);comment:'응답 결과 설명'" json:"resResultDesc"`
	ResData                string    `gorm:"column:res_data;type:varchar(9000);comment:'응답 데이터'" json:"resData"`
	Duration               int64     `gorm:"column:duration;comment:'요청 처리 시간(ms)'" json:"duration"`
	ReqHeaderAcceptService string    `gorm:"column:req_header_accept_service;type:varchar(255);comment:'Accept-Service 헤더'" json:"reqHeaderAcceptService"`
	ReqHeaderDeviceInfo    string    `gorm:"column:req_header_device_info;type:varchar(255);comment:'Device-Info 헤더'" json:"reqHeaderDeviceInfo"`
	ReqHeaderWearableInfo  string    `gorm:"column:req_header_wearable_info;type:varchar(255);comment:'Wearable-Info 헤더'" json:"reqHeaderWearableInfo"`
	ReqHeaderPackageInfo   string    `gorm:"column:req_header_package_info;type:varchar(255);comment:'Package-Info 헤더'" json:"reqHeaderPackageInfo"`
	ReqHeaderLocaleInfo    string    `gorm:"column:req_header_locale_info;type:varchar(255);comment:'Locale-Info 헤더'" json:"reqHeaderLocaleInfo"`
	ReqHeaderUserAgent     string    `gorm:"column:req_header_user_agent;type:varchar(255);comment:'User-Agent 헤더'" json:"reqHeaderUserAgent"`
	ReqHeaderContentType   string    `gorm:"column:req_header_content_type;type:varchar(255);comment:'Content-Type 헤더'" json:"reqHeaderContentType"`
	ReqHeaderAuthorization string    `gorm:"column:req_header_authorization;type:varchar(255);comment:'Authorization 헤더'" json:"reqHeaderAuthorization"`
}
