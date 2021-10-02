package checkpoint3

import (
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type CorrectRequest struct {
	ErrorCode string `json:"error_code"`
}

func GetMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrGet, nil, "error")
}

func PostMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrPost, nil, "error")

}

func PutMethod(c *gin.Context) {
	var data CorrectRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	errorCode, err := encrypt.AESDecodeAfterBase64([]byte(data.ErrorCode), []byte(viper.GetString("sercet_key")))
	if err != nil {
		handler.SendBadRequest(c, errno.ErrDecode, nil, err.Error())
		return
	}

	if string(errorCode) != viper.GetString("error_code") {
		handler.SendBadRequest(c, errno.ErrMatch, nil, err.Error())
	}

	handler.SendResponse(c, errno.OK, nil)
}

func DelMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrDel, nil, "error")
}

func PatchMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrPatch, nil, "error")
}
