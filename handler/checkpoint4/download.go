package checkpoint4

import (
	"encoding/base64"
	"http-theft-bank/handler"
	"http-theft-bank/log"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
	"http-theft-bank/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserGetImage
// @Summary 下载图片
// @Description 站点4，下载图片
// @Tags organization
// @Accept  application/json
// @Produce  application/json
// @Param code header string true "代号名"
// @Param passport header string true "通行证"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /organization/iris_sample [get]
func UserGetImage(c *gin.Context) {
	log.Info("Message UserGetImage function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text:      text.Text4scene,
		ExtraInfo: base64.StdEncoding.EncodeToString(text.ImageBytes),
	})

}
