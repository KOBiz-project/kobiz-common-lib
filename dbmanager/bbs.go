package dbmanager

import (
	"time"
)

// InqueryChannelBbs 채널 문의 관리 테이블
type InqueryChannelBbs struct {
	ID                      uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	InqueryChannelID        int64      `gorm:"column:inquery_channel_id;not null;type:bigint;comment:'인입 채널 ID';index:idx_channel_bbss_inquery_channel_id" json:"inqueryChannelId"`
	AdConcept               *string    `gorm:"column:ad_concept;type:varchar(100);comment:'광고 컨셉'" json:"adConcept"`
	InqueryPath             *string    `gorm:"column:inquery_path;type:varchar(200);comment:'유입경로 상세'" json:"inqueryPath"`
	InqueryKeyword          *string    `gorm:"column:inquery_keyword;type:varchar(100);comment:'유입키워드'" json:"inqueryKeyword"`
	InqueryTime             time.Time  `gorm:"column:inquery_time;not null;type:datetime;comment:'유입시간'" json:"inqueryTime"`
	CustomerName            string     `gorm:"column:customer_name;not null;type:varchar(100);comment:'고객명'" json:"customerName"`
	CustomerNumber          string     `gorm:"column:customer_number;not null;type:varchar(20);comment:'전화번호'" json:"customerNumber"`
	AssignmentStatus        int8       `gorm:"column:assignment_status;not null;type:tinyint;default:0;comment:'배정상태(0: 배정대기, 1: 팀배정(자동), 2: 팀배정(수동), 3: 상담원배정(자동), 4: 상담원배정(수동), 5: 실패'" json:"assignmentStatus"`
	AssignedTeamID          *int64     `gorm:"column:assigned_team_id;type:bigint;comment:'배정팀 ID'" json:"assignedTeamId"`
	AssignmentTeamDate      *time.Time `gorm:"column:assignment_team_date;type:datetime;comment:'팀배정일'" json:"assignmentTeamDate"`
	AssignedAdminMemberID   *int64     `gorm:"column:assigned_admin_member_id;type:bigint;comment:'배정상담원 ID'" json:"assignedAdminMemberId"`
	AssignedAdminMemberDate *time.Time `gorm:"column:assigned_admin_member_date;type:datetime;comment:'팀원배정일'" json:"assignedAdminMemberDate"`
	ProgressStatus          int8       `gorm:"column:progress_status;type:tinyint;default:0;comment:'진행상태(0: 접수, 1: 통화, 2: 재통화, 3: 부재중, 4: 거부, 5: 잔여, 6: 마감, 7: 실패);comment:'진행상태'" json:"progressStatus"`
	CallDate                *time.Time `gorm:"column:call_date;type:datetime;comment:'통화일'" json:"callDate"`
	CallDuration            *int64     `gorm:"column:call_duration;type:bigint;comment:'통화시간(초)'" json:"callDuration"`
	Gender                  int8       `gorm:"column:gender;type:tinyint;default:0;comment:'성별(0: 남자, 1: 여자, 2: 기타)';comment:'성별'" json:"gender"`
	Age                     int8       `gorm:"column:age;type:tinyint;comment:'연령'" json:"age"`
	CurrentCarrier          *string    `gorm:"column:current_carrier;type:varchar(20);comment:'현재통신사'" json:"currentCarrier"`
	DesiredCarrier          *string    `gorm:"column:desired_carrier;type:varchar(20);comment:'희망통신사'" json:"desiredCarrier"`
	Memo                    *string    `gorm:"column:memo;type:text;comment:'메모'" json:"memo"`
	CreatedAt               time.Time  `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt               time.Time  `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (InqueryChannelBbs) TableName() string {
	return "inquery_channel_bbss"
}

// 중복 DB
type InqueryChannelBbsDuplicate struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	InqueryChannelBbsID int64     `gorm:"column:inquery_channel_bbs_id;not null;type:bigint;comment:'채널 문의 관리 ID';index:idx_inquery_channel_bbs_duplicates_inquery_channel_bbs_id" json:"inqueryChannelBbsId"`
	CreatedAt           time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt           time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (InqueryChannelBbsDuplicate) TableName() string {
	return "inquery_channel_bbs_duplicates"
}
