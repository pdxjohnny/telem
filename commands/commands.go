package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/telem/crypto"
	"github.com/pdxjohnny/telem/frontend"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "frontend",
		Short: "Recives records from clients",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			frontend.Run()
		},
	},
	&cobra.Command{
		Use:   "crypto",
		Short: "Generats pgp key pairs",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			crypto.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
