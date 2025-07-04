package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type CLI struct{}

var rootCmd = &cobra.Command{
	Use:   "dev-tools",
	Short: "Compilation of tools that give a AWESOME developer experience",
}

func (c CLI) Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initConfig)
}

// BUG: This is temporary configInit
func initConfig() {
	cfgFile := ".dev-tools.yaml"

	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		newFile, err := os.Create(cfgFile)
		if err != nil {
			panic(err)
		}
		defer newFile.Close()
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
