package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	apiV2 "gin_demo/api/v2"
)

// @title RESTful API-v2
// @version 2.0
// @license.name Apache 2.0
// @host 127.0.0.1:9091
// @BasePath :9091/api/v2

func Router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery(), apiV2.StatCost())

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		c.InstanceName = "v2"
	}))

	router02 := e.Group("/api/v2")
	{
		router02.POST("/upload", apiV2.SingleUpload)
		router02.POST("/multiUpload", apiV2.MultiUpload)
		router02.GET("/testRedirect", apiV2.RedirectRequest)
		router02.GET("/returnJson", apiV2.ReturnJson)
	}
	return e
}
