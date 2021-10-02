package checkpoint4

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	//"http-theft-bank/pkg/errno"
	"io/ioutil"
	"net/http"
)

func UserGetImage(c *gin.Context) {
	bytes, err := ioutil.ReadFile("./file/testSample/MuXieye.jpg")
	if err != nil {
		fmt.Printf("fail to read MuXieye.jpg:%v\n", err)
	}
	var gBytes []byte = bytes

	c.Data(http.StatusOK, "image", gBytes)
	var response handler.Response
	response.Code = http.StatusOK
	var message string = "恭喜这位黑客你成功得到了你的虹膜样本，你所需要需要做的就是下载你的虹膜样本。"
	response.Message = message
	handler.SendResponse(c, nil, response)

}
