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
	auth.GET("/chat/sendmessage/private", service.WebSocketsendPrivateMessage())
	auth.GET("/chat/sendmessage/channel", service.WebSocketsendChannelMessage())
	//用户详情
	auth.GET("/user/detail", service.UserDetail())
	//聊天列表
	auth.GET("/getlist/private", service.PrivateChatlist())
	auth.GET("/getlist/channel", service.ChannelChatlist())
	return router
}
