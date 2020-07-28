package article

import (
	"dlsite/app/model"
	"dlsite/internal/db"
	"dlsite/internal/http/response"
	"dlsite/util"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type queryRequest struct {
	Title      string            `form:"title"`
	CategoryID int               `form:"category_id"`
	ProductID  int               `form:"product_id"`
	Filter     map[string]string `form:"filter"`
	Sorter     map[string]string `form:"sorter"`
}

type createOrUpdateRequest struct {
	Title      string      `json:"title" binding:"required"`
	CategoryID int         `json:"category_id" binding:"required"`
	ProductID  int         `json:"product_id" binding:"required"`
	Sort       int         `json:"sort,default=0" binding:"omitempty"`
	Image      model.Image `json:"image" binding:"required"`
	Content    string      `json:"content" binding:"omitempty"`
}

type deleteRequest struct {
	Key []int `json:"key" binding:"required"`
}

// List 分类列表
func List(c *gin.Context) {
	var request queryRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var article []model.Article

	db := db.New().Where(&model.Article{
		Title:      request.Title,
		CategoryID: request.CategoryID,
		ProductID:  request.ProductID,
	})

	if len(request.Sorter) > 0 {
		for k, v := range request.Sorter {
			db = db.Order(util.BuilderOrder(k, v))
		}
	} else {
		db = db.Order("created_at desc")
	}

	err := db.Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title")
	}).
		Find(&article).
		Error

	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.Success(c, article)
	}
}

// Detail 分类详情
func Detail(c *gin.Context) {
	var article model.Article

	record := db.New().Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, title")
	}).
		Where("id = ?", c.Param("id")).
		First(&article).
		RecordNotFound()

	if record {
		response.Success(c, struct{}{})
		return
	}

	response.Success(c, article)
}

// Create 创建分类
func Create(c *gin.Context) {
	var request createOrUpdateRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	article := &model.Article{
		Title:      request.Title,
		CategoryID: request.CategoryID,
		ProductID:  request.ProductID,
		Sort:       request.Sort,
		Image:      request.Image,
		Content:    request.Content,
	}

	if err := db.New().Create(&article).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.SuccessCreated(c, article)
	}
}

// Update 更新分类
func Update(c *gin.Context) {
	var request createOrUpdateRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var article model.Article
	if db.New().Where("id = ?", c.Param("id")).First(&article).RecordNotFound() {
		response.Fail(c, http.StatusBadRequest, "找不到该记录")
		return
	}

	if err := db.New().Model(&article).Updates(structs.Map(&request)).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.Success(c, article)
	}
}

// Delete 删除分类
func Delete(c *gin.Context) {
	var request deleteRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	if len(request.Key) <= 0 {
		response.Fail(c, http.StatusUnprocessableEntity, "请传入正确的数据")
		return
	}

	var article model.Article

	if err := db.New().Where(request.Key).Delete(&article).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.SuccessNoContent(c)
	}
}
