package v2

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// StatCost 统计请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		//可以设置值,后续的处理函数中可以使用
		c.Set("name", "renwen")
		// 调用该请求的剩余处理程序
		c.Next()
		// 如果不想调用该请求剩余的处理程序,可以调用 c.Abort()
		cost := time.Since(start)
		log.Println(cost)
	}
}
