package checkpoint3

import (
	"http-theft-bank/pkg/text"
	"strings"

	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrWrongMethod, nil, "")
}

func PostMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrWrongMethod, nil, "")

}

func PutMethod(c *gin.Context) {
	var data handler.Request
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	errorCode, err := encrypt.AESDecodeAfterBase64([]byte(data.Content), []byte(viper.GetString("sercet_key")))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrDecode, nil, err.Error())
		return
	}

	if strings.Compare(string(errorCode), viper.GetString("error_code")) != 0 {
		handler.SendBadRequest(c, errno.ErrMatch, nil, "输入病毒无效")
		return
	}

	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text: text.Text3Success,
	})
}

func DelMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrWrongMethod, nil, "")
}

func PatchMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrWrongMethod, nil, "")
}
