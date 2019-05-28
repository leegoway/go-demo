package handlers

import (
	"github.com/gin-gonic/gin"
	"demo/services"
	"demo/pkg/e"
	"demo/pkg/app"
	"github.com/jinzhu/gorm"
)

func UserRegisterHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form services.RegisterUserForm
	)
	errCode := app.Bind2Form(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil, "")
		return
	}
	userService := new(services.UserService)
	u, err := userService.NewUser(form)
	if err != nil {
		appG.Response(e.ERROR, nil, err.Error())
		return
	}
	appG.Response(e.SUCCESS, u, "")
}

func UserQueryHandler(c *gin.Context)  {
	var (
		appG = app.Gin{C: c}
		form services.QueryUserForm
	)
	errCode := app.Bind2Form(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(errCode, nil, "")
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