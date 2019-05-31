package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leegoway/go-demo/pkg/app"
	"github.com/leegoway/go-demo/services"
)

func UserRegisterHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form services.RegisterUserForm
	)
	err := appG.Bind2Form(c, &form)
	if err != nil {
		appG.Response(nil, err)
		return
	}
	userService := new(services.UserService)
	u, err := userService.NewUser(form)
	if err != nil {
		appG.Response(u, err)
		return
	}
	appG.Response(u, nil)
}

func UserQueryHandler(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form services.QueryUserForm
	)
	err := appG.Bind2Form(c, &form)
	if err != nil {
		appG.Response(nil, err)
		return
	}
	uservice := new(services.UserService)
	userModel, err := uservice.QueryUser(form)
	if err != nil {
		appG.Response(nil, err)
	} else {
		appG.Response(userModel, nil)
	}
}

func UserUpdateHandler(c *gin.Context) {

}



func UserMultiQueryHandler(c *gin.Context) {

}