package checkpoint4

import (
	//"log"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func VerifyParameter(c *gin.Context) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename

	// 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()

	// 将file的内容拷贝到out
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "upload successful \n")

	// 开始检验上传的文件是否合法
	ext := GetExt(filename)
	// 得到没有后缀的文件名
	filename = strings.TrimSuffix(filename, ext)
	// 判断文件名是否合法
	tag := strings.ContainsAny(filename, "MuXieye")
	if tag {
		var response handler.Response
		response.Code = http.StatusOK
		var message string = "恭喜这位黑客，你成功通过了第四关卡！"
		response.Message = message
		handler.SendResponse(c, nil, response)
	} else {
		var response handler.Response
		response.Code = errno.ErrValidation.Code
		err := errno.ErrValidation.Error()
		var message string = "虹膜识别失败！未通过第四关卡！"
		response.Message = message
		handler.SendError(c, errors.New("验证失败！"), response, err)
	}
}

// GetExt：获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}
