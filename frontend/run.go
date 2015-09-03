package frontend

import (
	"net/http"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/telem/web"
)

func Run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Form)
	mux.HandleFunc("/upload", Upload)
	web.Start(
		mux,
		viper.GetString("address"),
		viper.GetString("port"),
		viper.GetString("cert"),
		viper.GetString("key"),
	)
}
