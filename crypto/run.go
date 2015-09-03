package crypto

import (
  "fmt"

  "github.com/spf13/viper"

	"github.com/pdxjohnny/telem/log"
)

func Run() {
  if viper.GetBool("gen") {
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
  }
}
