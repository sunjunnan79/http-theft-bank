package text

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const dir = "./text/"

var (
	Text1Success string

	Text2Success string

	Text3Success string

	Text4scene   string
	Text4Success string

	Text5Scene   string
	Text5Success string

	ImageBytes []byte

	// cp5: 全排列答案(3组数据)
	Answers [3][]string
)

// InitText ... 程序启动时读取文件
func InitText() {
	Text1Success = readFile(dir + "checkpoint1/success.txt")

	Text2Success = readFile(dir + "checkpoint2/success.txt")

	Text3Success = readFile(dir + "checkpoint3/success.txt")

	Text4Success = readFile(dir + "checkpoint4/success.txt")
	Text4scene = readFile(dir + "checkpoint4/scene.txt")

	Text5Scene = readFile(dir + "checkpoint5/scene.txt")
	Text5Success = readFile(dir + "checkpoint5/success.txt")

	// 获取答案并处理为[]string
	for i := 0; i < 3; i++ {
		fimename := "./file/testSample/test" + strconv.Itoa(i+1) + "_result.txt"
		file, err := os.Open(fimename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		content, _ := ioutil.ReadAll(file)

		Answers[i] = strings.Split(string(content), string(rune(10)))
	}

	ImageBytes = []byte(readFile("./text/MuXieye.jpg"))
}

func readFile(file string) string {
	txt, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(txt)
}
