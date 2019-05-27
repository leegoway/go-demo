package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)

func Formatter() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		fmt.Println(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		fmt.Println(status)
	}
}