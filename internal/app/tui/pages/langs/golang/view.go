// Package golang
package golang
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
	Feature     lipgloss.Style
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
		Feature: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#98FB98")).
			Padding(0, 2),
	}
}
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
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	return true, nil
}
func (p *Page) GetTitle() string {
	return "Go/Golang Tools"
}
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "b", Description: "Blueprint", Action: "navigate_blueprint"},
		{Key: "m", Description: "Modules", Action: "manage_modules"},
		{Key: "t", Description: "Tests", Action: "run_tests"},
		{Key: "f", Description: "Format", Action: "format_code"},
	}
}
