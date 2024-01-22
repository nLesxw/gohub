//Package requests 处理请求数据和表单验证
package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool{

	//1. 解析请求， 支持 JSON 数据，表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 Multipart 标头，参数使用 JSON 格式",
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}

	//2. 表单验证
	errs := handler(obj, c)

	//3. 判断验证是否通过
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string{
	//配置初始化
	opts := govalidator.Options{
		Data: data,
		Rules: rules,
		Messages: messages,
		TagIdentifier: "valid", //模型中 struct 标签标识符
	}

	//开始验证
	return govalidator.New(opts).ValidateStruct()
	
}