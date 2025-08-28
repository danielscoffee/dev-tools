package home
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
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	return true, nil
}
func (p *Page) GetTitle() string {
	return "Dev Tools - Home"
}
func (p *Page) GetKeyBindings() []types.KeyBinding {
	return []types.KeyBinding{
		{Key: "l", Description: "Languages", Action: "navigate_languages"},
		{Key: "c", Description: "Configuration", Action: "navigate_config"},
		{Key: "?", Description: "Help", Action: "navigate_help"},
	}
}
