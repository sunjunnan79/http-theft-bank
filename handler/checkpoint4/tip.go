package checkpoint4

import (
	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	//"http-theft-bank/pkg/errno"
	"net/http"
)

func BackTips(c *gin.Context) {
	var response handler.Response
	response.Code = http.StatusOK
	var message string = "恭喜这位黑客，你成功来到了第四关卡！现在你的面前是一扇需要虹膜识别的门，你所需要做的是拿到一份虹膜样本。然后你才能拿着这份虹膜样本成功进入到下一关。"
	response.Message = message
	handler.SendResponse(c, nil, response)
}
