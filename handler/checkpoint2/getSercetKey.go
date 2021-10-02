package checkpoint2

import (
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetSecretKey(c *gin.Context) {
	content := viper.GetString("sercet_key") + " : " + viper.GetString("error_code")
	secretKey := encrypt.Base64Encode([]byte(content))
	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text:      text.Text2Success,
		ExtraInfo: secretKey,
	})
}
