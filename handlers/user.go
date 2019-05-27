package handlers

import (
	"github.com/gin-gonic/gin"
	"demo/services"
	"demo/pkg/e"
	"demo/pkg/app"
	"fmt"
)

func UserRegisterHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form services.RegisterUserForm
	)
	errCode := app.Bind2Form(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}
	userService := new(services.UserService)
	u, _ := userService.NewUser(form)
	appG.Response(e.SUCCESS, u)
}

func UserUpdateHandler(c *gin.Context) {

}

func UserQueryHandler(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form services.QueryUserForm
	)
	errCode := app.Bind2Form(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil)
		return
	}
	fmt.Println(form)
	uservice := new(services.UserService)
	userModel, _ := uservice.QueryUser(form)
	c.JSON(200, gin.H{"code":200, "msg":"", "data": userModel})
}

func UserMultiQueryHandler(c *gin.Context) {

}