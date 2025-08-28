package cli
import (
	"fmt"
	"log"
	"os"
	"github.com/danielscoffee/dev-tools/internal/app/tui"
	"github.com/spf13/cobra"
)
var (
	themeFlag string
)
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Start a SUPER AWESOME TUI mode",
	Long:  "Launch the Terminal User Interface for an interactive development tools experience",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Starting Dev Tools TUI...")
		selectedTheme := themeFlag
		if selectedTheme == "" {
			selectedTheme = os.Getenv("DEV_TOOLS_THEME")
		}
		if selectedTheme == "" {
			selectedTheme = "dark"
		}
		if err := tui.InitializeWithTheme(selectedTheme); err != nil {
			log.Fatalf("Failed to start TUI: %v", err)
		}
	},
}
func init() {
	tuiCmd.Flags().StringVarP(&themeFlag, "theme", "t", "", "Set theme (dark/light)")
	rootCmd.AddCommand(tuiCmd)
}
