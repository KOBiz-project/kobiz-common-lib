package dbmanager

import (
	"time"
)

type AdminMember struct {
	ModelCommon
	UserId       string `gorm:"type:varchar(50);unique;comment:'사용자 ID'" json:"userId"`
	Password     string `gorm:"type:varchar(1024);comment:'비밀번호(암호화)'" json:"password"`
	Name         string `gorm:"type:varchar(50);comment:'실명'" json:"name"`
	NickName     string `gorm:"type:varchar(50);unique;column:nick_name;comment:'닉네임'" json:"nickName"`
	Email        string `gorm:"type:varchar(255);unique;column:email;comment:'이메일 주소'" json:"email"`
	Phone        string `gorm:"type:varchar(20);unique;column:phone;comment:'전화번호'" json:"phone"`
	ProfileImage string `gorm:"type:varchar(1024);comment:'프로필 이미지 URL'" json:"profileImage"`
	IdAdminTeam  uint   `gorm:"column:id_admin_team;comment:'관리자 팀 번호'" json:"adminTeam"`
	RoleType     int8   `gorm:"column:role_type;type:tinyint;default:1;not null;comment:'관리자 역할 유형 (1~10 범위)'" json:"roleType"`
}

func (AdminMember) TableName() string {
	return "admin_members"
}

type AdminMemberPermission struct {
	ModelCommon
	IdAdminMember     uint `gorm:"column:id_admin_member;not null;comment:'관리자 번호'" json:"idAdminMember"`
	IdLevelPermission uint `gorm:"column:id_level_permission;not null;comment:'레벨별 권한 번호'" json:"idLevelPermission"`
}

func (AdminMemberPermission) TableName() string {
	return "admin_member_permissions"
}

type AdminTeam struct {
	ModelCommon
	TeamName string `gorm:"column:team_name;not null;type:varchar(255);comment:'팀 이름'" json:"teamName"`
}

// TableName AdminTeam 테이블명 반환
func (AdminTeam) TableName() string {
	return "admin_teams"
}

type ModelCommon struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ActiveYn    bool      `gorm:"column:active_yn;not null;type:boolean;default:true;comment:'활성 여부'" json:"activeYn"`
	Description string    `gorm:"column:description;type:varchar(255);default:'';comment:'설명'" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;type:datetime;comment:'생성일';default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;type:datetime;comment:'수정일';default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updatedAt"`
}
