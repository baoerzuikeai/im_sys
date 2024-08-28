package service

import (
	"net/http"
	"strconv"

	"github.com/baoer/im_sys/models"
	"github.com/gin-gonic/gin"
)

func PrivateChatlist() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// ?roomidentity & pageIdenx & pageSize
func ChannelChatlist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roomidentity := ctx.Query("roomidentity")
		if roomidentity == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "房间号不能为空",
			})
		}
		pageIdenx, _ := strconv.ParseInt(ctx.Query("pageIdenx"), 10, 32)
		pageSize, _ := strconv.ParseInt(ctx.Query("pageSize"), 10, 32)
		skip := (pageIdenx - 1) * pageSize
		data, err := models.GetPublicMsgbyRooMidentity(roomidentity, &pageSize, &skip)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "系统异常" + err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "数据加载成功",
			"data": gin.H{
				"list": data,
			},
		})
	}
}
