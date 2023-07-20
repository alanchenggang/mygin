package util

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// CustomValidator 自定义验证器注册
func CustomValidator() {
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("sign", SignValidator)
	}

}

func SignValidator(fl validator.FieldLevel) bool {
	var nameList = []string{
		"root",
		"admin",
		"sysmaster",
		"sysadmin",
	}
	flStr := fl.Field().String()
	for _, s := range nameList {
		if s == flStr {
			return false
		}
	}
	return true
}

// CustomizeFailMessage 自定义错误信息处理
func CustomizeFailMessage(err error, obj any) (errMsg string) {
	getObjType := reflect.TypeOf(obj)

	// 1. 判断 err 是否是 validator.ValidationErrors 类型
	if errors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range errors {
			// 循环每一个错误信息
			// 通过反射获取到 tag 中的 json tag
			if filed, exist := getObjType.Elem().FieldByName(fieldError.Field()); exist {
				// 2. 判断是否有 json tag
				if jsonTag := filed.Tag.Get("json"); len(jsonTag) > 0 {
					// 3. 判断是否有自定义错误信息
					if customTag := filed.Tag.Get("msg"); len(customTag) > 0 {
						// 4. 返回自定义错误信息
						errMsg = customTag
						return
					}
				}
			}
		}
	}
	return err.Error()
}
