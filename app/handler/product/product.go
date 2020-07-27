package product

import (
	"dlsite/app/model"
	"dlsite/internal/db"
	"dlsite/internal/http/response"
	"dlsite/util"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type queryRequest struct {
	Title      string            `form:"title"`
	Group      int               `form:"group"`
	CategoryID int               `form:"category_id"`
	Filter     map[string]string `form:"filter"`
	Sorter     map[string]string `form:"sorter"`
}

type createOrUpdateRequest struct {
	Title      string                 `json:"title" binding:"required"`
	Group      int                    `json:"group" binding:"required"`
	Sort       int                    `json:"sort,default=0" binding:"omitempty"`
	CategoryID int                    `json:"category_id" binding:"required"`
	Content    string                 `json:"content" binding:"omitempty"`
	Image      string                 `json:"image" binding:"required"`
	Intro      string                 `json:"intro" binding:"omitempty"`
	Language   int                    `json:"language" binding:"omitempty"`
	Size       string                 `json:"size" binding:"omitempty"`
	Slide      model.ProductSlideJSON `json:"slide" binding:"omitempty"`
	URL        string                 `json:"url" binding:"omitempty"`
	Version    string                 `json:"version" binding:"omitempty"`
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

	var product []model.Product

	db := db.New().Where(&model.Product{
		Title:      request.Title,
		Group:      request.Group,
		CategoryID: request.CategoryID,
	})

	if len(request.Sorter) > 0 {
		for k, v := range request.Sorter {
			db = db.Order(util.BuilderOrder(k, v))
		}
	} else {
		db = db.Order("created_at desc")
	}

	if err := db.Find(&product).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.Success(c, product)
	}
}

// Detail 分类详情
func Detail(c *gin.Context) {
	var product model.Product

	if db.New().Where("id = ?", c.Param("id")).First(&product).RecordNotFound() {
		response.Success(c, struct{}{})
		return
	}

	response.Success(c, product)
}

// Create 创建分类
func Create(c *gin.Context) {
	var request createOrUpdateRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	product := &model.Product{
		Title:      request.Title,
		Group:      request.Group,
		Sort:       request.Sort,
		CategoryID: request.CategoryID,
		Content:    request.Content,
		Image:      request.Image,
		Intro:      request.Intro,
		Language:   request.Language,
		Size:       request.Size,
		Slide:      request.Slide,
		URL:        request.URL,
		Version:    request.Version,
	}

	if err := db.New().Create(&product).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.SuccessCreated(c, product)
	}
}

// Update 更新分类
func Update(c *gin.Context) {
	var request createOrUpdateRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var product model.Product
	if db.New().Where("id = ?", c.Param("id")).First(&product).RecordNotFound() {
		response.Fail(c, http.StatusBadRequest, "找不到该记录")
		return
	}

	if err := db.New().Model(&product).Updates(structs.Map(&request)).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.Success(c, product)
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

	var product model.Product

	if err := db.New().Where(request.Key).Delete(&product).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.SuccessNoContent(c)
	}
}
