package tui

import tea "github.com/charmbracelet/bubbletea"

// KeyBinding represents a key binding for a page
type KeyBinding struct {
	Key         string
	Description string
	Action      string
}

// PageRenderer interface that all page components must implement
type PageRenderer interface {
	Render(width, height int) string
	HandleInput(msg tea.KeyMsg) (bool, tea.Cmd)
	GetTitle() string
	GetKeyBindings() []KeyBinding
}
