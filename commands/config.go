package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConfigOptions = map[string]interface{}{
	"frontend": map[string]interface{}{
		"port": map[string]interface{}{
			"value": 25000,
			"help":  "Port to listen on",
		},
		"address": map[string]interface{}{
			"value": "0.0.0.0",
			"help":  "Address  to bind to",
		},
	},
}

func ConfigDefaults(cmdList ...*cobra.Command) {
	ConfigEnv()
	ConfigSet()
	ConfigFlags(cmdList...)
}

func ConfigSet() {
	for index, item := range ConfigOptions {
		opt := item.(map[string]interface{})
		viper.SetDefault(index, opt["value"])
	}
}

func ConfigFlags(cmdList ...*cobra.Command) {
	for _, cmd := range cmdList {
		for index, item := range ConfigOptions[cmd.Use].(map[string]interface{}) {
			opt := item.(map[string]interface{})
			help := opt["help"].(string)
			switch value := opt["value"].(type) {
			case int:
				cmd.Flags().Int(index, value, help)
			case bool:
				cmd.Flags().Bool(index, value, help)
			case string:
				cmd.Flags().String(index, value, help)
			case float32:
				cmd.Flags().Float32(index, value, help)
			case float64:
				cmd.Flags().Float64(index, value, help)
			default:
			}
		}
	}
}

func ConfigBindFlags(cmd *cobra.Command) {
	for index, _ := range ConfigOptions {
		viper.BindPFlag(index, cmd.Flags().Lookup(index))
	}
}

func ConfigEnv() {
	viper.SetEnvPrefix("telem")
	viper.AutomaticEnv()
}
