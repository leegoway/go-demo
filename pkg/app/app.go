package app

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/leegoway/go-demo/pkg/e"
	"net/http"
	"strings"
)

type Gin struct {
	C *gin.Context
}

// BindAndValid binds and validates data
func (app Gin)Bind2Form(c *gin.Context, form interface{}) *e.ErrorCode {
	err := c.Bind(form)
	fmt.Println("parsing form", form)
	if err != nil {
		fmt.Println("binding error", err)
		return e.NewCoder(e.INVALID_PARAMS, err.Error())
	}
	return nil
}

func Valid(form interface{}) *e.ErrorCode {
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		fmt.Println("验证参数错误: ", err)
		return e.Wrap(err)
	}
	if !check {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
			index := strings.Index(err.Key, ".")
			msg := err.Key[:index]
			return e.NewCoder(e.INVALID_PARAMS, fmt.Sprintf("%s %s", msg, err.Message))
		}
	}
	return nil
}

type Response struct {
	Code uint32         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(data interface{}, err *e.ErrorCode) {
	var msg string
	var errCode uint32
	if err != nil{
		errCode = err.Code()
		msg = err.Msg()
		if msg == "" {
			msg = e.GetMsg(errCode)
		}
	}
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}
