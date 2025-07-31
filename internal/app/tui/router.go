package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
)

// Route represents a TUI route
type Route struct {
	Path        string
	Component   types.PageRenderer
	Title       string
	Description string
	KeyBinding  string
}

// Router manages TUI routes and navigation
type Router struct {
	routes       map[string]*Route
	currentRoute string
	history      []string
	styles       *RouterStyles
}

// RouterStyles defines styles for the router
type RouterStyles struct {
	Header    lipgloss.Style
	Content   lipgloss.Style
	Footer    lipgloss.Style
	StatusBar lipgloss.Style
	Error     lipgloss.Style
}

// NewRouter creates a new router instance
func NewRouter() *Router {
	return &Router{
		routes:       make(map[string]*Route),
		currentRoute: "/",
		history:      make([]string, 0),
		styles:       NewRouterStyles(),
	}
}

// NewRouterStyles creates default router styles
func NewRouterStyles() *RouterStyles {
	return &RouterStyles{
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

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Background(lipgloss.Color("#2D1B1B")).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF6B6B")),
	}
}

// RegisterRoute registers a new route
func (r *Router) RegisterRoute(path string, component types.PageRenderer, title, description, keyBinding string) {
	r.routes[path] = &Route{
		Path:        path,
		Component:   component,
		Title:       title,
		Description: description,
		KeyBinding:  keyBinding,
	}
}

// NavigateTo navigates to a specific route
func (r *Router) NavigateTo(path string) error {
	if _, exists := r.routes[path]; !exists {
		return fmt.Errorf("route '%s' not found", path)
	}

	// Add current route to history
	if r.currentRoute != path {
		r.history = append(r.history, r.currentRoute)
	}

	r.currentRoute = path
	return nil
}

// GoBack navigates back to the previous route
func (r *Router) GoBack() error {
	if len(r.history) == 0 {
		return fmt.Errorf("no previous route in history")
	}

	// Get the last route from history
	lastRoute := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]

	r.currentRoute = lastRoute
	return nil
}

// GetCurrentRoute returns the current route
func (r *Router) GetCurrentRoute() *Route {
	return r.routes[r.currentRoute]
}

// GetAllRoutes returns all registered routes
func (r *Router) GetAllRoutes() map[string]*Route {
	return r.routes
}

// HandleInput handles key input and routes to appropriate handlers
func (r *Router) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return false, tea.Quit
	case "esc":
		if err := r.GoBack(); err != nil {
			// If can't go back, try to go to home
			r.NavigateTo("/")
		}
		return true, nil
	}

	// Check for contextually relevant route key bindings
	// Only check routes that are accessible from the current context
	availableRoutes := r.getAvailableRoutes()
	for path, route := range availableRoutes {
		if route.KeyBinding == msg.String() {
			r.NavigateTo(path)
			return true, nil
		}
	}

	// Let the current page handle the input
	if currentRoute := r.GetCurrentRoute(); currentRoute != nil {
		return currentRoute.Component.HandleInput(msg)
	}

	return true, nil
}

// getAvailableRoutes returns routes that are accessible from the current route
func (r *Router) getAvailableRoutes() map[string]*Route {
	available := make(map[string]*Route)
	
	switch r.currentRoute {
	case "/":
		// From home, can access top-level routes
		for path, route := range r.routes {
			if path == "/langs" || path == "/config" || path == "/help" {
				available[path] = route
			}
		}
	case "/langs":
		// From languages, can access language-specific routes
		for path, route := range r.routes {
			if strings.HasPrefix(path, "/langs/") && strings.Count(path, "/") == 2 {
				available[path] = route
			}
		}
	case "/langs/golang":
		// From golang, can access golang-specific routes
		for path, route := range r.routes {
			if strings.HasPrefix(path, "/langs/golang/") && strings.Count(path, "/") == 3 {
				available[path] = route
			}
		}
	default:
		// For other routes, no child routes are directly accessible via key bindings
		// They need to use escape to go back or handle their own navigation
	}
	
	return available
}

// RenderCurrentPage renders the current page
func (r *Router) RenderCurrentPage(width, height int) string {
	currentRoute := r.GetCurrentRoute()
	if currentRoute == nil {
		return r.styles.Error.Render("Error: Route not found")
	}

	return currentRoute.Component.Render(width, height)
}

// RenderHeader renders the application header
func (r *Router) RenderHeader(width int) string {
	title := "🛠️  Dev Tools TUI"
	version := "v1.0.0"
	subtitle := "Developer Experience Enhancement Suite"

	currentRoute := r.GetCurrentRoute()
	if currentRoute != nil && currentRoute.Title != "" {
		subtitle = currentRoute.Title
	}

	headerTitle := lipgloss.JoinHorizontal(
		lipgloss.Left,
		title,
		strings.Repeat(" ", max(0, width-len(title)-len(version)-8)),
		version,
	)

	headerContent := lipgloss.JoinVertical(
		lipgloss.Left,
		headerTitle,
		r.styles.Header.
			Bold(false).
			Foreground(lipgloss.Color("#CCCCCC")).
			Render(subtitle),
	)

	return r.styles.Header.Width(width - 4).Render(headerContent)
}

// RenderFooter renders the application footer
func (r *Router) RenderFooter(width int) string {
	var footerItems []string

	// Add current page key bindings
	if currentRoute := r.GetCurrentRoute(); currentRoute != nil {
		keyBindings := currentRoute.Component.GetKeyBindings()
		for _, kb := range keyBindings {
			footerItems = append(footerItems, fmt.Sprintf("[%s] %s", kb.Key, kb.Description))
		}
	}

	// Add global key bindings
	footerItems = append(footerItems, "[esc] Go back", "[ctrl+c] Exit")

	footerContent := "💡 " + strings.Join(footerItems, " | ")
	return r.styles.Footer.Width(width - 4).Render(footerContent)
}

// RenderStatusBar renders the status bar with breadcrumb
func (r *Router) RenderStatusBar(width int) string {
	breadcrumb := r.GetBreadcrumb()
	statusContent := fmt.Sprintf("📍 %s", breadcrumb)
	return r.styles.StatusBar.Width(width - 4).Render(statusContent)
}

// GetBreadcrumb returns the current navigation breadcrumb
func (r *Router) GetBreadcrumb() string {
	if r.currentRoute == "/" {
		return "home"
	}

	// Convert route path to breadcrumb
	parts := strings.Split(strings.Trim(r.currentRoute, "/"), "/")
	return strings.Join(parts, " > ")
}
