package model

// Article 产品模型
type Article struct {
	Base
	Title      string `json:"title" gorm:"type:varchar(64);not null"`
	Image      Image  `json:"image" gorm:"default:''"`
	ProductID  int    `json:"product_id" gorm:"not null" sql:"index"`
	CategoryID int    `json:"category_id" gorm:"not null" sql:"index"`
	Content    string `json:"content" gorm:"type:text;not null"`
	Sort       int    `json:"sort" gorm:"not null;default:0" sql:"index"`
	Hot        int    `json:"hot" gorm:"not null;default:0" sql:"index"`
	Top        bool   `json:"top" gorm:"not null;default:false" sql:"index"`

	Category Category `json:"category" gorm:"foreignkey:CategoryID"`
	Product  Product  `json:"product" gorm:"foreignkey:ProductID"`
}
