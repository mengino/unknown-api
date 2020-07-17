package category

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
	Name   string            `form:"name"`
	Group  int               `form:"group"`
	Sorter map[string]string `form:"sorter"`
}

type createOrUpdateRequest struct {
	Name  string `json:"name" binding:"required"`
	Group int    `json:"group" binding:"required"`
	Sort  int    `json:"sort"`
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

	var category []model.Category

	db := db.New().Where(&model.Category{
		Name:  request.Name,
		Group: request.Group,
	})

	if len(request.Sorter) > 0 {
		for k, v := range request.Sorter {
			db = db.Order(util.BuilderOrder(k, v))
		}
	} else {
		db = db.Order("created_at desc")
	}

	if err := db.Find(&category).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.Success(c, category)
	}
}

// Detail 分类详情
func Detail(c *gin.Context) {
	var category model.Category

	if db.New().Where("id = ?", c.Param("id")).First(&category).RecordNotFound() {
		response.Success(c, struct{}{})
		return
	}

	response.Success(c, category)
}

// Create 创建分类
func Create(c *gin.Context) {
	var request createOrUpdateRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	category := &model.Category{
		Name:  request.Name,
		Group: request.Group,
		Sort:  request.Sort,
	}

	if err := db.New().Create(&category).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.SuccessCreated(c, category)
	}
}

// Update 更新分类
func Update(c *gin.Context) {
	var request createOrUpdateRequest
	if err := c.ShouldBind(&request); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var category model.Category
	if db.New().Where("id = ?", c.Param("id")).First(&category).RecordNotFound() {
		response.Fail(c, http.StatusBadRequest, "找不到该记录")
		return
	}

	if err := db.New().Model(&category).Updates(structs.Map(&request)).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.Success(c, category)
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

	var category model.Category

	if err := db.New().Where(request.Key).Delete(&category).Error; err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
	} else {
		response.SuccessNoContent(c)
	}
}
