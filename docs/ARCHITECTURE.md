# Router System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                        Dev Tools TUI                            │
├─────────────────────────────────────────────────────────────────┤
│  ┌─────────────────┐    ┌─────────────────┐    ┌──────────────┐ │
│  │     Router      │    │   PageRenderer  │    │    Types     │ │
│  │                 │    │   Interface     │    │              │ │
│  │ • Route Mgmt    │◄──►│ • Render()      │    │ • KeyBinding │ │
│  │ • Navigation    │    │ • HandleInput() │    │ • Route      │ │
│  │ • History       │    │ • GetTitle()    │    │              │ │
│  │ • Breadcrumb    │    │ • GetKeyBinds() │    │              │ │
│  └─────────────────┘    └─────────────────┘    └──────────────┘ │
│           │                       │                             │
│           ▼                       ▼                             │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │                    Page Components                          │ │
│  │  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌───────┐ │ │
│  │  │  Home   │ │  Langs  │ │ Config  │ │  Help   │ │  ...  │ │ │
│  │  │ Page    │ │ Page    │ │ Page    │ │ Page    │ │ Pages │ │ │
│  │  └─────────┘ └─────────┘ └─────────┘ └─────────┘ └───────┘ │ │
│  │       │           │           │           │           │     │ │
│  │       ▼           ▼           ▼           ▼           ▼     │ │
│  │  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌───────┐ │ │
│  │  │  /      │ │ /langs  │ │/config  │ │ /help   │ │  ...  │ │ │
│  │  │ Route   │ │ Route   │ │ Route   │ │ Route   │ │ Routes│ │ │
│  │  └─────────┘ └─────────┘ └─────────┘ └─────────┘ └───────┘ │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘

Router Flow:
1. User Input → Router.HandleInput()
2. Router checks key bindings for navigation
3. Router.NavigateTo() switches to new route
4. Router.RenderCurrentPage() calls PageRenderer.Render()
5. Page handles specific input via HandleInput()
6. Router maintains history and breadcrumb navigation

Page Discovery:
• Each page implements PageRenderer interface
• Pages are registered with routes in main application
• Router dynamically processes page components
• Key bindings are discovered from page components
• Footer is automatically generated with available keys

Directory Structure:
internal/app/tui/pages/
├── home/view.go         → Route: "/"
├── langs/view.go        → Route: "/langs"
├── langs/golang/view.go → Route: "/langs/golang"
├── config/view.go       → Route: "/config"
└── help/view.go         → Route: "/help"
```
