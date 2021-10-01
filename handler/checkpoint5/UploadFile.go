package checkpoint5

import (
	"http-theft-bank/handler"
	"http-theft-bank/pkg/errno"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFile ... 上传全排列.go 文件，线上 ac
func UploadFile(c *gin.Context) {
	// log

	// 解析 token

	// get file
	file, err := c.FormFile("file")
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 根据时间戳生成新的文件名
	fileName, fileNameOnly := newFileName(file.Filename)

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "./file/"+fileName)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, "")
		return
	}

	// ac
	err = testProgramme(fileName, fileNameOnly)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// response
	handler.SendResponse(c, errno.OK, nil)
	return
}

func newFileName(fileName string) (string, string) {
	// 根据时间戳生成新的文件名
	extString := path.Ext(fileName)                           // 获取后缀
	fileNameTimeInt := time.Now().Unix()                      // 获取时间戳
	fileNameTimeStr := strconv.FormatInt(fileNameTimeInt, 10) // 时间戳格式化
	filenameOnly := strings.TrimSuffix(fileName, extString)   // 去掉原文件名后缀

	return filenameOnly + fileNameTimeStr + extString, filenameOnly + fileNameTimeStr
}
