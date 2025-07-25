package home

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
)

// Page represents the home page component
type Page struct {
	styles *PageStyles
}

// PageStyles defines styles for the home page
type PageStyles struct {
	Title       lipgloss.Style
	Description lipgloss.Style
	MenuItem    lipgloss.Style
	KeyBinding  lipgloss.Style
}

// NewPage creates a new home page instance
func NewPage() *Page {
	return &Page{
		styles: NewPageStyles(),
	}
}

// NewPageStyles creates default styles for the home page
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

// Render renders the home page
func (p *Page) Render(width, height int) string {
	var items []string

	items = append(items, "🏠 Welcome to Dev Tools!")
	items = append(items, "")
	items = append(items, p.styles.Description.Render("A comprehensive suite of development tools designed to enhance your"))
	items = append(items, p.styles.Description.Render("developer experience across multiple programming languages and technologies."))
	items = append(items, "")
	items = append(items, "🚀 Features:")
	items = append(items, "  • Project scaffolding with go-blueprint")
	items = append(items, "  • Multi-language development tools")
	items = append(items, "  • Docker containerization utilities")
	items = append(items, "  • Configuration management")
	items = append(items, "")
	items = append(items, "📋 Choose an option:")
	items = append(items, "")

	// Available routes
	keyStyle := p.styles.KeyBinding.Render("[l]")
	item := lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Programming Languages - ",
		p.styles.Description.Render("Tools for different languages"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	keyStyle = p.styles.KeyBinding.Render("[c]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Configuration - ",
		p.styles.Description.Render("Application settings"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	keyStyle = p.styles.KeyBinding.Render("[?]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Help & Documentation - ",
		p.styles.Description.Render("Usage instructions and help"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// HandleInput handles input for the home page
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	// Home page doesn't handle any specific input, let router handle it
	return true, nil
}

// GetTitle returns the page title
func (p *Page) GetTitle() string {
	return "Dev Tools - Home"
}

// GetKeyBindings returns the key bindings for this page
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "l", Description: "Languages", Action: "navigate_languages"},
		{Key: "c", Description: "Configuration", Action: "navigate_config"},
		{Key: "?", Description: "Help", Action: "navigate_help"},
	}
}
