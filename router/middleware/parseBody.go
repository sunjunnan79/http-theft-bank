package middleware

import (
	"bytes"
	"encoding/json"
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"io/ioutil"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/gin-gonic/gin"
)

func paresBodyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		// check whether the data abide by the role of json structure
		newBody, err := encrypt.Base64Decode(string(data))
		if err != nil {
			handler.SendResponse(c, errno.ErrDecode, nil)
			c.Abort()
			return
		}

		if !json.Valid(data) {
			handler.SendResponse(c, errno.ErrBodyInvalid, nil)
			c.Abort()
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(newBody)))
		c.Next()
	}
}
