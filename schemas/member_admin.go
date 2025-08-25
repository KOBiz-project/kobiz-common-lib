package schemas

import "gorm.io/gorm"

type MemberAdmin struct {
	gorm.Model
	UserId       string `gorm:"type:varchar(50);unique;comment:'사용자 ID'" json:"userId"`
	Password     string `gorm:"type:varchar(1024);comment:'비밀번호(암호화)'" json:"password"`
	ActiveYn     bool   `gorm:"default:false;comment:'활성화 여부'" json:"activeYn"`
	Name         string `gorm:"type:varchar(50);comment:'실명'" json:"name"`
	NickName     string `gorm:"type:varchar(50);unique;column:nick_name;comment:'닉네임'" json:"nickName"`
	UserLevel    int8   `gorm:"default:5;comment:'등급(1:최고관리자,2:일반관리자,3:팀관리자,4:팀원,5:일반회원)'" json:"userLevel"`
	ProfileImage string `gorm:"type:varchar(1024);comment:'프로필 이미지 URL'" json:"profileImage"`
}

func (MemberAdmin) TableName() string {
	return "member_admins"
}
