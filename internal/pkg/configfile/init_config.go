package configfile
import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)
type ConfigFile struct{}
func (c *ConfigFile) InitConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		cobra.CheckErr(err)
		fmt.Printf("\n Error getting user home directory: %v", err)
		os.Exit(1)
	}
	if _, err := os.Stat(home + "/.dev-tools.yaml"); os.IsNotExist(err) {
		os.Create(home + "/.dev-tools.yaml")
	}
	cobra.CheckErr(err)
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".dev-tools")
	viper.ReadInConfig()
}
