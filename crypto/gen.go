package crypto

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"

	"github.com/pdxjohnny/telem/log"
)

func Gen() {
	pair, err := NewPair(
		viper.GetString("name"),
		viper.GetString("comment"),
		viper.GetString("email"),
	)
	if log.PrintError("crypto", err) != nil {
		return
	}
	public, err := Public(pair)
	if log.PrintError("crypto", err) != nil {
		return
	}
	private, err := Private(pair)
	if log.PrintError("crypto", err) != nil {
		return
	}
	if viper.GetBool("stdout") {
		fmt.Println(public)
		fmt.Println(private)
	}
	if viper.GetString("path") != "" {
		err := os.MkdirAll(viper.GetString("path"), 0700)
		if log.PrintError("crypto", err) != nil {
			return
		}
		err = ioutil.WriteFile(
			viper.GetString("path")+"public.pgp",
			[]byte(public),
			0600,
		)
		if log.PrintError("crypto", err) != nil {
			return
		}
		err = ioutil.WriteFile(
			viper.GetString("path")+"private.pgp",
			[]byte(private),
			0600,
		)
		if log.PrintError("crypto", err) != nil {
			return
		}
	}
}
