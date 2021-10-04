package checkpoint1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"http-theft-bank/handler"
	"http-theft-bank/log"
	"http-theft-bank/pkg/constvar"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
	"http-theft-bank/pkg/token"
	"http-theft-bank/util"
	"net/http"
)

// CheckCode ... header的 code 生成token
// @Summary send user token
// @Description 站点1，返回token在头部
// @Tags organization
// @Accept  application/json
// @Produce  application/json
// @Param code header string true "代号名"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /organization/code [get]
func CheckCode(c *gin.Context) {
	log.Info("Message CheckCode function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	code := c.Request.Header.Get("code")
	if code == "" {
		//err := errors.New("未接收到code字段")
		//handler.SendBadRequest(c, err, "", "请在request头添加 code 字段，值为你的组织代号名")

		// 仍然返回文本，但是不给 token
		c.HTML(http.StatusOK, "start.html", nil)
		return
	}

	Token, err := token.Sign(c, token.Context{Code: code}, "")

	if err != nil {
		log.Info("sign token wrong")
		handler.SendError(c, err, "try again", "sign token wrong")
		return
	}

	handler.SetResponseHeader(c, "passport", Token)
	handler.SetResponseHeader(c, constvar.FragmentField, constvar.Fragment1)

	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text: text.Text1Success,
	})
}
