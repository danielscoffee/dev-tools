package tui
import tea "github.com/charmbracelet/bubbletea"
type KeyBinding struct {
	Key         string
	Description string
	Action      string
}
type PageRenderer interface {
	Render(width, height int) string
	HandleInput(msg tea.KeyMsg) (bool, tea.Cmd)
	GetTitle() string
	GetKeyBindings() []KeyBinding
}
