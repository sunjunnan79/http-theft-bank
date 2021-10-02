package checkpoint4

import (
	"github.com/gin-gonic/gin"
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
	"path"
	"strings"
)

func VerifyParameter(c *gin.Context) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	header, err := c.FormFile("image")
	if err != nil {
		handler.SendResponse(c, errno.ErrPost, handler.Response{
			Code:    errno.ErrBind.Code,
			Message: "上传文件失败！", //上传游戏图片失败的的提示
			Data:    nil,
		})

	}
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename

	handler.SendResponse(c, errno.OK, handler.Response{
		Code:    errno.OK.Code,
		Message: "上传虹膜图片成功！",
		Data:    nil,
	})
	// 开始检验上传的文件是否合法
	ext := GetExt(filename)
	// 得到没有后缀的文件名
	filename = strings.TrimSuffix(filename, ext)
	// 判断文件名是否合法
	tag := strings.ContainsAny(filename, "MuXieye")
	if tag {
		handler.SendResponse(c, errno.OK, handler.Response{
			Code:    errno.OK.Code,
			Message: text.Text4Scene, //通过游戏的提示
			Data:    nil,
		})

	} else {

		handler.SendResponse(c, errno.ErrValidation, handler.Response{
			Code:    errno.ErrValidation.Code,
			Message: text.Text4Scene, //未通过游戏的提示
			Data:    nil,
		})
	}
}

// GetExt：获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}
