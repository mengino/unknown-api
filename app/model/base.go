package model

import (
	"database/sql/driver"
	"dlsite/internal/config"
	"encoding/json"
	"strings"
	"time"
)

// Base 基础模型
type Base struct {
	ID        int        `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at" sql:"index"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Status    bool       `json:"status" gorm:"type:tinyint(1);not null;default:false" sql:"index"`
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
