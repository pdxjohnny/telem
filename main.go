package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/telem/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "telem"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
