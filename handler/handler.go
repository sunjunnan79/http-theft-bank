package handler

import (
	"http-theft-bank/log"
	"http-theft-bank/util"
	"net/http"

	"http-theft-bank/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Request struct {
	Content string `json:"content"` // 目前只有 checkpoint3 需要这个请求，都放这
	ExtraInfo string `json:"extra_info"` // 此字段暂时不用
}


type TextInfo struct {
	Text      string `json:"text"`
	ExtraInfo string `json:"extra_info"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SetResponseHeader ... 设置响应头
func SetResponseHeader(c *gin.Context, key, value string) {
	c.Header(key, value)
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	log.Info(message,
		zap.String("X-Request-Id", util.GetReqID(c)))

	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause))

	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}, cause string) {
	code, message := errno.DecodeErr(err)
	log.Error(message,
		zap.String("X-Request-Id", util.GetReqID(c)),
		zap.String("cause", cause))

	c.JSON(http.StatusInternalServerError, Response{
		Code:    code,
		Message: message + ": " + cause,
		Data:    data,
	})
}
