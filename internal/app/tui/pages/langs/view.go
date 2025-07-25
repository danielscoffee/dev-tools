package langs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
)

// Page represents the languages page component
type Page struct {
	styles *PageStyles
}

// PageStyles defines styles for the languages page
type PageStyles struct {
	Title       lipgloss.Style
	Description lipgloss.Style
	MenuItem    lipgloss.Style
	KeyBinding  lipgloss.Style
}

// NewPage creates a new languages page instance
func NewPage() *Page {
	return &Page{
		styles: NewPageStyles(),
	}
}

// NewPageStyles creates default styles for the languages page
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

// Render renders the languages page
func (p *Page) Render(width, height int) string {
	var items []string

	items = append(items, "💻 Programming Languages")
	items = append(items, "")
	items = append(items, p.styles.Description.Render("Choose a programming language to see available tools and utilities."))
	items = append(items, "")

	// Available language routes
	keyStyle := p.styles.KeyBinding.Render("[g]")
	item := lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Go/Golang - ",
		p.styles.Description.Render("Go development tools with go-blueprint integration"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	keyStyle = p.styles.KeyBinding.Render("[j]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" JavaScript - ",
		p.styles.Description.Render("JavaScript/Node.js development tools"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	keyStyle = p.styles.KeyBinding.Render("[p]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Python - ",
		p.styles.Description.Render("Python development tools"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// HandleInput handles input for the languages page
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	// Languages page doesn't handle any specific input, let router handle it
	return true, nil
}

// GetTitle returns the page title
func (p *Page) GetTitle() string {
	return "Programming Languages"
}

// GetKeyBindings returns the key bindings for this page
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "g", Description: "Go/Golang", Action: "navigate_golang"},
		{Key: "j", Description: "JavaScript", Action: "navigate_javascript"},
		{Key: "p", Description: "Python", Action: "navigate_python"},
	}
}
