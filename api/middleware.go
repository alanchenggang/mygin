package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddMiddleWare 多个中间件按照顺序执行
func AddMiddleWare(router *gin.Engine) {
	// 注册全局中间件
	router.Use(m10)
	// 从前到后依次执行中间件 从后往前执行响应中间件
	//m1 ... in
	//m2 ... in
	//m3 ... in
	//m3 ... out
	//m2 ... out
	//m1 ... out
	router.GET("/ping", m1, m2, m3)
}
func m10(context *gin.Context) {
	fmt.Println("m10 ... in")

}

func m3(context *gin.Context) {
	fmt.Println("m3 ... in")
	context.Next() //

	//context.Abort()// 终止执行
	fmt.Println("m3 ... out")

}

func m2(context *gin.Context) {
	fmt.Println("m2 ... in")
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"code":    http.StatusOK,
	})
	context.Next()
	fmt.Println("m2 ... out")

}

func m1(context *gin.Context) {
	fmt.Println("m1 ... in")
	context.Next()
	fmt.Println("m1 ... out")
}
