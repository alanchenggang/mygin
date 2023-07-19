package api

import (
	"MyGin/core"
	"MyGin/model"
	"MyGin/server"
	"MyGin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetList(c *gin.Context) {
	resp := server.GetListService()
	c.JSON(http.StatusOK, resp)
}

func GetInfo(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "id is empty"))
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "id is wrong"))

	}
	resp := server.GetDetailService(idInt)
	c.JSON(http.StatusOK, resp)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "id is empty"))
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "id is wrong"))
	}
	var article model.ArticleModel

	util.BindJson(c, &article)

	resp := server.ModifyDetailService(idInt, article)
	c.JSON(http.StatusOK, resp)
}

func Create(c *gin.Context) {
	var article model.ArticleModel

	err := util.BindJson(c, &article)
	if err != nil {
		return
	}
	resp := server.CreateService(article)
	c.JSON(http.StatusOK, resp)
}
func Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "id is empty"))
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "id is wrong"))
	}
	resp := server.DeleteService(idInt)
	c.JSON(http.StatusOK, resp)
}

// Upload 文件上传 单个文件
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "file is empty"))
	}
	filename := file.Filename
	err = c.SaveUploadedFile(file, "static/upload/"+filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "file save error"))
	}
	c.JSON(http.StatusOK, core.Success(filename))
}

// Uploads 文件上传 多个文件
func Uploads(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "file is empty"))
	}
	files := form.File["upload[]"]
	var filenames []string
	for _, file := range files {
		filename := file.Filename
		size := file.Size / 1024
		fmt.Println("filename ", filename, " , size ", size)
		err = c.SaveUploadedFile(file, "static/upload/"+filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, "file save error"))
		}
		filenames = append(filenames, filename)
	}
	c.JSON(http.StatusOK, core.Success(filenames))
}

// 文件下载
func Download(c *gin.Context) {
	c.File("static/pic/pic1.jpg")
	// 标识是文件流 强制浏览器进行下载行为
	c.Header("Content-Type", "application/octet-stream")
	// 指定下载后文件名
	c.Header("Content-Disposition", "attachment; filename=pic1.jpg")
	// 标识传输过程中编码形式 避免乱码问题
	c.Header("Content-Transfer-Encoding", "binary")

	// 前后端分离的情况下，一般将msg和filename塞入header中
}
