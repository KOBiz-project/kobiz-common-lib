package dbmanager

import (
	"time"
)

// Channel 채널 설정 테이블
type Channel struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name             string    `gorm:"column:name;not null;type:varchar(100);comment:'채널명';uniqueIndex:idx_channel_name" json:"name"`
	Code             string    `gorm:"column:code;not null;type:varchar(100);comment:'채널 코드';uniqueIndex:idx_channel_code" json:"code"`
	Description      string    `gorm:"column:description;type:varchar(255);comment:'채널 설명'" json:"description"`
	IdAdminTeam      uint      `gorm:"column:id_admin_team;not null;type:int;comment:'배정 관리자 팀 ID'" json:"idAdminTeam"`
	AssignmentMethod int8      `gorm:"column:assignment_method;not null;type:tinyint;default:0;comment:'배정 방법(0: 자동 순차, 1: 수동, 2: 팀 배정)'" json:"assignmentMethod"`
	RepeatDbCheckYn  bool      `gorm:"column:repeat_db_check_yn;not null;type:boolean;default:false;comment:'중복 DB 체크 여부'" json:"repeatDbCheckYn"`
	Title            string    `gorm:"column:title;not null;type:varchar(100);comment:'타이틀'" json:"title"`
	MainImageUrl     string    `gorm:"column:main_image_url;type:varchar(255);comment:'메인이미지 주소'" json:"mainImageUrl"`
	SubImageUrl      string    `gorm:"column:sub_image_url;type:varchar(255);comment:'서브이미지 주소'" json:"subImageUrl"`
	ButtonText       string    `gorm:"column:button_text;type:varchar(255);comment:'버튼 문구'" json:"buttonText"`
	ActiveYn         bool      `gorm:"column:active_yn;not null;type:boolean;default:true;comment:'활성 여부'" json:"activeYn"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (Channel) TableName() string {
	return "channels"
}

// ChannelGroup 채널 그룹 설정 테이블
type ChannelGroup struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name;not null;type:varchar(100);comment:'그룹명'" json:"name"`
	Description string    `gorm:"column:description;type:varchar(255);comment:'그룹 설명'" json:"description"`
	ActiveYn    bool      `gorm:"column:active_yn;not null;type:boolean;default:true;comment:'활성 여부'" json:"activeYn"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelGroup) TableName() string {
	return "channel_groups"
}

// Channel ChannelGroup 중간 테이블
type ChannelChannelGroup struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID      uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	ChannelGroupID uint      `gorm:"column:channel_group_id;not null;type:int;comment:'그룹 ID';index" json:"channelGroupId"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelChannelGroup) TableName() string {
	return "channel_channel_groups"
}

// 신청서 항목 설정 테이블
type ApplicationItem struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID    uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	Name         string    `gorm:"column:name;not null;type:varchar(100);comment:'항목명'" json:"name"`
	Description  string    `gorm:"column:description;type:varchar(255);comment:'항목 설명'" json:"description"`
	Type         int8      `gorm:"column:type;not null;type:tinyint;comment:'항목 타입(1: 고객명, 2: 고객 연락처, 2: 날짜, 3: 체크박스, 4: 라디오, 5: 드롭다운)'" json:"type"`
	TypeRequired bool      `gorm:"column:type_required;not null;type:boolean;default:false;comment:'항목 필수 여부'" json:"typeRequired"`
	ActiveYn     bool      `gorm:"column:active_yn;not null;type:boolean;default:true;comment:'활성 여부'" json:"activeYn"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ApplicationItem) TableName() string {
	return "application_items"
}

type ChannelBbs struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID      uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	CustomerName   string    `gorm:"column:customer_name;not null;type:varchar(100);comment:'고객명'" json:"customerName"`
	CustomerNumber string    `gorm:"column:customer_number;not null;type:varchar(20);comment:'고객 연락처'" json:"customerNumber"`
	IdAdminUser    uint      `gorm:"column:id_admin_user;type:int;comment:'배정 관리자 사용자 ID'" json:"idAdminUser"`
	CallDate       time.Time `gorm:"column:call_date;type:datetime;comment:'통화일'" json:"callDate"`
	CallDuration   int       `gorm:"column:call_duration;type:int;comment:'통화시간(초)'" json:"callDuration"`
	Status         int8      `gorm:"column:status;not null;type:tinyint;default:0;comment:'상태(0: 신규, 1: 진행중, 2: 완료, 3: 실패)'" json:"status"`
	Gender         int8      `gorm:"column:gender;not null;type:tinyint;default:0;comment:'성별(0: 남자, 1: 여자)'" json:"gender"`
	Age            int       `gorm:"column:age;type:int;comment:'나이'" json:"age"`
	Memo           string    `gorm:"column:memo;type:text;comment:'메모'" json:"memo"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelBbs) TableName() string {
	return "channel_bbss"
}
