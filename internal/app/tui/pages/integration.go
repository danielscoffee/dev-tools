package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TUIIntegration connects the page system with the TUI
type TUIIntegration struct {
	manager *PageManager
	styles  *TUIStyles
}

// TUIStyles defines styles for the TUI
type TUIStyles struct {
	Title       lipgloss.Style
	Breadcrumb  lipgloss.Style
	MenuItem    lipgloss.Style
	Selected    lipgloss.Style
	Description lipgloss.Style
	KeyBinding  lipgloss.Style
	Border      lipgloss.Style
}

// NewTUIIntegration creates a new TUI integration
func NewTUIIntegration() *TUIIntegration {
	return &TUIIntegration{
		manager: GetManager(),
		styles:  NewTUIStyles(),
	}
}

// NewTUIStyles creates default styles for the TUI
func NewTUIStyles() *TUIStyles {
	return &TUIStyles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00D7FF")).
			MarginBottom(1),

		Breadcrumb: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			MarginBottom(1),

		MenuItem: lipgloss.NewStyle().
			Padding(0, 2).
			MarginBottom(1),

		Selected: lipgloss.NewStyle().
			Background(lipgloss.Color("#00D7FF")).
			Foreground(lipgloss.Color("#000000")).
			Padding(0, 2).
			MarginBottom(1),

		Description: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Italic(true),

		KeyBinding: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true),

		Border: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00D7FF")).
			Padding(1, 2),
	}
}

// RenderCurrentPage renders the current page
func (ti *TUIIntegration) RenderCurrentPage() string {
	current := ti.manager.GetCurrentPage()

	// Page title
	title := ti.styles.Title.Render(current.Title)

	// Breadcrumb
	breadcrumb := ti.renderBreadcrumb()

	// Page-specific content
	content := ti.renderPageContent()

	// Combine everything
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		breadcrumb,
		"",
		content,
	)
}

// renderBreadcrumb renders the navigation path
func (ti *TUIIntegration) renderBreadcrumb() string {
	breadcrumb := ti.manager.GetBreadcrumb()
	if len(breadcrumb) <= 1 {
		return ""
	}

	var parts []string
	for i, page := range breadcrumb {
		if i == len(breadcrumb)-1 {
			// Current page highlighted
			parts = append(parts, ti.styles.KeyBinding.Render(page.Name))
		} else {
			parts = append(parts, page.Name)
		}
	}

	path := lipgloss.JoinHorizontal(lipgloss.Left, parts...)
	return ti.styles.Breadcrumb.Render("📍 " + path)
}

// renderPageContent renders page-specific content
func (ti *TUIIntegration) renderPageContent() string {
	currentType := ti.manager.GetCurrentPageType()

	switch currentType {
	case HomePage:
		return ti.renderHomePage()
	case LanguagesPage:
		return ti.renderLanguagesPage()
	case GolangPage:
		return ti.renderGolangPage()
	case JavascriptPage:
		return ti.renderJavascriptPage()
	case PythonPage:
		return ti.renderPythonPage()
	case ToolsPage:
		return ti.renderToolsPage()
	case ConfigPage:
		return ti.renderConfigPage()
	case HelpPage:
		return ti.renderHelpPage()
	default:
		return "Page not implemented"
	}
}

// renderHomePage renders the home page
func (ti *TUIIntegration) renderHomePage() string {
	children := ti.manager.GetChildrenInfo(HomePage)

	var items []string
	items = append(items, "🏠 Welcome to Dev Tools!")
	items = append(items, "")
	items = append(items, ti.styles.Description.Render("A comprehensive suite of development tools designed to enhance your"))
	items = append(items, ti.styles.Description.Render("developer experience across multiple programming languages and technologies."))
	items = append(items, "")
	items = append(items, "🚀 Features:")
	items = append(items, "  • Project scaffolding with go-blueprint")
	items = append(items, "  • Multi-language development tools")
	items = append(items, "  • Docker containerization utilities")
	items = append(items, "  • Configuration management")
	items = append(items, "")
	items = append(items, "📋 Choose an option:")
	items = append(items, "")

	for _, child := range children {
		keyStyle := ti.styles.KeyBinding.Render("[" + child.KeyBinding + "]")
		item := lipgloss.JoinHorizontal(
			lipgloss.Left,
			keyStyle,
			" ",
			child.Title,
			" - ",
			ti.styles.Description.Render(child.Description),
		)
		items = append(items, ti.styles.MenuItem.Render(item))
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// renderLanguagesPage renders the languages page
func (ti *TUIIntegration) renderLanguagesPage() string {
	children := ti.manager.GetChildrenInfo(LanguagesPage)

	var items []string
	items = append(items, "💻 Programming Languages")
	items = append(items, "")

	for _, child := range children {
		keyStyle := ti.styles.KeyBinding.Render("[" + child.KeyBinding + "]")
		item := lipgloss.JoinHorizontal(
			lipgloss.Left,
			keyStyle,
			" ",
			child.Title,
		)
		items = append(items, ti.styles.MenuItem.Render(item))
	}

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// renderGolangPage renders the Go page
func (ti *TUIIntegration) renderGolangPage() string {
	var items []string
	items = append(items, "🐹 Go/Golang Development Tools")
	items = append(items, "")
	items = append(items, ti.styles.Description.Render("Integrated Go development tools including go-blueprint project scaffolding"))
	items = append(items, "")
	items = append(items, "🏗️  Project Creation:")
	items = append(items, ti.styles.MenuItem.Render("[b] Blueprint - Create Go projects with go-blueprint"))
	items = append(items, ti.styles.MenuItem.Render("    • REST API projects"))
	items = append(items, ti.styles.MenuItem.Render("    • CLI applications"))
	items = append(items, ti.styles.MenuItem.Render("    • Web applications with Fiber/Gin"))
	items = append(items, "")
	items = append(items, "🔧 Development Tools:")
	items = append(items, ti.styles.MenuItem.Render("[m] Modules - Manage go.mod and dependencies"))
	items = append(items, ti.styles.MenuItem.Render("[t] Tests - Run tests and benchmarks"))
	items = append(items, ti.styles.MenuItem.Render("[f] Format - Format code with gofmt/goimports"))
	items = append(items, ti.styles.MenuItem.Render("[l] Lint - Code analysis with golangci-lint"))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// renderJavascriptPage renders the JavaScript page
func (ti *TUIIntegration) renderJavascriptPage() string {
	var items []string
	items = append(items, "⚡ JavaScript/Node.js Tools")
	items = append(items, "")
	items = append(items, ti.styles.MenuItem.Render("[n] NPM - Manage packages"))
	items = append(items, ti.styles.MenuItem.Render("[v] Vite - Create Vite project"))
	items = append(items, ti.styles.MenuItem.Render("[r] React - Create React app"))
	items = append(items, ti.styles.MenuItem.Render("[e] ESLint - Configure linting"))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// renderPythonPage renders the Python page
func (ti *TUIIntegration) renderPythonPage() string {
	var items []string
	items = append(items, "🐍 Python Tools")
	items = append(items, "")
	items = append(items, ti.styles.MenuItem.Render("[v] Venv - Manage virtual environments"))
	items = append(items, ti.styles.MenuItem.Render("[p] Pip - Manage packages"))
	items = append(items, ti.styles.MenuItem.Render("[f] Flask - Create Flask app"))
	items = append(items, ti.styles.MenuItem.Render("[d] Django - Create Django project"))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// renderToolsPage renders the tools page
func (ti *TUIIntegration) renderToolsPage() string {
	var items []string
	items = append(items, "🔧 General Tools")
	items = append(items, "")
	items = append(items, ti.styles.MenuItem.Render("[d] Docker - Manage containers"))
	items = append(items, ti.styles.MenuItem.Render("[g] Git - Git operations"))
	items = append(items, ti.styles.MenuItem.Render("[s] SSH - SSH configurations"))
	items = append(items, ti.styles.MenuItem.Render("[e] Editor - Configure editors"))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// renderConfigPage renders the configuration page
func (ti *TUIIntegration) renderConfigPage() string {
	var items []string
	items = append(items, "⚙️ Configuration")
	items = append(items, "")
	items = append(items, ti.styles.MenuItem.Render("[t] Theme - Change theme"))
	items = append(items, ti.styles.MenuItem.Render("[k] Keys - Configure shortcuts"))
	items = append(items, ti.styles.MenuItem.Render("[p] Paths - Configure paths"))
	items = append(items, ti.styles.MenuItem.Render("[r] Reset - Restore defaults"))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// renderHelpPage renders the help page
func (ti *TUIIntegration) renderHelpPage() string {
	var items []string
	items = append(items, "❓ Help & Documentation")
	items = append(items, "")
	items = append(items, "🎯 Navigation:")
	items = append(items, ti.styles.MenuItem.Render("  [h] - Home"))
	items = append(items, ti.styles.MenuItem.Render("  [l] - Languages"))
	items = append(items, ti.styles.MenuItem.Render("  [t] - Tools"))
	items = append(items, ti.styles.MenuItem.Render("  [esc/q] - Go back"))
	items = append(items, ti.styles.MenuItem.Render("  [ctrl+c] - Exit"))
	items = append(items, "")
	items = append(items, "📚 Features:")
	items = append(items, ti.styles.MenuItem.Render("  • Project creation"))
	items = append(items, ti.styles.MenuItem.Render("  • Dependency management"))
	items = append(items, ti.styles.MenuItem.Render("  • Customizable settings"))

	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// HandleInput processes user input
func (ti *TUIIntegration) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return false, tea.Quit
	case "esc", "q":
		if ti.manager.GoBack() {
			return true, nil
		}
		return false, tea.Quit
	default:
		// Try to navigate using key bindings
		if HandleKeyBinding(msg.String()) {
			return true, nil
		}
	}

	return true, nil
}

// GetFooter returns the footer with instructions
func (ti *TUIIntegration) GetFooter() string {
	footer := "💡 [esc/q] Go back | [ctrl+c] Exit | [?] Help"
	return ti.styles.Description.Render(footer)
}
