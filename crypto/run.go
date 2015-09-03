package crypto

import (
	"github.com/spf13/viper"
)

func Run() {
	if viper.GetBool("gen") {
		Gen()
	}
}
