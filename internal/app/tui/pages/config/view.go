package config

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
)

// Page represents the configuration page component
type Page struct {
	styles *PageStyles
}

// PageStyles defines styles for the configuration page
type PageStyles struct {
	Title       lipgloss.Style
	Description lipgloss.Style
	MenuItem    lipgloss.Style
	KeyBinding  lipgloss.Style
}

// NewPage creates a new configuration page instance
func NewPage() *Page {
	return &Page{
		styles: NewPageStyles(),
	}
}

// NewPageStyles creates default styles for the configuration page
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

// Render renders the configuration page
func (p *Page) Render(width, height int) string {
	var items []string

	items = append(items, "⚙️ Configuration")
	items = append(items, "")
	items = append(items, p.styles.Description.Render("Configure dev-tools settings and preferences."))
	items = append(items, "")

	// Available configuration options
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

// HandleInput handles input for the configuration page
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "t":
		// TODO: Implement theme configuration
		return true, nil
	case "k":
		// TODO: Implement keybinding configuration
		return true, nil
	case "p":
		// TODO: Implement path configuration
		return true, nil
	case "r":
		// TODO: Implement reset functionality
		return true, nil
	}
	return true, nil
}

// GetTitle returns the page title
func (p *Page) GetTitle() string {
	return "Configuration"
}

// GetKeyBindings returns the key bindings for this page
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "t", Description: "Theme", Action: "configure_theme"},
		{Key: "k", Description: "Keybindings", Action: "configure_keys"},
		{Key: "p", Description: "Paths", Action: "configure_paths"},
		{Key: "r", Description: "Reset", Action: "reset_config"},
	}
}
