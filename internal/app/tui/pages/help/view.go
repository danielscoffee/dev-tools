package help

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
)

// Page represents the help page component
type Page struct {
	styles *PageStyles
}

// PageStyles defines styles for the help page
type PageStyles struct {
	Title       lipgloss.Style
	Description lipgloss.Style
	MenuItem    lipgloss.Style
	KeyBinding  lipgloss.Style
	Section     lipgloss.Style
}

// NewPage creates a new help page instance
func NewPage() *Page {
	return &Page{
		styles: NewPageStyles(),
	}
}

// NewPageStyles creates default styles for the help page
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

		Section: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#98FB98")).
			MarginBottom(1),
	}
}

// Render renders the help page
func (p *Page) Render(width, height int) string {
	var items []string

	items = append(items, "❓ Help & Documentation")
	items = append(items, "")
	items = append(items, p.styles.Description.Render("Welcome to Dev Tools TUI - your comprehensive development toolkit."))
	items = append(items, "")

	items = append(items, p.styles.Section.Render("🎯 Navigation:"))
	items = append(items, p.styles.MenuItem.Render("  [h] - Home"))
	items = append(items, p.styles.MenuItem.Render("  [l] - Programming Languages"))
	items = append(items, p.styles.MenuItem.Render("  [c] - Configuration"))
	items = append(items, p.styles.MenuItem.Render("  [esc] - Go back"))
	items = append(items, p.styles.MenuItem.Render("  [ctrl+c] - Exit"))
	items = append(items, "")

	items = append(items, p.styles.Section.Render("🚀 Features:"))
	items = append(items, p.styles.MenuItem.Render("  • Project scaffolding with go-blueprint"))
	items = append(items, p.styles.MenuItem.Render("  • Multi-language development tools"))
	items = append(items, p.styles.MenuItem.Render("  • Router-based page system"))
	items = append(items, p.styles.MenuItem.Render("  • Dynamic page discovery"))
	items = append(items, p.styles.MenuItem.Render("  • Customizable settings"))
	items = append(items, "")

	items = append(items, p.styles.Section.Render("🛠️ Architecture:"))
	items = append(items, p.styles.MenuItem.Render("  • Modular page components"))
	items = append(items, p.styles.MenuItem.Render("  • Interface-based design"))
	items = append(items, p.styles.MenuItem.Render("  • Hierarchical routing"))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// HandleInput handles input for the help page
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	// Help page doesn't handle any specific input, let router handle it
	return true, nil
}

// GetTitle returns the page title
func (p *Page) GetTitle() string {
	return "Help & Documentation"
}

// GetKeyBindings returns the key bindings for this page
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "h", Description: "Home", Action: "navigate_home"},
		{Key: "l", Description: "Languages", Action: "navigate_languages"},
		{Key: "c", Description: "Configuration", Action: "navigate_config"},
	}
}
