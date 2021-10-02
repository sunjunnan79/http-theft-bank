package checkpoint4

import (
	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
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

	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text:      text.Text4scene,
		ExtraInfo: string(text.ImageBytes),
	})

}
