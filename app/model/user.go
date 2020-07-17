package model

// User 用户模型
type User struct {
	Base
	Nickname string `json:"nickname" gorm:"type:varchar(32);not null"`
	Password string `json:"-" gorm:"type:varchar(64);not null"`
}
