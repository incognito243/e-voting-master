package main

import (
	"github.com/spf13/cobra"

	"e-voting-mater/cmd/api"
	"e-voting-mater/configs"
)

var (
	// Used for flags.
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:          "e-voting-master",
	Short:        "",
	Long:         "",
	SilenceUsage: true,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(func() { configs.Init(cfgFile) })
	rootCmd.AddCommand(api.Command)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (required)")
	if err := rootCmd.MarkFlagRequired("config"); err != nil {
		return
	}
}
