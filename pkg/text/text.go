package text

import (
	"io/ioutil"
)

const dir = "./text/"

var (
	Text1Success string

	Text2Success string

	Text3Success string

	Text4scene string
	Text4Success string

	Text5Scene string
	Text5Success string
)

// InitText ... 程序启动时读取文件
func InitText() {
	Text1Success = readFile(dir + "checkpoint1/success.txt")

	Text2Success = readFile(dir+"checkpoint2/success.txt")

	Text3Success = readFile(dir + "checkpoint3/success.txt")

	Text4Success = readFile(dir+"checkpoint4/success.txt")
	Text4scene = readFile(dir+"checkpoint4/scene.txt")

	Text5Scene = readFile(dir + "checkpoint5/scene.txt")
	Text5Success = readFile(dir + "checkpoint5/success.txt")
}

func readFile(file string) string {
	txt, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(txt)
}
