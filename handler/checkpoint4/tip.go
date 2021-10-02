package checkpoint4

import (
	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
)

func BackTips(c *gin.Context) {

	handler.SendResponse(c, errno.OK, handler.Response{
		Code:    errno.OK.Code,
		Message: text.Text4Scene,
		Data:    nil,
	})
}
