package api

import "github.com/gin-gonic/gin"

func AddRouterGroup(router *gin.Engine) {
	router.Group("/users")
}
