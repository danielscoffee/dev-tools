package pages

// pages.go - Main entry point for the page system

var (
	// GlobalPageManager is the global instance of the page manager
	GlobalPageManager *PageManager
)

// InitPages initializes the page system
func InitPages() {
	GlobalPageManager = NewPageManager()
}

// GetManager returns the global page manager
func GetManager() *PageManager {
	if GlobalPageManager == nil {
		InitPages()
	}
	return GlobalPageManager
}

// NavigateToPage navigates to a page using the global manager
func NavigateToPage(pageType PageType) bool {
	return GetManager().NavigateTo(pageType)
}

// GoBackPage goes back one page using the global manager
func GoBackPage() bool {
	return GetManager().GoBack()
}

// GetCurrentPageInfo returns information about the current page
func GetCurrentPageInfo() PageInfo {
	return GetManager().GetCurrentPage()
}

// GetAvailablePages returns all available pages organized by category
func GetAvailablePages() map[string][]PageInfo {
	manager := GetManager()
	categories := make(map[string][]PageInfo)

	for _, page := range manager.GetAllPages() {
		categories[page.Category] = append(categories[page.Category], page)
	}

	return categories
}

// GetMainMenuPages returns main menu pages
func GetMainMenuPages() []PageInfo {
	return GetManager().GetPagesByCategory("main")
}

// GetLanguagePages retorna páginas de linguagens
func GetLanguagePages() []PageInfo {
	return GetManager().GetPagesByCategory("language")
}

// HandleKeyBinding processa teclas de atalho para navegação
func HandleKeyBinding(key string) bool {
	manager := GetManager()

	// Teclas especiais
	switch key {
	case "esc", "q":
		return manager.GoBack()
	case "ctrl+c":
		return false // Sinaliza para sair da aplicação
	}

	// Buscar página por tecla de atalho
	if pageType, _, found := manager.GetPageByKeyBinding(key); found {
		return manager.NavigateTo(pageType)
	}

	return false
}
