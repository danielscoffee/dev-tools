package golang

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
)

// Page represents the golang page component
type Page struct {
	styles *PageStyles
}

// PageStyles defines styles for the golang page
type PageStyles struct {
	Title       lipgloss.Style
	Description lipgloss.Style
	MenuItem    lipgloss.Style
	KeyBinding  lipgloss.Style
	Feature     lipgloss.Style
}

// NewPage creates a new golang page instance
func NewPage() *Page {
	return &Page{
		styles: NewPageStyles(),
	}
}

// NewPageStyles creates default styles for the golang page
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

		Feature: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#98FB98")).
			Padding(0, 2),
	}
}

// Render renders the golang page
func (p *Page) Render(width, height int) string {
	var items []string

	items = append(items, "🐹 Go/Golang Development Tools")
	items = append(items, "")
	items = append(items, p.styles.Description.Render("Integrated Go development tools including go-blueprint project scaffolding"))
	items = append(items, "")
	items = append(items, "🏗️  Project Creation:")
	items = append(items, p.styles.Feature.Render("  • REST API projects"))
	items = append(items, p.styles.Feature.Render("  • CLI applications"))
	items = append(items, p.styles.Feature.Render("  • Web applications with Fiber/Gin"))
	items = append(items, "")

	// Available actions
	keyStyle := p.styles.KeyBinding.Render("[b]")
	item := lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Blueprint - ",
		p.styles.Description.Render("Create Go projects with go-blueprint"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	keyStyle = p.styles.KeyBinding.Render("[m]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Modules - ",
		p.styles.Description.Render("Manage go.mod and dependencies"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	keyStyle = p.styles.KeyBinding.Render("[t]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Tests - ",
		p.styles.Description.Render("Run tests and benchmarks"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	keyStyle = p.styles.KeyBinding.Render("[f]")
	item = lipgloss.JoinHorizontal(
		lipgloss.Left,
		keyStyle,
		" Format - ",
		p.styles.Description.Render("Format code with gofmt/goimports"),
	)
	items = append(items, p.styles.MenuItem.Render(item))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// HandleInput handles input for the golang page
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	// Let the router handle navigation for blueprint and other actions
	return true, nil
}

// GetTitle returns the page title
func (p *Page) GetTitle() string {
	return "Go/Golang Tools"
}

// GetKeyBindings returns the key bindings for this page
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "b", Description: "Blueprint", Action: "navigate_blueprint"},
		{Key: "m", Description: "Modules", Action: "manage_modules"},
		{Key: "t", Description: "Tests", Action: "run_tests"},
		{Key: "f", Description: "Format", Action: "format_code"},
	}
}
