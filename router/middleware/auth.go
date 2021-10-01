package middleware

import (
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}
		c.Set("code", ctx.Code)

		c.Next()
	}
}
