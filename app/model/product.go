package model

import (
	"database/sql/driver"
	"encoding/json"
)

// Product 产品模型
type Product struct {
	Base
	Title      string           `json:"title" gorm:"type:varchar(64);not null"`
	Image      string           `json:"image" gorm:"default:''"`
	Group      int              `json:"group" gorm:"type:tinyint(1);not null" sql:"index"`
	CategoryID int              `json:"category_id" gorm:"not null" sql:"index"`
	Sort       int              `json:"sort" gorm:"not null;default:0" sql:"index"`
	Size       string           `json:"size" gorm:"type:varchar(64);default:''"`
	Version    string           `json:"version" gorm:"type:varchar(64);default:0"`
	Language   int              `json:"language" gorm:"type:tinyint(4);default:1"`
	Intro      string           `json:"intro" gorm:"type:text"`
	Content    string           `json:"content" gorm:"type:text"`
	Slide      ProductSlideJSON `json:"slide" gorm:"type:json"`
	URL        string           `json:"url" gorm:"default:''"`
	Hot        int              `json:"hot" gorm:"not null;default:0" sql:"index"`
	Top        bool             `json:"top" gorm:"not null;default:false" sql:"index"`
}

// ProductSlideJSON 产品轮播图
type ProductSlideJSON []string

// Value 存入方法
func (s *ProductSlideJSON) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}

// Scan 获取方法
func (s *ProductSlideJSON) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), s)
}
