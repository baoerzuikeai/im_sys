package middleware

import (
	"github.com/baoer/im_sys/util"
	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Token")
		userclaims, err := util.Parsetoken(token)
		if err != nil {
			ctx.Abort()
			ctx.JSON(200, gin.H{
				"code": -1,
				"msg":  "认证不通过",
			})
			return
		}
		ctx.Set("user_claims", userclaims)
	}
}
