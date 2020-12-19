package model

var (
	GroupGame       = 1
	GroupSoft       = 2
	GroupNews       = 3
	GroupCollection = 4
	GroupTop        = 5
)

// Category 分类模型
type Category struct {
	Base
	Name  string `json:"name" gorm:"type:varchar(64);not null"`
	Group int    `json:"group" gorm:"type:tinyint(1);not null" sql:"index"`
	Sort  int    `json:"sort" gorm:"not null;default:0" sql:"index"`
}
