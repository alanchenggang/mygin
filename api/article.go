package api

import (
	"MyGin/core"
	"MyGin/model"
	"MyGin/server"
	"MyGin/util"
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
