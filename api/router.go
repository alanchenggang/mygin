package api

import (
	"MyGin/core"
	"MyGin/model"
	"MyGin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddRouterGroup(router *gin.Engine) {
	userGroup := router.Group("/users", UserGroupFunc)
	{
		userGroup.GET("/list", UserListFunc)
		userGroup.POST("/", CreateUser)

	}

}

func CreateUser(context *gin.Context) {
	var user model.UserModel
	err := context.ShouldBindJSON(&user)
	if err != nil {
		errMsg := util.CustomizeFailMessage(err, &user)
		context.JSON(http.StatusBadRequest, core.Fail(http.StatusBadRequest, errMsg))
		return
	}
	context.JSON(http.StatusOK, core.Success(user))

}

func UserListFunc(context *gin.Context) {
	userList := []model.UserModel{
		{"jack", 18},
		{"nack", 18},
		{"fuck", 18},
	}
	context.JSON(http.StatusOK, core.Success(userList))
}

func UserGroupFunc(context *gin.Context) {
	fmt.Println("进入了用户组router")
}
