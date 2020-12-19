package repository

import (
	"dlsite/app/model"
	"dlsite/internal/db"
)

// GetCategoryList 获取类型列表
func GetCategoryList(group int) []model.Category {
	var categoryList []model.Category

	db.New().Where("`group` = ? AND status = ?", group, 1).Find(&categoryList)

	return categoryList
}
