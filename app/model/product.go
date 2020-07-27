package model

import (
	"database/sql/driver"
	"dlsite/internal/config"
	"encoding/json"
	"strings"
)

// Product 产品模型
type Product struct {
	Base
	Title      string           `json:"title" gorm:"type:varchar(64);not null"`
	Image      Image            `json:"image" gorm:"default:''"`
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
	// Category   Category         `json:"category" gorm:"foreignkey:CategoryID"`
}

// Image 产品封面图
type Image string

// ProductSlideJSON 产品轮播图
type ProductSlideJSON []string

// Value 存入方法
func (s ProductSlideJSON) Value() (driver.Value, error) {
	domain := config.New().GetString("static_domain")
	for k, v := range s {
		if strings.Contains(string(v), domain) {
			s[k] = v[len(domain)+1:]
		}

	}

	b, err := json.Marshal(s)
	return string(b), err
}

// Scan 获取方法
func (s *ProductSlideJSON) Scan(input interface{}) error {
	if err := json.Unmarshal(input.([]byte), s); err != nil {
		return err
	}

	domain := config.New().GetString("static_domain")
	for k, v := range *s {
		(*s)[k] = domain + "/" + v
	}

	return nil
}

// Value 存入方法
func (i Image) Value() (driver.Value, error) {
	domain := config.New().GetString("static_domain")
	if strings.Contains(string(i), domain) {
		return string(i[len(domain)+1:]), nil
	}

	return string(i), nil
}

// Scan 获取方法
func (i *Image) Scan(input interface{}) error {
	domain := config.New().GetString("static_domain")
	*i = Image(domain + "/" + string(input.([]uint8)))

	return nil
}
