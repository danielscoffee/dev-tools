package tui
import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/config"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/help"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/home"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/langs"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/langs/golang"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/langs/golang/blueprint"
	"github.com/danielscoffee/dev-tools/internal/app/tui/theme"
)
type Model struct {
	router *Router
	ready  bool
	width  int
	height int
	styles *AppStyles
	theme  *theme.Theme
}
type AppStyles struct {
	App lipgloss.Style
}
func NewModel() *Model {
	return NewModelWithTheme("themeless")
}
func NewModelWithTheme(themeName string) *Model {
	router := NewRouter()
	currentTheme := theme.Themeless()
	router.RegisterRoute("/", home.NewPage(), "Dev Tools - Home", "Main menu and navigation", "h")
	router.RegisterRoute("/langs", langs.NewPage(), "Programming Languages", "Tools for different languages", "l")
	router.RegisterRoute("/langs/golang", golang.NewPage(), "Go/Golang Tools", "Go development tools", "g")
	router.RegisterRoute("/langs/golang/blueprint", blueprint.NewPage(), "Go Blueprint Creator", "Create Go projects with go-blueprint", "b")
	router.RegisterRoute("/config", config.NewPage(), "Configuration", "Application settings", "c")
	router.RegisterRoute("/help", help.NewPage(), "Help & Documentation", "Usage instructions and help", "?")
	return &Model{
		router: router,
		styles: NewAppStyles(currentTheme),
		theme:  currentTheme,
	}
}
func NewAppStyles(t *theme.Theme) *AppStyles {
	return &AppStyles{
		App: lipgloss.NewStyle().
			Padding(1, 2).
			Background(t.Background).
			Foreground(t.Foreground),
	}
}
func (m *Model) Init() tea.Cmd {
	return nil
}
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil
	case tea.KeyMsg:
		if msg.String() == "t" {
			if m.theme.Name == "Dark" {
				m.theme = theme.Light()
			} else {
				m.theme = theme.Dark()
			}
			m.styles = NewAppStyles(m.theme)
			m.router.UpdateTheme(m.theme)
			return m, nil
		}
		if msg.String() == "?" {
			return m, nil
		}
		if cont, cmd := m.router.HandleInput(msg); !cont {
			return m, cmd
		}
		return m, nil
	default:
		currentRoute := m.router.GetCurrentRoute()
		if currentRoute != nil && currentRoute.Path == "/langs/golang/blueprint" {
			if blueprintPage, ok := currentRoute.Component.(*blueprint.Page); ok {
				return m, blueprintPage.Update(msg)
			}
		}
	}
	return m, nil
}
func (m *Model) View() string {
	if !m.ready {
		return "Loading..."
	}
	header := m.router.RenderHeader(m.width)
	content := m.router.RenderCurrentPage(m.width, m.height)
	footer := m.router.RenderFooter(m.width)
	status := m.router.RenderStatusBar(m.width)
	app := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		m.router.styles.Content.Render(content),
		footer,
		status,
	)
	return m.styles.App.Render(app)
}
func Initialize() error {
	return InitializeWithTheme("dark")
}
func InitializeWithTheme(themeName string) error {
	model := NewModelWithTheme(themeName)
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("failed to run TUI: %w", err)
	}
	return nil
}
