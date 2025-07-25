# Adding New Pages to the Router System

This guide shows how to add new pages to the Dev Tools TUI router system.

## Step 1: Create a Page Component

Create a new directory and `view.go` file for your page:

```bash
mkdir -p internal/app/tui/pages/mytools
```

```go
// internal/app/tui/pages/mytools/view.go
package mytools

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    
    "github.com/danielscoffee/dev-tools/internal/app/tui/types"
)

// Page represents the mytools page component
type Page struct {
    styles *PageStyles
}

// PageStyles defines styles for the page
type PageStyles struct {
    Title       lipgloss.Style
    Description lipgloss.Style
    MenuItem    lipgloss.Style
    KeyBinding  lipgloss.Style
}

// NewPage creates a new page instance
func NewPage() *Page {
    return &Page{
        styles: NewPageStyles(),
    }
}

// NewPageStyles creates default styles
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

// Render renders the page content
func (p *Page) Render(width, height int) string {
    var items []string
    
    items = append(items, "🔧 My Tools")
    items = append(items, "")
    items = append(items, p.styles.Description.Render("Custom development tools and utilities."))
    items = append(items, "")
    
    // Add your menu items
    keyStyle := p.styles.KeyBinding.Render("[1]")
    item := lipgloss.JoinHorizontal(
        lipgloss.Left,
        keyStyle,
        " Tool One - ",
        p.styles.Description.Render("Description of tool one"),
    )
    items = append(items, p.styles.MenuItem.Render(item))
    
    keyStyle = p.styles.KeyBinding.Render("[2]")
    item = lipgloss.JoinHorizontal(
        lipgloss.Left,
        keyStyle,
        " Tool Two - ",
        p.styles.Description.Render("Description of tool two"),
    )
    items = append(items, p.styles.MenuItem.Render(item))

    return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// HandleInput handles input for the page
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
    switch msg.String() {
    case "1":
        // Handle tool one
        return true, nil
    case "2":
        // Handle tool two
        return true, nil
    }
    return true, nil
}

// GetTitle returns the page title
func (p *Page) GetTitle() string {
    return "My Tools"
}

// GetKeyBindings returns the key bindings for this page
func (p *Page) GetKeyBindings() []types.KeyBinding {
    return []types.KeyBinding{
        {Key: "1", Description: "Tool One", Action: "run_tool_one"},
        {Key: "2", Description: "Tool Two", Action: "run_tool_two"},
    }
}
```

## Step 2: Register the Route

Add the import and route registration in `internal/app/tui/view.go`:

```go
import (
    // ... existing imports ...
    "github.com/danielscoffee/dev-tools/internal/app/tui/pages/mytools"
)

// In NewModel() function:
func NewModel() *Model {
    router := NewRouter()
    
    // ... existing routes ...
    router.RegisterRoute("/mytools", mytools.NewPage(), "My Tools", "Custom development tools", "m")
    
    return &Model{
        router: router,
        styles: NewAppStyles(),
    }
}
```

## Step 3: Add Navigation from Parent Page

Update a parent page (e.g., home page) to include navigation to your new page:

```go
// In home/view.go, add to the Render() method:
keyStyle = p.styles.KeyBinding.Render("[m]")
item = lipgloss.JoinHorizontal(
    lipgloss.Left,
    keyStyle,
    " My Tools - ",
    p.styles.Description.Render("Custom development tools"),
)
items = append(items, p.styles.MenuItem.Render(item))

// And in GetKeyBindings():
func (p *Page) GetKeyBindings() []types.KeyBinding {
    return []types.KeyBinding{
        // ... existing bindings ...
        {Key: "m", Description: "My Tools", Action: "navigate_mytools"},
    }
}
```

## Step 4: Test Your New Page

Build and test:

```bash
make build
./bin/dev-tools tui
```

Navigate to your new page using the assigned key binding!

## Advanced Examples

### Nested Routes

Create nested routes for sub-pages:

```go
// Register nested routes
router.RegisterRoute("/mytools", mytools.NewPage(), "My Tools", "Custom tools", "m")
router.RegisterRoute("/mytools/advanced", mytools.NewAdvancedPage(), "Advanced Tools", "Advanced utilities", "a")
```

### Dynamic Content

Add dynamic content based on user actions:

```go
type Page struct {
    styles *PageStyles
    selectedTool string
    tools []Tool
}

func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
    switch msg.String() {
    case "1":
        p.selectedTool = "tool1"
        return true, nil
    case "2":
        p.selectedTool = "tool2"
        return true, nil
    }
    return true, nil
}
```

### Custom Styling

Create custom styles for your page:

```go
func NewPageStyles() *PageStyles {
    return &PageStyles{
        Title: lipgloss.NewStyle().
            Bold(true).
            Foreground(lipgloss.Color("#FF6B6B")).  // Custom color
            Background(lipgloss.Color("#2D1B1B")).
            Padding(0, 1),
        // ... other custom styles
    }
}
```

## Tips

1. **Keep it Simple**: Start with basic functionality and add complexity gradually
2. **Consistent Styling**: Use similar styling patterns as existing pages
3. **Clear Key Bindings**: Use intuitive key combinations
4. **Error Handling**: Add proper error handling for user actions
5. **Documentation**: Update README.md with your new page functionality

The router system automatically handles navigation, breadcrumbs, and key binding display, so you can focus on your page's specific functionality!
