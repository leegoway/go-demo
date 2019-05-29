package app

import (
	"github.com/gin-gonic/gin"
	"github.com/leegoway/go-demo/pkg/e"
	"github.com/astaxie/beego/validation"
	"fmt"
	)

// BindAndValid binds and validates data
func Bind2Form(c *gin.Context, form interface{}) int {
	err := c.Bind(form)
	if err != nil {
		return e.INVALID_PARAMS
	}
	return e.SUCCESS
}

func Valid(form interface{}) int {
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		fmt.Println("valid error: ", err)
		return e.ERROR
	}
	if !check {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
		}
		return e.INVALID_PARAMS
	}
	return e.SUCCESS
}