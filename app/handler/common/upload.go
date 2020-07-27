package common

import (
	"dlsite/internal/http/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type fileResponse struct {
	FileName string `json:"file_name"`
	URL      string `json:"url"`
}

// Upload 上传图片
func Upload(c *gin.Context) {
	if file, err := c.FormFile("file"); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
	} else {
		// 上传文件至指定目录
		name := strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
		c.SaveUploadedFile(file, "./upload/"+name)

		response.Success(c, &fileResponse{FileName: name, URL: "http://localhost:30088/image/" + name})
	}
}

// MultiUpload 多图上传
func MultiUpload(c *gin.Context) {
	var resp []*fileResponse

	if form, err := c.MultipartForm(); err != nil {
		response.Fail(c, http.StatusUnprocessableEntity, err.Error())
	} else {
		for _, file := range form.File["files[]"] {
			name := strconv.FormatInt(time.Now().Unix(), 10) + "_" + file.Filename
			c.SaveUploadedFile(file, "./upload/"+name)
			resp = append(resp, &fileResponse{FileName: name, URL: "http://localhost:30088/upload/" + name})
		}

		response.Success(c, resp)
	}
}
