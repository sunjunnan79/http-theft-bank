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
			handler.SendBadRequest(c,
				errno.ErrTokenInvalid,
				nil,
				"your request header doesn't have passport")
			c.Abort()
			return
		}
		c.Set("code", ctx.Code)

		c.Next()
	}
}
