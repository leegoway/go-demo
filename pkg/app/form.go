package app

import (
	"github.com/gin-gonic/gin"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
)

// BindAndValid binds and validates data
func Bind2Form(c *gin.Context, form interface{}) int {
	err := c.Bind(form)
	if err != nil {
		return e.INVALID_PARAMS
	}
	return e.SUCCESS
}
