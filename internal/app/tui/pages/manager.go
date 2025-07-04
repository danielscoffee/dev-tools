package pages

// PageType represents different page types in the TUI
type PageType int

const (
	HomePage PageType = iota
	LanguagesPage
	GolangPage
	JavascriptPage
	PythonPage
	DockerPage
	ToolsPage
	ConfigPage
	HelpPage
)

// PageInfo contains metadata about each page
type PageInfo struct {
	Name        string
	Title       string
	Description string
	KeyBinding  string
	Category    string
}

// PageManager manages all available pages
type PageManager struct {
	pages    map[PageType]PageInfo
	current  PageType
	history  []PageType
	parent   map[PageType]PageType
	children map[PageType][]PageType
}

// NewPageManager creates a new page manager
func NewPageManager() *PageManager {
	pm := &PageManager{
		pages:    make(map[PageType]PageInfo),
		current:  HomePage,
		history:  make([]PageType, 0),
		parent:   make(map[PageType]PageType),
		children: make(map[PageType][]PageType),
	}

	pm.registerPages()
	pm.setupHierarchy()
	return pm
}

// registerPages registers all available pages
func (pm *PageManager) registerPages() {
	pm.pages[HomePage] = PageInfo{
		Name:        "home",
		Title:       "Dev Tools - Home",
		Description: "Main menu and navigation",
		KeyBinding:  "h",
		Category:    "main",
	}

	pm.pages[LanguagesPage] = PageInfo{
		Name:        "languages",
		Title:       "Programming Languages",
		Description: "Tools for different languages",
		KeyBinding:  "l",
		Category:    "main",
	}

	pm.pages[GolangPage] = PageInfo{
		Name:        "golang",
		Title:       "Go/Golang Tools",
		Description: "Tools for Go development",
		KeyBinding:  "g",
		Category:    "language",
	}

	pm.pages[JavascriptPage] = PageInfo{
		Name:        "javascript",
		Title:       "JavaScript/Node.js Tools",
		Description: "Tools for JavaScript development",
		KeyBinding:  "j",
		Category:    "language",
	}

	pm.pages[PythonPage] = PageInfo{
		Name:        "python",
		Title:       "Python Tools",
		Description: "Tools for Python development",
		KeyBinding:  "p",
		Category:    "language",
	}

	pm.pages[DockerPage] = PageInfo{
		Name:        "docker",
		Title:       "Docker Tools",
		Description: "Tools for containerization",
		KeyBinding:  "d",
		Category:    "tool",
	}

	pm.pages[ToolsPage] = PageInfo{
		Name:        "tools",
		Title:       "General Tools",
		Description: "General development utilities",
		KeyBinding:  "t",
		Category:    "main",
	}

	pm.pages[ConfigPage] = PageInfo{
		Name:        "config",
		Title:       "Configuration",
		Description: "Application settings",
		KeyBinding:  "c",
		Category:    "main",
	}

	pm.pages[HelpPage] = PageInfo{
		Name:        "help",
		Title:       "Help & Documentation",
		Description: "Usage instructions and help",
		KeyBinding:  "?",
		Category:    "main",
	}
}

// setupHierarchy configures the hierarchy between pages
func (pm *PageManager) setupHierarchy() {
	// Configure home page children
	pm.children[HomePage] = []PageType{
		LanguagesPage, ToolsPage, ConfigPage, HelpPage,
	}

	// Configure language page children
	pm.children[LanguagesPage] = []PageType{
		GolangPage, JavascriptPage, PythonPage,
	}

	// Configure parents
	pm.parent[LanguagesPage] = HomePage
	pm.parent[ToolsPage] = HomePage
	pm.parent[ConfigPage] = HomePage
	pm.parent[HelpPage] = HomePage

	pm.parent[GolangPage] = LanguagesPage
	pm.parent[JavascriptPage] = LanguagesPage
	pm.parent[PythonPage] = LanguagesPage
}

// NavigateTo navigates to a specific page
func (pm *PageManager) NavigateTo(pageType PageType) bool {
	if _, exists := pm.pages[pageType]; exists {
		pm.history = append(pm.history, pm.current)
		pm.current = pageType
		return true
	}
	return false
}

// GoBack goes back to the previous page
func (pm *PageManager) GoBack() bool {
	if len(pm.history) > 0 {
		pm.current = pm.history[len(pm.history)-1]
		pm.history = pm.history[:len(pm.history)-1]
		return true
	}
	return false
}

// GoToParent navigates to the parent page
func (pm *PageManager) GoToParent() bool {
	if parent, exists := pm.parent[pm.current]; exists {
		pm.history = append(pm.history, pm.current)
		pm.current = parent
		return true
	}
	return false
}

// GetCurrentPage retorna a página atual
func (pm *PageManager) GetCurrentPage() PageInfo {
	return pm.pages[pm.current]
}

// GetCurrentPageType retorna o tipo da página atual
func (pm *PageManager) GetCurrentPageType() PageType {
	return pm.current
}

// GetPage retorna informações de uma página específica
func (pm *PageManager) GetPage(pageType PageType) (PageInfo, bool) {
	page, exists := pm.pages[pageType]
	return page, exists
}

// GetChildren retorna as páginas filhas de uma página
func (pm *PageManager) GetChildren(pageType PageType) []PageType {
	return pm.children[pageType]
}

// GetChildrenInfo retorna informações das páginas filhas
func (pm *PageManager) GetChildrenInfo(pageType PageType) []PageInfo {
	children := pm.children[pageType]
	infos := make([]PageInfo, len(children))

	for i, child := range children {
		infos[i] = pm.pages[child]
	}

	return infos
}

// GetPageByName encontra uma página pelo nome
func (pm *PageManager) GetPageByName(name string) (PageType, PageInfo, bool) {
	for pageType, info := range pm.pages {
		if info.Name == name {
			return pageType, info, true
		}
	}
	return HomePage, PageInfo{}, false
}

// GetPageByKeyBinding encontra uma página pela tecla de atalho
func (pm *PageManager) GetPageByKeyBinding(key string) (PageType, PageInfo, bool) {
	for pageType, info := range pm.pages {
		if info.KeyBinding == key {
			return pageType, info, true
		}
	}
	return HomePage, PageInfo{}, false
}

// GetBreadcrumb retorna o caminho de navegação atual
func (pm *PageManager) GetBreadcrumb() []PageInfo {
	breadcrumb := []PageInfo{}
	current := pm.current

	// Adicionar página atual
	breadcrumb = append([]PageInfo{pm.pages[current]}, breadcrumb...)

	// Adicionar páginas pais
	for {
		if parent, exists := pm.parent[current]; exists {
			breadcrumb = append([]PageInfo{pm.pages[parent]}, breadcrumb...)
			current = parent
		} else {
			break
		}
	}

	return breadcrumb
}

// GetAllPages retorna todas as páginas registradas
func (pm *PageManager) GetAllPages() map[PageType]PageInfo {
	return pm.pages
}

// GetPagesByCategory retorna páginas de uma categoria específica
func (pm *PageManager) GetPagesByCategory(category string) []PageInfo {
	var pages []PageInfo

	for _, info := range pm.pages {
		if info.Category == category {
			pages = append(pages, info)
		}
	}

	return pages
}

// IsValidTransition verifica se uma transição é válida
func (pm *PageManager) IsValidTransition(from, to PageType) bool {
	// Sempre pode voltar para home
	if to == HomePage {
		return true
	}

	// Pode navegar para páginas filhas
	children := pm.children[from]
	for _, child := range children {
		if child == to {
			return true
		}
	}

	// Pode navegar para página pai
	if parent, exists := pm.parent[from]; exists && parent == to {
		return true
	}

	return false
}

// Reset volta para a página inicial e limpa o histórico
func (pm *PageManager) Reset() {
	pm.current = HomePage
	pm.history = make([]PageType, 0)
}
