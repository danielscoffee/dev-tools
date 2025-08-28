package theme
import "github.com/charmbracelet/lipgloss"
type Theme struct {
	Name       string
	Primary    lipgloss.Color
	Secondary  lipgloss.Color
	Accent     lipgloss.Color
	Background lipgloss.Color
	Foreground lipgloss.Color
	Muted      lipgloss.Color
	Success    lipgloss.Color
	Warning    lipgloss.Color
	Error      lipgloss.Color
	Info       lipgloss.Color
}
func Themeless() *Theme {
	return &Theme{
		Name:       "Themeless",
		Primary:    lipgloss.Color(""),
		Secondary:  lipgloss.Color(""),
		Accent:     lipgloss.Color(""),
		Background: lipgloss.Color(""),
		Foreground: lipgloss.Color(""),
		Muted:      lipgloss.Color(""),
		Success:    lipgloss.Color("2"),
		Warning:    lipgloss.Color("3"),
		Error:      lipgloss.Color("1"),
		Info:       lipgloss.Color("4"),
	}
}
func Dark() *Theme {
	return Themeless()
}
func Light() *Theme {
	return Themeless()
}
func (t *Theme) GetResetSequence() string {
	return "\033[0m"
}
func (t *Theme) ApplyTerminalReset() lipgloss.Style {
	return lipgloss.NewStyle()
}
