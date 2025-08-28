package cli
import (
	"os"
	"github.com/danielscoffee/dev-tools/internal/pkg/configfile"
	"github.com/spf13/cobra"
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
	cf := &configfile.ConfigFile{}
	cobra.OnInitialize(cf.InitConfig)
}
