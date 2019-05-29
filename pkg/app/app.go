package app

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/leegoway/go-demo/pkg/e"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

// BindAndValid binds and validates data
func (app Gin)Bind2Form(c *gin.Context, form interface{}) *e.ErrorCode {
	err := c.Bind(form)
	fmt.Println("parsing form", form)
	if err != nil {
		return e.NewCoder(e.INVALID_PARAMS, "")
	}
	return nil
}

func Valid(form interface{}) uint32 {
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

type Response struct {
	Code uint32         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(errCode uint32, data interface{}, msg string) {
	if msg == ""{
		msg = e.GetMsg(errCode)
	}
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}
