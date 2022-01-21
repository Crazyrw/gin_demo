package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	apiV1 "gin_demo/api/v1"
)

// @title RESTful API-v1
// @version 1.0
// @license.name Apache 2.0
// @host 127.0.0.1:9090
// @BasePath :9090/api/v1

func Router01() http.Handler {
	e := gin.New()
	//加入两个中间件
	e.Use(gin.Recovery(), apiV1.StatCost())

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		c.InstanceName = "v1"
	}))

	// 路由组
	router01 := e.Group("/api/v1")
	{
		router01.GET("/hello", apiV1.HandleHello)
		router01.POST("/formData", apiV1.HandleForm)
		router01.POST("/json", apiV1.GetJson)
	}
	return e
}
