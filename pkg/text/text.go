package text

import "github.com/spf13/viper"

var Text5Scene string

func InitText() {
	// TODO: get text in another file
	Text5Scene = viper.GetString("scene")
}
