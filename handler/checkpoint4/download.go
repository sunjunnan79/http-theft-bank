package checkpoint4

import (
	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
)

func UserGetImage(c *gin.Context) {

	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text:      text.Text4Scene,
		ExtraInfo: string(handler.ImageBytes),
	})

}
