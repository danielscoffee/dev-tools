// Package langs
package langs
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
	items = append(items, "💻 Programming Languages")
	items = append(items, "")
	items = append(items, p.styles.Description.Render("Choose a programming language to see available tools and utilities."))
	items = append(items, "")
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
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	return true, nil
}
func (p *Page) GetTitle() string {
	return "Programming Languages"
}
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "g", Description: "Go/Golang", Action: "navigate_golang"},
		{Key: "j", Description: "JavaScript", Action: "navigate_javascript"},
		{Key: "p", Description: "Python", Action: "navigate_python"},
	}
}
