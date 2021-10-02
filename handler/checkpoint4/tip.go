package checkpoint4

import (
	"http-theft-bank/handler"
	"http-theft-bank/log"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
	"http-theft-bank/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// BackTips
// @Summary 站点4获取文本
// @Description 站点4，获取游戏场景文本
// @Tags bank
// @Accept  application/json
// @Produce  application/json
// @Param code header string true "代号名"
// @Param passport header string true "通行证"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /bank/iris_recognition_gate [get]
func BackTips(c *gin.Context) {
	log.Info("Message BackTips function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text: text.Text4scene,
	})
}
