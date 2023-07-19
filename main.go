package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Message struct {
	Name    string `json:"user" xml:"user"`
	Message string `json:"message" xml:"message"`
	Status  int    `json:"status" xml:"status"`
	// - 表示不会被序列化
	Password string `json:"-" xml:"-"`
}

// 自定义解析json ---- 极简版
func bindJson(c *gin.Context, obj any) (err error) {
	body, err := c.GetRawData()
	if err != nil {
		return err
	}
	if c.GetHeader("Content-Type") != "application/json" {
		return err
	}
	err = sonic.Unmarshal(body, &obj)
	if err != nil {
		return err
	}
	return nil
}

// 获取查询参数
func query(context *gin.Context) {
	value := context.Query("name")
	if len(value) == 0 {
		fmt.Printf("name is empty\n")
		return
	}
	// ok 表示是否有值
	value, ok := context.GetQuery("name")
	if !ok {
		fmt.Printf("name is empty\n")
		return
	}
	fmt.Printf("name:%s\n", value)
	array := context.QueryArray("name")
	fmt.Println(array)
	dicts := context.QueryMap("name")
	fmt.Println(dicts)

}

// 获取动态参数
func param(context *gin.Context) {
	fmt.Println(context.Param("userid"))
	fmt.Println(context.Param("bookid"))
}

// 获取表单参数
func _postForm(context *gin.Context) {
	fmt.Println(context.PostForm("name"))
	fmt.Println(context.PostFormArray("name"))
	fmt.Println(context.DefaultPostForm("addr", "beijing"))

	fmt.Println(context.MultipartForm()) // 接受所有的表单数据 包括文件
}

// 原始参数
func raw(context *gin.Context) {
	var message Message
	err := bindJson(context, &message)
	if err != nil {
		return
	}
	fmt.Println(message)
}
func addPathQuery(router *gin.Engine) {
	router.GET("/query", query)
	router.GET("/param/:userid", param)
	router.GET("/param/:userid/:bookid", param)
	router.POST("/form", _postForm)
	router.POST("/raw", raw)

}
func main() {
	engine := gin.Default()
	load(engine)
	addPathReturn(engine)
	addPathQuery(engine)

	engine.Run(":8080")
}

// 重定向
// 301和302都是HTTP状态码，表示重定向。
//
// 301状态码表示永久性重定向，即请求的资源已经被永久性地移动到了新的位置，搜索引擎会把新的位置作为该资源的唯一有效地址，
// 因此对于搜索引擎来说，301重定向是最为友好的。
//
// 302状态码表示临时性重定向，即请求的资源暂时被移动到了新的位置，但是该资源的原始地址仍然有效，
// 搜索引擎会继续抓取原始地址，而不是新的地址。
func redirect(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}

// 加载资源和模板
func load(engine *gin.Engine) {
	// 加载html模板
	engine.LoadHTMLGlob("templates/*")
	// 在golang中 没有相对文件的路径概念 只有相对项目路径概念
	engine.StaticFile("/file", "static/123.pdf")
	// 加载静态文件
	//http://127.0.0.1:8080/static/pic3.jpg
	engine.StaticFS("/static", http.Dir("static/pic"))
}

// 配置路径及handler
func addPathReturn(router *gin.Engine) {

	router.GET("/hello", sayHello)
	router.GET("/json", jsonRe)
	router.GET("/json2", jsonRe2)
	router.GET("/xml", xmlRe)
	router.GET("/map", mapRe)
	router.GET("/yml", ymlRe)
	router.GET("/html", htmlRe)
	router.GET("/baidu", redirect)
}

// 响应html
// <!DOCTYPE html>
// <html lang="en">
//
// <head>
//
//	<meta charset="UTF-8">
//	<title>Title</title>
//
// </head>
//
// <body>
//
//	<header>你好,zhangsan</header>
//
//	hello world
//
// </body>
//
// </html>
func htmlRe(context *gin.Context) {
	msg := Message{
		Name:    "zhangsan",
		Message: "hello world",
		Status:  http.StatusOK,
	}
	context.HTML(http.StatusOK, "index.html", gin.H{
		"name": msg.Name,
		"msg":  msg.Message,
	})
}

// 拼接map响应 但是不推荐
func mapRe(context *gin.Context) {
	msg := make(map[string]string, 16)
	msg["name"] = "zhangsan"
	msg["message"] = "hello world"
	msg["status"] = strconv.Itoa(http.StatusOK)

	context.JSON(http.StatusOK, msg)

}

// 拼接xml响应
func xmlRe(context *gin.Context) {
	msg := Message{
		Name:    "zhangsan",
		Message: "hello world",
		Status:  http.StatusOK,
	}
	context.XML(http.StatusOK, msg)
}

func sayHello(context *gin.Context) {
	// 响应字符串
	context.String(http.StatusOK, "hello world")
}

func jsonRe(context *gin.Context) {
	// 响应json
	context.JSON(http.StatusOK, gin.H{
		"message": "hello world",
		"status":  http.StatusOK,
	})
}

func jsonRe2(context *gin.Context) {

	msg := Message{
		Name:    "zhangsan",
		Message: "hello world",
		Status:  http.StatusOK,
	}
	// 响应json
	context.JSON(http.StatusOK, msg)
}

// 拼接yml响应
func ymlRe(context *gin.Context) {
	msg := Message{
		Name:    "zhangsan",
		Message: "hello world",
		Status:  http.StatusOK,
	}
	context.YAML(http.StatusOK, msg)
}
