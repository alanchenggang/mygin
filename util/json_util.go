package util

import (
	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
)

// 自定义解析json ---- 极简版
func BindJson(c *gin.Context, obj any) (err error) {
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
