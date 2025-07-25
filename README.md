# 🛠️ Dev Tools TUI

A comprehensive suite of development tools designed to enhance your developer experience across multiple programming languages and technologies.

## 🚀 Features

- **Router-Based Architecture**: Dynamic page discovery and navigation system
- **Multi-Language Support**: Tools for Go, JavaScript, Python, and more
- **Project Scaffolding**: Integrated go-blueprint for Go project creation
- **Interactive TUI**: Beautiful terminal user interface with intuitive navigation
- **Modular Design**: Easy to extend with new pages and functionality
- **Configuration Management**: Customizable settings and preferences

## 📦 Installation

```bash
# Clone the repository
git clone https://github.com/danielscoffee/dev-tools.git
cd dev-tools

# Build the application
make build

# Or build directly
go build -o bin/dev-tools cmd/main.go
```

## 🎯 Usage

### Launch TUI
```bash
# Start the interactive TUI
./bin/dev-tools tui

# Or using make
make tui
```

### CLI Commands
```bash
# Show help
./bin/dev-tools --help

# Start TUI mode
./bin/dev-tools tui
```

## 🏗️ Router System Architecture

The TUI uses a sophisticated router system that dynamically discovers and processes page components from the `pages/` directory structure.

### Page Structure
```
internal/app/tui/pages/
├── home/
│   └── view.go              # Home page component
├── langs/
│   ├── view.go              # Languages overview page
│   └── golang/
│       └── view.go          # Go-specific tools page
├── config/
│   └── view.go              # Configuration page
└── help/
    └── view.go              # Help and documentation
```

### Adding New Pages

To add a new page, create a component that implements the `PageRenderer` interface:

```go
package mypage

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/danielscoffee/dev-tools/internal/app/tui/types"
)

type Page struct {
    // Your page state
}

func NewPage() *Page {
    return &Page{}
}

// Implement PageRenderer interface
func (p *Page) Render(width, height int) string {
    // Return your page content
}

func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
    // Handle page-specific input
}

func (p *Page) GetTitle() string {
    // Return page title
}

func (p *Page) GetKeyBindings() []types.KeyBinding {
    // Return page key bindings
}
```

Then register it in the router:

```go
router.RegisterRoute("/mypage", mypage.NewPage(), "My Page", "Description", "key")
```

## 🔧 Development

### Build and Test
```bash
# Build the application
make build

# Run tests
make test

# Run TUI tests
make tui-test

# Test router system
make router-test

# Run demo
make demo
```

### Router Demo
```bash
# Run the router system demonstration
make router-demo
```

## 🎨 Navigation

### Global Keys
- `esc` - Go back to previous page
- `ctrl+c` - Exit application
- `q` - Quit (context-dependent)

### Home Page
- `l` - Navigate to Languages
- `c` - Navigate to Configuration
- `?` - Navigate to Help

### Languages Page
- `g` - Go/Golang tools
- `j` - JavaScript tools
- `p` - Python tools

### Go/Golang Page
- `b` - Blueprint project creation
- `m` - Module management
- `t` - Test runner
- `f` - Code formatter

## 🔄 Router System Details

### Key Components

1. **Router**: Core routing engine that manages navigation
2. **PageRenderer Interface**: Contract that all page components must implement
3. **Route Registration**: Dynamic route registration system
4. **Breadcrumb Navigation**: Hierarchical navigation with history
5. **Key Binding System**: Dynamic key binding discovery and display

### Router Features

- **Dynamic Page Discovery**: Automatically processes page components
- **Hierarchical Navigation**: Support for nested routes (e.g., `/langs/golang`)
- **History Management**: Back navigation with route history
- **Key Binding Management**: Automatic footer generation with available keys
- **Error Handling**: Graceful handling of navigation errors

### Example Router Usage

```go
// Create router
router := tui.NewRouter()

// Register routes
router.RegisterRoute("/", home.NewPage(), "Home", "Main menu", "h")
router.RegisterRoute("/langs", langs.NewPage(), "Languages", "Programming languages", "l")
router.RegisterRoute("/langs/golang", golang.NewPage(), "Go Tools", "Go development", "g")

// Navigate
router.NavigateTo("/langs/golang")

// Go back
router.GoBack()

// Get current page
currentPage := router.GetCurrentRoute()
```

## 🧩 Go-Blueprint Integration

The TUI includes integrated support for go-blueprint project scaffolding:

- **REST API Projects**: Create REST APIs with various frameworks
- **CLI Applications**: Generate CLI applications with Cobra
- **Web Applications**: Build web apps with Fiber, Gin, or other frameworks
- **Database Integration**: Support for various database configurations

## 🔧 Configuration

The configuration system allows customization of:

- **Themes**: Application colors and styling
- **Key Bindings**: Custom keyboard shortcuts
- **Default Paths**: Project creation directories
- **Tool Preferences**: Default tools and settings

## 📚 Contributing

1. Fork the repository
2. Create a feature branch
3. Add your page component following the `PageRenderer` interface
4. Register your route in the main router
5. Add tests for your component
6. Submit a pull request

## 🎯 Future Enhancements

- **Plugin System**: Dynamic plugin loading
- **Themes**: Multiple color themes
- **Persistent Settings**: Configuration file support
- **More Languages**: Additional language support
- **Tool Integration**: More development tools integration

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Acknowledgments

- Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) for the TUI framework
- Styled with [Lipgloss](https://github.com/charmbracelet/lipgloss) for beautiful terminal UI
- Integrated with [go-blueprint](https://github.com/Melkeydev/go-blueprint) for project scaffolding
