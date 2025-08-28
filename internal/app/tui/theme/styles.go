// Package theme
package theme
import "github.com/charmbracelet/lipgloss"
type Styles struct {
	App         lipgloss.Style
	Header      lipgloss.Style
	Subtext     lipgloss.Style
	Content     lipgloss.Style
	Footer      lipgloss.Style
	Status      lipgloss.Style
	Error       lipgloss.Style
	Success     lipgloss.Style
	Warning     lipgloss.Style
	Info        lipgloss.Style
	Panel       lipgloss.Style
	Button      lipgloss.Style
	ButtonFocus lipgloss.Style
}
func NewStyles(theme *Theme) *Styles {
	return &Styles{
		App: lipgloss.NewStyle().
			Padding(1, 2),
		Header: lipgloss.NewStyle().
			Bold(true).
			Padding(0, 1).
			MarginBottom(1),
		Subtext: lipgloss.NewStyle().
			Bold(false),
		Content: lipgloss.NewStyle().
			Padding(1, 0),
		Footer: lipgloss.NewStyle().
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder()).
			MarginTop(1).
			Padding(1, 0),
		Status: lipgloss.NewStyle().
			Padding(0, 1).
			Reverse(true),
		Error: lipgloss.NewStyle().
			Foreground(theme.Error).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(theme.Error),
		Success: lipgloss.NewStyle().
			Foreground(theme.Success).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(theme.Success),
		Warning: lipgloss.NewStyle().
			Foreground(theme.Warning).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(theme.Warning),
		Info: lipgloss.NewStyle().
			Foreground(theme.Info).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(theme.Info),
		Panel: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			MarginBottom(1),
		Button: lipgloss.NewStyle().
			Padding(0, 2).
			Border(lipgloss.RoundedBorder()),
		ButtonFocus: lipgloss.NewStyle().
			Padding(0, 2).
			Border(lipgloss.RoundedBorder()).
			Reverse(true).
			Bold(true),
	}
}
