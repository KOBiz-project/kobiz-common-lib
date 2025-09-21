package dbmanager

import (
	"gorm.io/gorm"
)

type AdminMember struct {
	gorm.Model
	UserId            string `gorm:"type:varchar(50);unique;comment:'사용자 ID'" json:"userId"`
	Password          string `gorm:"type:varchar(1024);comment:'비밀번호(암호화)'" json:"password"`
	ActiveYn          bool   `gorm:"default:false;comment:'활성화 여부'" json:"activeYn"`
	Name              string `gorm:"type:varchar(50);comment:'실명'" json:"name"`
	NickName          string `gorm:"type:varchar(50);unique;column:nick_name;comment:'닉네임'" json:"nickName"`
	Email             string `gorm:"type:varchar(255);unique;column:email;comment:'이메일 주소'" json:"email"`
	Phone             string `gorm:"type:varchar(100);unique;column:phone;comment:'전화번호'" json:"phone"`
	ProfileImage      string `gorm:"type:varchar(1024);comment:'프로필 이미지 URL'" json:"profileImage"`
	IdAdminTeam       uint   `gorm:"column:id_admin_team;comment:'관리자 팀 번호'" json:"adminTeam"`
	IdLevelPermission uint   `gorm:"column:id_level_permission;comment:'레벨별 권한 번호'" json:"idLevelPermission"`
}

func (AdminMember) TableName() string {
	return "admin_members"
}

type AdminMemberLevelPermissionMenuPermission struct {
	gorm.Model
	IdAdminMember                   uint `gorm:"column:id_admin_member;comment:'관리자 멤버 번호'" json:"idAdminMember"`
	IdLevelPermissionMenuPermission uint `gorm:"column:id_level_permission_menu_permission;comment:'레벨별 권한 메뉴 권한 번호'" json:"idLevelPermissionMenuPermission"`
	ActiveYn                        bool `gorm:"column:use_yn;default:false;comment:'사용 여부'" json:"activeYn"`
}

func (AdminMemberLevelPermissionMenuPermission) TableName() string {
	return "admin_member_level_permission_menu_permissions"
}

type AdminTeam struct {
	gorm.Model
	TeamName string `gorm:"column:team_name;not null;type:varchar(255);comment:'팀 이름'" json:"teamName"`
	ActiveYn bool   `gorm:"column:active_yn;default:false;comment:'활성화 여부'" json:"activeYn"`
}

// TableName AdminTeam 테이블명 반환
func (AdminTeam) TableName() string {
	return "admin_teams"
}
