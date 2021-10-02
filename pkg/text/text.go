package text

import "github.com/spf13/viper"

var Text1Scene string
var Text5Scene string
var Text4Scene string

func InitText() {
	// TODO: get text in another file
	Text1Scene = viper.GetString("scene")
	Text4Scene = viper.GetString("scene")
	Text5Scene = viper.GetString("scene")

}
