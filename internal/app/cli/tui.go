package cli

import (
	"fmt"
	"log"

	"github.com/danielscoffee/dev-tools/internal/app/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start a SUPER AWESOME TUI mode",
	Long:  "Launch the Terminal User Interface for an interactive development tools experience",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting Dev Tools TUI...")

		if err := tui.Initialize(); err != nil {
			log.Fatalf("Failed to start TUI: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
