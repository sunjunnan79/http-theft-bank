package checkpoint5

import (
	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
)

// GetText ... 获取站点6文本
// TODO：解析 token 然后把 代号名代入 Text5Scene
func GetText(c *gin.Context) {
	// log

	// 解析 token

	handler.SendResponse(c, errno.OK, text.Text5Scene)
}
