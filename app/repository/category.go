package repository

import (
	"dlsite/app/model"
	"dlsite/internal/db"
)

// GetCategoryList 获取类型列表
func GetCategoryList(group int) []model.Category {
	var categoryList []model.Category

	db.New().Where(&model.Category{
		Base: model.Base{
			Status: true,
		},
		Group: group,
	}).Order("sort desc, created_at desc").Find(&categoryList)

	return categoryList
}
