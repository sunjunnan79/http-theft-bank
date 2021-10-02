package checkpoint2

import (
	"http-theft-bank/handler"
	"http-theft-bank/log"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
	"http-theft-bank/util"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// GetSecretKey ... 获取
// @Summary send user secretKey and error function
// @Description 站点2，返回加密的密钥和错误代码片段，在 body
// @Tags organization
// @Accept  application/json
// @Produce  application/json
// @Param code header string true "代号名"
// @Param passport header string true "通行证"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /organization/secret_key [get]
func GetSecretKey(c *gin.Context) {
	log.Info("Message GetSecretKey function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	content := viper.GetString("sercet_key") + " : " + viper.GetString("error_code")
	secretKey := encrypt.Base64Encode([]byte(content))
	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text:      text.Text2Success,
		ExtraInfo: secretKey,
	})
}
