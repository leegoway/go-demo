package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/leegoway/go-demo/pkg/app"
	"github.com/leegoway/go-demo/pkg/e"
	"github.com/leegoway/go-demo/services"
)

func UserRegisterHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form services.RegisterUserForm
	)
	err := appG.Bind2Form(c, &form)
	if err != nil {
		appG.Response(err.Code(), nil, err.Msg())
		return
	}
	userService := new(services.UserService)
	u, err1 := userService.NewUser(form)
	if err1 != nil {
		appG.Response(e.ERROR, nil, err1.Error())
		return
	}
	appG.Response(e.SUCCESS, u, "")
}

func UserQueryHandler(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form services.QueryUserForm
	)
	err := appG.Bind2Form(c, &form)
	if err != nil {
		appG.Response(err.Code(), nil, err.Msg())
		return
	}
	uservice := new(services.UserService)
	userModel, err := uservice.QueryUser(form)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			appG.Response(e.ERROR_USER_NOT_EXIST, nil, "")
		} else {
			appG.Response(e.ERROR, nil, err.Error())
		}
	} else {
		appG.Response(e.SUCCESS, userModel, "")
	}
}

func UserUpdateHandler(c *gin.Context) {

}



func UserMultiQueryHandler(c *gin.Context) {

}