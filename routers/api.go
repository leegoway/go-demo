package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/leegoway/go-demo/middlewares"
	"github.com/leegoway/go-demo/controller"
	)

func InitRouters() *gin.Engine  {
	r := gin.Default()

	v1 := r.Group("/v1")
	v1.Use(middlewares.Formatter())
	{
		v1.POST("/user/register", controller.UserRegisterHandler)
		v1.POST("/user/update", controller.UserUpdateHandler)
		v1.GET("/user/query", controller.UserQueryHandler)
		v1.POST("/user/multi_query", controller.UserMultiQueryHandler)
	}
	return r
}
