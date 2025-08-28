// Package config
package config
import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
)
type Page struct {
	styles *PageStyles
}
type PageStyles struct {
	Title       lipgloss.Style
	Description lipgloss.Style
	MenuItem    lipgloss.Style
	KeyBinding  lipgloss.Style
}
func NewPage() *Page {
	return &Page{
		styles: NewPageStyles(),
	}
}
func NewPageStyles() *PageStyles {
	return &PageStyles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00D7FF")).
			MarginBottom(1),
		Description: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC")).
			MarginBottom(1),
		MenuItem: lipgloss.NewStyle().
			Padding(0, 2).
			MarginBottom(1),
		KeyBinding: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00D7FF")).
			Background(lipgloss.Color("#1A1A1A")).
			Padding(0, 1),
	}
}
func (p *Page) Render(width, height int) string {
	var items []string
	items = append(items, "⚙️ Configuration")
	items = append(items, "")
	items = append(items, p.styles.Description.Render("Configure dev-tools settings and preferences."))
	items = append(items, "")
	keyStyle := p.styles.KeyBinding.Render("[t]")
	item := lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Theme - ",
		p.styles.Description.Render("Change application theme and colors"),
	)
	items = append(items, p.styles.MenuItem.Render(item))
	keyStyle = p.styles.KeyBinding.Render("[k]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Keybindings - ",
		p.styles.Description.Render("Configure keyboard shortcuts"),
	)
	items = append(items, p.styles.MenuItem.Render(item))
	keyStyle = p.styles.KeyBinding.Render("[p]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Paths - ",
		p.styles.Description.Render("Configure default project paths"),
	)
	items = append(items, p.styles.MenuItem.Render(item))
	keyStyle = p.styles.KeyBinding.Render("[r]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Reset - ",
		p.styles.Description.Render("Reset all settings to defaults"),
	)
	items = append(items, p.styles.MenuItem.Render(item))
	return lipgloss.JoinVertical(lipgloss.Left, items...)
}
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "t":
		return true, nil
	case "k":
		return true, nil
	case "p":
		return true, nil
	case "r":
		return true, nil
	}
	return true, nil
}
func (p *Page) GetTitle() string {
	return "Configuration"
}
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "t", Description: "Theme", Action: "configure_theme"},
		{Key: "k", Description: "Keybindings", Action: "configure_keys"},
		{Key: "p", Description: "Paths", Action: "configure_paths"},
		{Key: "r", Description: "Reset", Action: "reset_config"},
	}
}
