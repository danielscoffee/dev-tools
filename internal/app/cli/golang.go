package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// golangCmd represents the golang command
var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please select a subcommand for golang")
	},
}

func init() {
	rootCmd.AddCommand(golangCmd)
}
