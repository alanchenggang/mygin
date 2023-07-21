package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func init() {
	file, err := os.Create("D:\\CodeProjects\\logs\\gin\\gin.log")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
