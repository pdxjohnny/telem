package main

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/telem/commands"
)

func main() {
	var rootCmd = &cobra.Command{Use: "telem"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
