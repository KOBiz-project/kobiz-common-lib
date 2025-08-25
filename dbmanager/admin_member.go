package dbmanager

import "gorm.io/gorm"

type AdminMember struct {
	gorm.Model
	UserId       string `gorm:"type:varchar(50);unique;comment:'사용자 ID'" json:"userId"`
	UserLevel    int8   `gorm:"column:level;not null;default:4;comment:'등급(1:최고관리자,2:일반관리자,3:팀관리자,4:팀원)'" json:"userLevel"`
	Password     string `gorm:"type:varchar(1024);comment:'비밀번호(암호화)'" json:"password"`
	ActiveYn     bool   `gorm:"default:false;comment:'활성화 여부'" json:"activeYn"`
	Name         string `gorm:"type:varchar(50);comment:'실명'" json:"name"`
	NickName     string `gorm:"type:varchar(50);unique;column:nick_name;comment:'닉네임'" json:"nickName"`
	Email        string `gorm:"type:varchar(255);unique;column:email;comment:'이메일 주소'" json:"email"`
	Phone        string `gorm:"type:varchar(20);unique;column:phone;comment:'전화번호'" json:"phone"`
	ProfileImage string `gorm:"type:varchar(1024);comment:'프로필 이미지 URL'" json:"profileImage"`
	IdAdminTeam  uint   `gorm:"column:id_admin_team;comment:'관리자 팀 번호'" json:"adminTeam"`
}

func (AdminMember) TableName() string {
	return "admin_members"
}
