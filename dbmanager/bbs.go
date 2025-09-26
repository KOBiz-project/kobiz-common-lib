package dbmanager

import (
	"time"
)

// ChannelBbs 채널 BBS 테이블
type ChannelBbs struct {
	ID             uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID      int64      `gorm:"column:channel_id;not null;type:bigint;comment:'채널 ID';index:idx_channel_bbss_channel_id" json:"channelId"`
	CustomerName   string     `gorm:"column:customer_name;not null;type:varchar(100);comment:'고객명'" json:"customerName"`
	CustomerNumber string     `gorm:"column:customer_number;not null;type:varchar(20);comment:'고객 연락처'" json:"customerNumber"`
	InqueryTime    time.Time  `gorm:"column:inquery_time;not null;type:datetime;comment:'인입시간'" json:"inqueryTime"`
	IdAdminUser    *int64     `gorm:"column:id_admin_user;type:bigint;comment:'배정 관리자 사용자 ID'" json:"idAdminUser"`
	CallDate       *time.Time `gorm:"column:call_date;type:datetime;comment:'통화일'" json:"callDate"`
	CallDuration   *int64     `gorm:"column:call_duration;type:bigint;comment:'통화시간(초)'" json:"callDuration"`
	Status         int8       `gorm:"column:status;not null;type:tinyint;default:0;comment:'상태(0: 신규, 1: 진행중, 2: 완료, 3: 실패)'" json:"status"`
	Gender         int8       `gorm:"column:gender;not null;type:tinyint;default:0;comment:'성별(0: 남자, 1: 여자)'" json:"gender"`
	Age            *int64     `gorm:"column:age;type:bigint;comment:'나이'" json:"age"`
	Memo           *string    `gorm:"column:memo;type:text;comment:'메모'" json:"memo"`
	CreatedAt      time.Time  `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelBbs) TableName() string {
	return "channel_bbss"
}
