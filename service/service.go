package service

import (
	"log"
	"net/http"
	"net/smtp"

	"github.com/baoer/im_sys/models"
	"github.com/baoer/im_sys/template"
	"github.com/baoer/im_sys/util"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		account := ctx.PostForm("account")
		password := ctx.PostForm("password")
		if account == "" || password == "" {
			ctx.JSON(200, gin.H{
				"code": -1,
				"msg":  "账号或密码不能为空",
			})
			return
		}
		ub, err := models.GetUserBasicBy_AccountPassword(account, util.Getmd5(password))
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": -1,
				"msg":  "账号或密码不正确",
			})
			log.Println(err)
			return
		}
		token, err := util.Gettoken(ub.Identity, ub.Email)
		if err != nil {
			ctx.JSON(200, gin.H{
				"code": -1,
				"msg":  "System error:" + err.Error(),
			})
			return
		}
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "登录成功",
			"data": gin.H{
				"token": token,
			},
		})
	}
}

func UserDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userclaims, _ := ctx.Get("user_claims")
		uc := userclaims.(*util.Myclaims)
		userbasic, err := models.GetUserBasicBy_Identity(uc.UserId)
		if err != nil {
			log.Printf("[DB ERROR]:%v\n", err)
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "数据查询异常",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "数据加载成功",
			"data": userbasic,
		})
	}
}

// 授权码pdswwunzbwnubgih
func SenCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uemail := ctx.PostForm("email")
		if uemail == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "邮箱不能为空",
			})
			return
		}
		count, err := models.GetUserBasicCountBy_Email(uemail)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "查询出错" + err.Error(),
			})
			return
		}
		if count > 0 {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "邮箱已经被注册",
			})
			return
		}
		e := email.NewEmail()
		e.From = "BaoEr <485191245@qq.com>"
		e.To = []string{uemail}
		e.Subject = "Verification Code"
		e.Text = []byte("Text Body is, of course, supported!")

		e.HTML = template.ParseSendCodetemplate()
		err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "485191245@qq.com", "pdswwunzbwnubgih", "smtp.qq.com"))
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "发送邮件失败" + err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "验证码发送成功",
		})
	}
}
