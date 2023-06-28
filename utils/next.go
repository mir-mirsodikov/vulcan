package utils

import (
	"github.com/spf13/viper"
)

func GetFileExtension() string {
	useJsx := viper.GetBool("next.jsx")
	if useJsx {
		return ".jsx"
	}
	return ".tsx"
}
