package routers

import (
	"github.com/gin-gonic/gin"
	"demo/middlewares"
	"demo/handlers"
	)

func InitRouters(r *gin.Engine)  {
	v1 := r.Group("/v1")
	v1.Use(middlewares.Formatter())
	{
		v1.POST("/user/register", handlers.UserRegisterHandler)
		v1.POST("/user/update", handlers.UserUpdateHandler)
		v1.POST("/user/query", handlers.UserQueryHandler)
		v1.POST("/user/multi_query", handlers.UserMultiQueryHandler)
	}

}
