package checkpoint1

import (
	"http-theft-bank/handler"
	"http-theft-bank/log"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
	"http-theft-bank/pkg/token"

	"github.com/gin-gonic/gin"
)

// CheckCode ... header的 code 生成token
func CheckCode(c *gin.Context) {
	code := c.Request.Header.Get("code")
	if err := c.BindJSON(&code); err != nil || code == "" {
		handler.SendError(c, err, "", "请在request头添加 code 字段，值为你的组织代号名")
		return
	}

	Token, err := token.Sign(c, token.Context{Code: code}, "")

	if err != nil {
		log.Info("sign token wrong")
		handler.SendError(c, err, "try again", "sign token wrong")
		return
	}

	handler.SetResponseHeader(c, "token", Token)

	handler.SendResponse(c, errno.OK, text.Text1Scene)
}
