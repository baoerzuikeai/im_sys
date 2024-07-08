package router

import (
	"github.com/baoer/im_sys/middleware"
	"github.com/baoer/im_sys/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.POST("/login", service.Login())
	router.POST("/send/code", service.SenCode())
	router.POST("/register", func(ctx *gin.Context) {

	})

	auth := router.Group("/u", middleware.AuthCheck())

	auth.GET("/user/detail", service.UserDetail())

	return router
}
