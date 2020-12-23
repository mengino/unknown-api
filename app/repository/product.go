package repository

import (
	"dlsite/app/model"
	"dlsite/internal/db"
)

// GetProductList 获取App列表
func GetProductList(group, categoryID int, sort string) []model.Product {
	var productList []model.Product

	db := db.New().Where(&model.Product{
		Base: model.Base{
			Status: true,
		},
		Group:      group,
		CategoryID: categoryID,
	})

	db.Order(sort + " desc, created_at desc").Find(&productList)

	return productList
}

// GetProduct 获取App
func GetProduct(id int) model.Product {
	var product model.Product

	db.New().Preload("Category").Where("id = ?", id).Find(&product)

	return product
}
