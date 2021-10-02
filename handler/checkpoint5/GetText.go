package checkpoint5

import (
	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
)

// GetText ... 获取站点5文本
// @Summary 站点5 获取游戏场景文本
// @Description 站点5，获取文本
// @Tags end
// @Accept  application/json
// @Produce  application/json
// @Param code header string true "代号名"
// @Param passport header string true "通行证"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /muxi/backend/computer/examination [get]
// TODO：解析 token 然后把 代号名代入 Text5Scene
func GetText(c *gin.Context) {
	// log

	// 解析 token

	handler.SendResponse(c, errno.OK, handler.TextInfo{
		Text: text.Text5Scene,
	})
}
