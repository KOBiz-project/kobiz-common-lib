package dbmanager

import (
	"time"
)

// ChannelBbs 채널 BBS 테이블
type ChannelBbs struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID      int64     `gorm:"column:channel_id;not null;type:bigint;comment:'채널 ID';index:idx_channel_bbss_channel_id" json:"channelId"`
	CustomerName   string    `gorm:"column:customer_name;not null;type:varchar(100);comment:'고객명'" json:"customerName"`
	CustomerNumber string    `gorm:"column:customer_number;not null;type:varchar(20);comment:'고객 연락처'" json:"customerNumber"`
	IdAdminUser    *int64    `gorm:"column:id_admin_user;type:bigint;comment:'배정 관리자 사용자 ID'" json:"idAdminUser"`
	CallDate       *time.Time `gorm:"column:call_date;type:datetime;comment:'통화일'" json:"callDate"`
	CallDuration   *int64    `gorm:"column:call_duration;type:bigint;comment:'통화시간(초)'" json:"callDuration"`
	Status         int8      `gorm:"column:status;not null;type:tinyint;default:0;comment:'상태(0: 신규, 1: 진행중, 2: 완료, 3: 실패)'" json:"status"`
	Gender         int8      `gorm:"column:gender;not null;type:tinyint;default:0;comment:'성별(0: 남자, 1: 여자)'" json:"gender"`
	Age            *int64    `gorm:"column:age;type:bigint;comment:'나이'" json:"age"`
	Memo           *string   `gorm:"column:memo;type:text;comment:'메모'" json:"memo"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelBbs) TableName() string {
	return "channel_bbss"
}

// 인입 원본 DB 데이터
type ChannelBbsOrigin struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID        uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	CustomerName     string    `gorm:"column:customer_name;not null;type:varchar(100);comment:'고객명'" json:"customerName"`
	CustomerNumber   string    `gorm:"column:customer_number;not null;type:varchar(20);comment:'고객 연락처'" json:"customerNumber"`
	InqueryTime      time.Time `gorm:"column:inquery_time;not null;type:datetime;comment:'인입시간'" json:"inqueryTime"`
	AssignmentStatus int8      `gorm:"column:assignment_status;not null;type:tinyint;default:0;comment:'배정상태(0: 배정 대기(신규 인입), 1: 팀 배정 완료, 2: 중복으로 인한 제외, 3: 관리자 삭제, 4: 배정보류)'" json:"assignmentStatus"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelBbsOrigin) TableName() string {
	return "channel_bbss_origins"
}

// 인입 원본 DB 상태 변경 로그
type ChannelBbsOriginLog struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IdChannelBbsOrigin uint      `gorm:"column:id_channel_bbs_origin;not null;type:int;comment:'원본 데이터 ID';index" json:"idChannelBbsOrigin"`
	PrevStatus         int8      `gorm:"column:prev_status;not null;type:tinyint;comment:'이전 배정상태(0: 배정 대기(신규 인입), 1: 팀 배정 완료, 2: 중복으로 인한 제외, 3: 관리자 삭제, 4: 배정보류)'" json:"prevStatus"`
	NewStatus          int8      `gorm:"column:new_status;not null;type:tinyint;comment:'변경 배정상태(0: 배정 대기(신규 인입), 1: 팀 배정 완료, 2: 중복으로 인한 제외, 3: 관리자 삭제, 4: 배정보류)'" json:"newStatus"`
	ChangedBy          string    `gorm:"column:changed_by;not null;type:varchar(100);comment:'변경자(시스템 또는 관리자명)'" json:"changedBy"`
	ChangeReason       string    `gorm:"column:change_reason;type:varchar(255);comment:'변경 사유'" json:"changeReason"`
	CreatedAt          time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
}

func (ChannelBbsOriginLog) TableName() string {
	return "channel_bbss_origin_logs"
}

// 팀별 DB 테이블
type ChannelBbsTeam struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IDChannelBbsOrigin uint      `gorm:"column:id_channel_bbs_origin_data;not null;type:int;comment:'원본 데이터 ID';index" json:"idChannelBbsOrigin"`
	ChannelID          uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	CustomerName       string    `gorm:"column:customer_name;not null;type:varchar(100);comment:'고객명'" json:"customerName"`
	CustomerNumber     string    `gorm:"column:customer_number;not null;type:varchar(20);comment:'고객 연락처'" json:"customerNumber"`
	InqueryTime        time.Time `gorm:"column:inquery_time;not null;type:datetime;comment:'인입시간'" json:"inqueryTime"`
	IdAdminTeam        uint      `gorm:"column:id_admin_team;not null;type:int;comment:'관리자 팀 ID'" json:"idAdminTeam"`
	IdAdminMember      uint      `gorm:"column:id_admin_member;not null;type:int;comment:'관리자 멤버 ID'" json:"idAdminMember"`
	AssiginmentStatus  int8      `gorm:"column:assignment_status;not null;type:tinyint;default:0;comment:'배정상태(0: 개인 배정 대기, 1: 개인 배정 완료, 2: 재배정됨, 3: 팀 보류, 4: 팀에서 삭제)'" json:"assignmentStatus"`
	CreatedAt          time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelBbsTeam) TableName() string {
	return "channel_bbs_teams"
}

// 팀별 DB 상태 변경 로그
type ChannelBbsTeamLog struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IdChannelBbsTeam uint      `gorm:"column:id_channel_bbs_team;not null;type:int;comment:'팀별 데이터 ID';index" json:"idChannelBbsTeam"`
	PrevStatus       int8      `gorm:"column:prev_status;not null;type:tinyint;comment:'이전 배정상태(0: 개인 배정 대기, 1: 개인 배정 완료, 2: 재배정됨, 3: 팀 보류, 4: 팀에서 삭제)'" json:"prevStatus"`
	NewStatus        int8      `gorm:"column:new_status;not null;type:tinyint;comment:'변경 배정상태(0: 개인 배정 대기, 1: 개인 배정 완료, 2: 재배정됨, 3: 팀 보류, 4: 팀에서 삭제)'" json:"newStatus"`
	ChangedBy        string    `gorm:"column:changed_by;not null;type:varchar(100);comment:'변경자(시스템 또는 관리자명)'" json:"changedBy"`
	ChangeReason     string    `gorm:"column:change_reason;type:varchar(255);comment:'변경 사유'" json:"changeReason"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
}

func (ChannelBbsTeamLog) TableName() string {
	return "channel_bbs_team_logs"
}

// 개인별(팀원별) 테이블
type ChannelBbsMember struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IdChannelBbsOrigin uint      `gorm:"column:id_channel_bbs_origin;not null;type:int;comment:'원본 데이터 ID';index" json:"idChannelBbsOrigin"`
	IDChannelBbsTeam   uint      `gorm:"column:id_channel_bbs_team;not null;type:int;comment:'팀별 데이터 ID';index" json:"idChannelBbsTeam"`
	ChannelID          uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	CustomerName       string    `gorm:"column:customer_name;not null;type:varchar(100);comment:'고객명'" json:"customerName"`
	CustomerNumber     string    `gorm:"column:customer_number;not null;type:varchar(20);comment:'고객 연락처'" json:"customerNumber"`
	InqueryTime        time.Time `gorm:"column:inquery_time;not null;type:datetime;comment:'인입시간'" json:"inqueryTime"`
	IdAdminTeam        uint      `gorm:"column:id_admin_team;not null;type:int;comment:'관리자 팀 ID'" json:"idAdminTeam"`
	IdAdminMember      uint      `gorm:"column:id_admin_member;not null;type:int;comment:'관리자 멤버 ID'" json:"idAdminMember"`
	AssignmentStatus   int8      `gorm:"column:assignment_status;not null;type:tinyint;default:0;comment:'배정상태(0: 신규 배정, 1: 상담 진행 중, 2: 재통화 예정, 3: 부재중, 4: 상담 보류, 5: 상담 완료, 6: 상담 취소)'" json:"assignmentStatus"`
	CallDate           time.Time `gorm:"column:call_date;type:datetime;comment:'통화일'" json:"callDate"`
	CallDuration       int       `gorm:"column:call_duration;type:int;comment:'통화 시간(초)'" json:"callDuration"`
	Gender             int8      `gorm:"column:gender;not null;type:tinyint;default:0;comment:'성별(0: 남자, 1: 여자)'" json:"gender"`
	Age                int       `gorm:"column:age;type:int;comment:'나이'" json:"age"`
	Memo               string    `gorm:"column:memo;type:text;comment:'메모'" json:"memo"`
	CreatedAt          time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelBbsMember) TableName() string {
	return "channel_bbs_members"
}

// 개인별(팀원별) 상태 변경 로그
type ChannelBbsMemberLog struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	IdChannelBbsMember uint      `gorm:"column:id_channel_bbs_member;not null;type:int;comment:'개인별 데이터 ID';index" json:"idChannelBbsMember"`
	PrevStatus         int8      `gorm:"column:prev_status;not null;type:tinyint;comment:'이전 배정상태(0: 신규 배정, 1: 상담 진행 중, 2: 재통화 예정, 3: 부재중, 4: 상담 보류, 5: 상담 완료, 6: 상담 취소)'" json:"prevStatus"`
	NewStatus          int8      `gorm:"column:new_status;not null;type:tinyint;comment:'변경 배정상태(0: 신규 배정, 1: 상담 진행 중, 2: 재통화 예정, 3: 부재중, 4: 상담 보류, 5: 상담 완료, 6: 상담 취소)'" json:"newStatus"`
	ChangedBy          string    `gorm:"column:changed_by;not null;type:varchar(100);comment:'변경자(시스템 또는 관리자명)'" json:"changedBy"`
	ChangeReason       string    `gorm:"column:change_reason;type:varchar(255);comment:'변경 사유'" json:"changeReason"`
	CreatedAt          time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
}

func (ChannelBbsMemberLog) TableName() string {
	return "channel_bbs_member_logs"
}

// 중복 테이블
type ChannelBbsDuplicate struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ChannelID      uint      `gorm:"column:channel_id;not null;type:int;comment:'채널 ID';index" json:"channelId"`
	CustomerName   string    `gorm:"column:customer_name;not null;type:varchar(100);comment:'고객명'" json:"customerName"`
	CustomerNumber string    `gorm:"column:customer_number;not null;type:varchar(20);comment:'고객 연락처'" json:"customerNumber"`
	InqueryTime    time.Time `gorm:"column:inquery_time;not null;type:datetime;comment:'인입시간'" json:"inqueryTime"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (ChannelBbsDuplicate) TableName() string {
	return "channel_bbss_duplicates"
}
