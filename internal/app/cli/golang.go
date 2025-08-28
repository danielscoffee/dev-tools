package cli
import (
	"fmt"
	"github.com/spf13/cobra"
)
var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please select a subcommand for golang")
	},
}
var buildCmd = &cobra.Command{
	Use:   "build golang project",
	Short: "Builds a Golang project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building Golang project...")
	},
}
func init() {
	rootCmd.AddCommand(golangCmd)
	golangCmd.AddCommand(buildCmd)
}
