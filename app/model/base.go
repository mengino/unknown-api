package model

import (
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
