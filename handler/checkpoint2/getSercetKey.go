package checkpoint2

import (
	"http-theft-bank/handler"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type getSercetKeyResponse struct {
	SercetKey string `json:"sercet_key"`
}

func GetSercetKey(c *gin.Context) {
	content := viper.GetString("sercet_key") + " : " + viper.GetString("error_code")
	getSercetKeyRes := getSercetKeyResponse{
		SercetKey: encrypt.Base64Encode([]byte(content)),
	}
	handler.SendResponse(c, nil, getSercetKeyRes)
}
