package checkpoint3

import (
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GetMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrGet, nil, "error")
}

func PostMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrPost, nil, "error")

}

func PutMethod(c *gin.Context) {
	handler.SendResponse(c, errno.OK, nil)
}

func DelMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrDel, nil, "error")
}

func PatchMethod(c *gin.Context) {
	handler.SendBadRequest(c, errno.ErrPatch, nil, "error")
}
