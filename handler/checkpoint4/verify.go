package checkpoint4

import (
	"http-theft-bank/handler"
	"http-theft-bank/log"
	"http-theft-bank/pkg/constvar"
	"http-theft-bank/pkg/errno"
	"http-theft-bank/pkg/text"
	"http-theft-bank/util"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// VerifyParameter
// @Summary 站点4 上传图片
// @Description 站点4，上传图片
// @Tags bank
// @Accept  application/json
// @Produce  application/json
// @Param file body string true "file,这个是用表单上传！！！！"
// @Param code header string true "代号名"
// @Param passport header string true "通行证"
// @Success 200 {object} handler.Response
// @Failure 401 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /organization/iris_recognition_gate [post]
func VerifyParameter(c *gin.Context) {
	log.Info("Message VerifyParameter function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	header, err := c.FormFile("file")
	if err != nil {
		handler.SendBadRequest(c, errno.ErrFormFile, "", "表单字段错误或缺失，请统一改为file")
		return
	}
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename

	// 开始检验上传的文件是否合法
	ext := GetExt(filename)

	// 得到没有后缀的文件名
	filename = strings.TrimSuffix(filename, ext)

	// 判断文件名是否合法
	tag := strings.ContainsAny(filename, "MuXieye")
	if tag {
		handler.SendResponse(c, errno.OK, handler.TextInfo{
			Text: text.Text4Success,
		})
		return
	}

	handler.SetResponseHeader(c, constvar.FragmentField, constvar.Fragment4)
	handler.SendBadRequest(c, errno.ErrPicture, nil, "")
}

// GetExt ... 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}
