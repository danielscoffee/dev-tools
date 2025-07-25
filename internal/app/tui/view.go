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
)

// Model represents the main TUI application model
type Model struct {
	router *Router
	ready  bool
	width  int
	height int
	styles *AppStyles
}

// AppStyles defines the main application styles
type AppStyles struct {
	App lipgloss.Style
}

// NewModel creates a new TUI model
func NewModel() *Model {
	router := NewRouter()

	// Register all routes with their page components
	router.RegisterRoute("/", home.NewPage(), "Dev Tools - Home", "Main menu and navigation", "h")
	router.RegisterRoute("/langs", langs.NewPage(), "Programming Languages", "Tools for different languages", "l")
	router.RegisterRoute("/langs/golang", golang.NewPage(), "Go/Golang Tools", "Go development tools", "g")
	router.RegisterRoute("/config", config.NewPage(), "Configuration", "Application settings", "c")
	router.RegisterRoute("/help", help.NewPage(), "Help & Documentation", "Usage instructions and help", "?")

	// TODO: Add more language routes
	// router.RegisterRoute("/langs/javascript", javascript.NewPage(), "JavaScript Tools", "JavaScript development tools", "j")
	// router.RegisterRoute("/langs/python", python.NewPage(), "Python Tools", "Python development tools", "p")

	return &Model{
		router: router,
		styles: NewAppStyles(),
	}
}

// NewAppStyles creates the main application styles
func NewAppStyles() *AppStyles {
	return &AppStyles{
		App: lipgloss.NewStyle().
			Padding(1, 2),
	}
}

// Init initializes the TUI model
func (m *Model) Init() tea.Cmd {
	return nil
}

// Update handles TUI events
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case tea.KeyMsg:
		// Let the router handle input
		if cont, cmd := m.router.HandleInput(msg); !cont {
			return m, cmd
		}
		return m, nil
	}

	return m, nil
}

// View renders the TUI
func (m *Model) View() string {
	if !m.ready {
		return "Loading..."
	}

	// Header
	header := m.router.RenderHeader(m.width)

	// Main content
	content := m.router.RenderCurrentPage(m.width, m.height)

	// Footer
	footer := m.router.RenderFooter(m.width)

	// Status bar
	status := m.router.RenderStatusBar(m.width)

	// Combine all sections
	app := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		m.router.styles.Content.Render(content),
		footer,
		status,
	)

	return m.styles.App.Render(app)
}

// Initialize starts the TUI application
func Initialize() error {
	model := NewModel()

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
