package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danielscoffee/dev-tools/internal/app/tui/pages"
)

// Model represents the main TUI application model
type Model struct {
	integration *pages.TUIIntegration
	ready       bool
	width       int
	height      int
	styles      *AppStyles
}

// AppStyles defines the main application styles
type AppStyles struct {
	App       lipgloss.Style
	Header    lipgloss.Style
	Content   lipgloss.Style
	Footer    lipgloss.Style
	StatusBar lipgloss.Style
}

// NewModel creates a new TUI model
func NewModel() *Model {
	// Initialize pages system
	pages.InitPages()

	return &Model{
		integration: pages.NewTUIIntegration(),
		styles:      NewAppStyles(),
	}
}

// NewAppStyles creates the main application styles
func NewAppStyles() *AppStyles {
	return &AppStyles{
		App: lipgloss.NewStyle().
			Padding(1, 2),

		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1).
			MarginBottom(1),

		Content: lipgloss.NewStyle().
			Padding(1, 0).
			Height(20),

		Footer: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#383838")).
			MarginTop(1).
			Padding(1, 0),

		StatusBar: lipgloss.NewStyle().
			Background(lipgloss.Color("#00D7FF")).
			Foreground(lipgloss.Color("#000000")).
			Padding(0, 1),
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
		// Let the integration handle input
		if cont, cmd := m.integration.HandleInput(msg); !cont {
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
	header := m.renderHeader()

	// Main content
	content := m.integration.RenderCurrentPage()

	// Footer
	footer := m.renderFooter()

	// Status bar
	status := m.renderStatusBar()

	// Combine all sections
	app := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		m.styles.Content.Render(content),
		footer,
		status,
	)

	return m.styles.App.Render(app)
}

// renderHeader renders the application header
func (m *Model) renderHeader() string {
	title := "🛠️  Dev Tools TUI"
	version := "v1.0.0"
	subtitle := "Developer Experience Enhancement Suite"

	headerTitle := lipgloss.JoinHorizontal(
		lipgloss.Left,
		title,
		strings.Repeat(" ", max(0, m.width-len(title)-len(version)-8)),
		version,
	)

	headerContent := lipgloss.JoinVertical(
		lipgloss.Left,
		headerTitle,
		m.styles.Header.
			Bold(false).
			Foreground(lipgloss.Color("#CCCCCC")).
			Render(subtitle),
	)

	return m.styles.Header.Width(m.width - 4).Render(headerContent)
}

// renderFooter renders the application footer
func (m *Model) renderFooter() string {
	return m.styles.Footer.Width(m.width - 4).Render(m.integration.GetFooter())
}

// renderStatusBar renders the status bar
func (m *Model) renderStatusBar() string {
	breadcrumb := m.getBreadcrumbString()

	statusContent := fmt.Sprintf("📍 %s", breadcrumb)

	return m.styles.StatusBar.Width(m.width - 4).Render(statusContent)
}

// getBreadcrumbString returns the breadcrumb as a string
func (m *Model) getBreadcrumbString() string {
	manager := pages.GetManager()
	breadcrumb := manager.GetBreadcrumb()

	if len(breadcrumb) == 0 {
		return "home"
	}

	var parts []string
	for _, page := range breadcrumb {
		parts = append(parts, page.Name)
	}

	return strings.Join(parts, " > ")
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
