package read

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Folder map[string]string `yaml:"folder"`
}

func ReadText() map[string]string {
	var c conf
	txtMap := make(map[string]string, 10)

	con := c.getConf()
	fmt.Println(con)

	// 遍历map中的路径
	for key, value := range con.Folder {
		txt, _ := ioutil.ReadFile(value)

		content := string(txt)
		txtMap[key] = content
		// fmt.Println(content)
	}

	return txtMap
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("/Users/zhangkuang/Desktop/bank/http-theft-bank/conf/txt.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}
