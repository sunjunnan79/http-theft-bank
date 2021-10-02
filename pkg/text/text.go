package text

import (
	"io/ioutil"
)

const dir = "./text/"

var Text1Scene string
var Text5Scene string

func InitText() {
	// TODO: get text in another file
	Text1Scene = readFile(dir + "scene1.txt")
	Text5Scene = readFile(dir + "scene5.txt")
}

func readFile(file string) string {
	txt, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(txt)
}
