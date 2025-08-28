// Package tui
package tui
import (
	"fmt"
	"strings"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/danielscoffee/dev-tools/internal/app/tui/theme"
	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
)
type Route struct {
	Path        string
	Component   types.PageRenderer
	Title       string
	Description string
	KeyBinding  string
}
type Router struct {
	routes       map[string]*Route
	currentRoute string
	history      []string
	styles       *RouterStyles
	theme        *theme.Theme
}
type RouterStyles struct {
	Header    lipgloss.Style
	Content   lipgloss.Style
	Footer    lipgloss.Style
	StatusBar lipgloss.Style
	Error     lipgloss.Style
}
func NewRouter() *Router {
	currentTheme := theme.Themeless()
	return &Router{
		routes:       make(map[string]*Route),
		currentRoute: "/",
		history:      make([]string, 0),
		styles:       NewRouterStyles(currentTheme),
		theme:        currentTheme,
	}
}
func NewRouterStyles(t *theme.Theme) *RouterStyles {
	return &RouterStyles{
		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(t.Foreground).
			Background(t.Primary).
			Padding(0, 1).
			MarginBottom(1),
		Content: lipgloss.NewStyle().
			Padding(1, 0).
			Height(20).
			Foreground(t.Foreground),
		Footer: lipgloss.NewStyle().
			Foreground(t.Muted).
			BorderTop(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(t.Muted).
			MarginTop(1).
			Padding(1, 0),
		StatusBar: lipgloss.NewStyle().
			Background(t.Secondary).
			Foreground(t.Background).
			Padding(0, 1),
		Error: lipgloss.NewStyle().
			Foreground(t.Error).
			Background(t.Background).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(t.Error),
	}
}
func (r *Router) RegisterRoute(path string, component types.PageRenderer, title, description, keyBinding string) {
	r.routes[path] = &Route{
		Path:        path,
		Component:   component,
		Title:       title,
		Description: description,
		KeyBinding:  keyBinding,
	}
}
func (r *Router) NavigateTo(path string) error {
	if _, exists := r.routes[path]; !exists {
		return fmt.Errorf("route '%s' not found", path)
	}
	if r.currentRoute != path {
		r.history = append(r.history, r.currentRoute)
	}
	r.currentRoute = path
	return nil
}
func (r *Router) GoBack() error {
	if len(r.history) == 0 {
		return fmt.Errorf("no previous route in history")
	}
	lastRoute := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	r.currentRoute = lastRoute
	return nil
}
func (r *Router) GetCurrentRoute() *Route {
	return r.routes[r.currentRoute]
}
func (r *Router) GetAllRoutes() map[string]*Route {
	return r.routes
}
func (r *Router) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return false, tea.Quit
	case "esc":
		if err := r.GoBack(); err != nil {
			r.NavigateTo("/")
		}
		return true, nil
	}
	availableRoutes := r.getAvailableRoutes()
	for path, route := range availableRoutes {
		if route.KeyBinding == msg.String() {
			r.NavigateTo(path)
			return true, nil
		}
	}
	if currentRoute := r.GetCurrentRoute(); currentRoute != nil {
		return currentRoute.Component.HandleInput(msg)
	}
	return true, nil
}
func (r *Router) getAvailableRoutes() map[string]*Route {
	available := make(map[string]*Route)
	switch r.currentRoute {
	case "/":
		for path, route := range r.routes {
			if path == "/langs" || path == "/config" || path == "/help" {
				available[path] = route
			}
		}
	case "/langs":
		for path, route := range r.routes {
			if strings.HasPrefix(path, "/langs/") && strings.Count(path, "/") == 2 {
				available[path] = route
			}
		}
	case "/langs/golang":
		for path, route := range r.routes {
			if strings.HasPrefix(path, "/langs/golang/") && strings.Count(path, "/") == 3 {
				available[path] = route
			}
		}
	default:
	}
	return available
}
func (r *Router) RenderCurrentPage(width, height int) string {
	currentRoute := r.GetCurrentRoute()
	if currentRoute == nil {
		return r.styles.Error.Render("Error: Route not found")
	}
	return currentRoute.Component.Render(width, height)
}
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
func (r *Router) RenderFooter(width int) string {
	var footerItems []string
	if currentRoute := r.GetCurrentRoute(); currentRoute != nil {
		keyBindings := currentRoute.Component.GetKeyBindings()
		for _, kb := range keyBindings {
			footerItems = append(footerItems, fmt.Sprintf("[%s] %s", kb.Key, kb.Description))
		}
	}
	footerItems = append(footerItems, "[esc] Go back", "[ctrl+c] Exit")
	footerContent := "💡 " + strings.Join(footerItems, " | ")
	return r.styles.Footer.Width(width - 4).Render(footerContent)
}
func (r *Router) RenderStatusBar(width int) string {
	breadcrumb := r.GetBreadcrumb()
	statusContent := fmt.Sprintf("📍 %s", breadcrumb)
	return r.styles.StatusBar.Width(width - 4).Render(statusContent)
}
func (r *Router) GetBreadcrumb() string {
	if r.currentRoute == "/" {
		return "home"
	}
	parts := strings.Split(strings.Trim(r.currentRoute, "/"), "/")
	return strings.Join(parts, " > ")
}
func (r *Router) UpdateTheme(newTheme *theme.Theme) {
	r.theme = newTheme
	r.styles = NewRouterStyles(newTheme)
}
