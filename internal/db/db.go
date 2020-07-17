package db

import (
	"dlsite/internal/config"

	"github.com/jinzhu/gorm"
)

// Conn mysql实例
var db *gorm.DB

// New 构造函数
func New() *gorm.DB {
	return db
}

// Init 初始化数据库
func Init() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open("mysql", config.New().GetString("dsn"))
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(40)
	db.DB().SetMaxOpenConns(100)

	return db, nil
}
