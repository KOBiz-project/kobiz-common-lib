package dbmanager

import (
	"time"
)

// Channel 채널 설정 테이블
type Channel struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IdChannelGroup   uint      `gorm:"column:id_channel_group;not null;type:int;default:0;comment:'채널 그룹 ID'" json:"idChannelGroup"`
	Name             string    `gorm:"column:name;not null;type:varchar(100);comment:'채널명';uniqueIndex:idx_channel_name" json:"name"`
	Description      string    `gorm:"column:description;type:varchar(255);comment:'채널 설명'" json:"description"`
	IdAdminTeam      uint      `gorm:"column:id_admin_team;not null;type:int;comment:'배정 관리자 팀 ID'" json:"idAdminTeam"`
	AssignmentMethod int8      `gorm:"column:assignment_method;not null;type:tinyint;default:0;comment:'배정 방법(0: 자동 순차, 1: 수동, 2: 팀 배정)'" json:"assignmentMethod"`
	RepeatDbCheckYn  bool      `gorm:"column:repeat_db_check_yn;not null;type:boolean;default:false;comment:'중복 DB 체크 여부'" json:"repeatDbCheckYn"`
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
