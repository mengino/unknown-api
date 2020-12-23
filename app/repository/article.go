package repository

import (
	"dlsite/app/model"
	"dlsite/internal/db"

	"github.com/jinzhu/gorm"
)

// GetArticleList 获取新闻列表
func GetArticleList(categoryID int) []model.Article {
	var articleList []model.Article

	db.New().Where(&model.Article{
		Base: model.Base{
			Status: true,
		},
		CategoryID: categoryID,
	}).Order("sort desc, created_at desc").Find(&articleList)

	return articleList
}

// GetArticle 获取文章
func GetArticle(id int) model.Article {
	var Article model.Article

	db.New().Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Category")
	}).Where("id = ?", id).Find(&Article)

	return Article
}
