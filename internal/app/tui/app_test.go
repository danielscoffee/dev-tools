package tui

import (
	"testing"

	"github.com/danielscoffee/dev-tools/internal/app/tui/pages"
)

func TestNewModel(t *testing.T) {
	model := NewModel()

	if model == nil {
		t.Fatal("Expected model to be created, got nil")
	}

	if model.integration == nil {
		t.Fatal("Expected integration to be initialized")
	}

	if model.styles == nil {
		t.Fatal("Expected styles to be initialized")
	}
}

func TestPageSystemIntegration(t *testing.T) {
	// Initialize the page system
	pages.InitPages()

	manager := pages.GetManager()
	if manager == nil {
		t.Fatal("Expected page manager to be initialized")
	}

	// Test navigation
	currentPage := manager.GetCurrentPage()
	if currentPage.Name != "home" {
		t.Errorf("Expected initial page to be 'home', got %s", currentPage.Name)
	}

	// Test page navigation
	success := manager.NavigateTo(pages.LanguagesPage)
	if !success {
		t.Error("Expected navigation to LanguagesPage to succeed")
	}

	currentPage = manager.GetCurrentPage()
	if currentPage.Name != "languages" {
		t.Errorf("Expected current page to be 'languages', got %s", currentPage.Name)
	}
}

func TestTUIIntegration(t *testing.T) {
	integration := pages.NewTUIIntegration()

	if integration == nil {
		t.Fatal("Expected TUI integration to be created")
	}

	// Test page rendering
	content := integration.RenderCurrentPage()
	if len(content) == 0 {
		t.Error("Expected rendered content to have length > 0")
	}

	// Test footer
	footer := integration.GetFooter()
	if len(footer) == 0 {
		t.Error("Expected footer to have content")
	}
}
