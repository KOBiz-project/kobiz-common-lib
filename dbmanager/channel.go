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
	AssignmentMethod int8      `gorm:"column:assignment_method;not null;type:tinyint;default:0;comment:'배정 방법(0: 자동 순차, 1: 수동, 2: 팀 배정)'" json:"assignmentMethod"`
	ActiveYn         bool      `gorm:"column:active_yn;not null;type:boolean;default:true;comment:'활성 여부'" json:"activeYn"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (Channel) TableName() string {
	return "channels"
}

// 채널 팀 1:N 설정 테이블
type ChannelTeam struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	TeamID    uint      `gorm:"column:team_id;not null;type:int;comment:'팀 ID';index" json:"teamId"`
	CreatedAt time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelTeam) TableName() string {
	return "channel_teams"
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

// 인입 채널 테이블 (채널 당 각 게시판)
type InqueryChannel struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID      uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	Name           string    `gorm:"column:name;not null;type:varchar(100);comment:'인입 채널명'" json:"name"`
	Code           string    `gorm:"column:code;not null;type:varchar(100);comment:'채널 코드';uniqueIndex:idx_inquery_channels_code" json:"code"`
	MainImageUrl   string    `gorm:"column:main_image_url;type:varchar(255);comment:'메인이미지 주소'" json:"mainImageUrl"`
	SubImageUrl    string    `gorm:"column:sub_image_url;type:varchar(255);comment:'서브이미지 주소'" json:"subImageUrl"`
	Description    string    `gorm:"column:description;type:varchar(255);comment:'채널 설명'" json:"description"`
	Title          string    `gorm:"column:title;not null;type:varchar(100);comment:'타이틀'" json:"title"`
	InputName      string    `gorm:"column:input_name;not null;type:varchar(100);comment:'고객명 입력 항목명'" json:"inputName"`
	InputNameDesc  string    `gorm:"column:input_name_desc;type:varchar(255);comment:'고객명 입력 항목 설명'" json:"inputNameDesc"`
	InputPhone     string    `gorm:"column:input_phone;not null;type:varchar(100);comment:'고객 연락처 입력 항목명'" json:"inputPhone"`
	InputPhoneDesc string    `gorm:"column:input_phone_desc;type:varchar(255);comment:'고객 연락처 입력 항목 설명'" json:"inputPhoneDesc"`
	ButtonText     string    `gorm:"column:button_text;type:varchar(255);comment:'버튼 문구'" json:"buttonText"`
	ActiveYn       bool      `gorm:"column:active_yn;not null;type:boolean;default:true;comment:'활성 여부'" json:"activeYn"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (InqueryChannel) TableName() string {
	return "inquery_channels"
}

// 신청서 항목 설정 테이블 (삭제)
type ApplicationItem struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID    uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	Name         string    `gorm:"column:name;not null;type:varchar(100);comment:'항목명'" json:"name"`
	Description  string    `gorm:"column:description;type:varchar(255);comment:'항목 설명'" json:"description"`
	Type         int8      `gorm:"column:type;not null;type:tinyint;comment:'항목 타입(1: 고객명, 2: 고객 연락처)'" json:"type"`
	TypeRequired bool      `gorm:"column:type_required;not null;type:boolean;default:false;comment:'항목 필수 여부'" json:"typeRequired"`
	ActiveYn     bool      `gorm:"column:active_yn;not null;type:boolean;default:true;comment:'활성 여부'" json:"activeYn"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ApplicationItem) TableName() string {
	return "application_items"
}
