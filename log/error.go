package log

import (
  "fmt"
)

func PrintError(prepend string, err error)  {
  if err != nil {
    fmt.Printf("%s:\t%s\n", prepend, err)
  }
}
