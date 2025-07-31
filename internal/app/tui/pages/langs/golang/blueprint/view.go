package blueprint

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/danielscoffee/dev-tools/internal/app/tui/types"
	"github.com/danielscoffee/dev-tools/internal/pkg/backend/golang"
)

// FormStep represents different steps in the blueprint creation form
type FormStep int

const (
	StepProjectName FormStep = iota
	StepFramework
	StepDatabase
	StepFeatures
	StepGitOption
	StepConfirm
	StepCreating
	StepComplete
	StepError
)

// Page represents the blueprint page component
type Page struct {
	styles      *PageStyles
	currentStep FormStep
	blueprint   *golang.Blueprint

	// Form data
	projectName string
	framework   string
	database    string
	features    []string
	gitOption   string
	advanced    bool

	// UI state
	input             string
	cursor            int
	selectedIndex     int
	multiSelectStates map[string]bool
	error             string
	creationOutput    []string
	isCreating        bool

	// Available options
	frameworks []string
	databases  []string
	allFeatures []string
	gitOptions  []string
}

// PageStyles defines styles for the blueprint page
type PageStyles struct {
	Title         lipgloss.Style
	Description   lipgloss.Style
	FormLabel     lipgloss.Style
	Input         lipgloss.Style
	InputFocus    lipgloss.Style
	Option        lipgloss.Style
	OptionFocus   lipgloss.Style
	Selected      lipgloss.Style
	Checkbox      lipgloss.Style
	CheckboxFocus lipgloss.Style
	Button        lipgloss.Style
	ButtonFocus   lipgloss.Style
	Error         lipgloss.Style
	Success       lipgloss.Style
	Output        lipgloss.Style
}

// NewPage creates a new blueprint page instance
func NewPage() *Page {
	bp := &golang.Blueprint{}
	
	return &Page{
		styles:            NewPageStyles(),
		currentStep:       StepProjectName,
		blueprint:         bp,
		multiSelectStates: make(map[string]bool),
		
		// Initialize options
		frameworks:  bp.GetSupportedFrameworks(),
		databases:   bp.GetSupportedDrivers(),
		allFeatures: bp.GetSupportedFeatures(),
		gitOptions:  []string{"init", "commit", "skip"},
		
		// Default values
		gitOption: "commit",
		advanced:  true,
	}
}

// NewPageStyles creates default styles for the blueprint page
func NewPageStyles() *PageStyles {
	return &PageStyles{
		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00D7FF")).
			MarginBottom(1),

		Description: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CCCCCC")).
			MarginBottom(1),

		FormLabel: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#98FB98")).
			MarginBottom(1),

		Input: lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color("#1A1A1A")).
			Foreground(lipgloss.Color("#CCCCCC")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#383838")),

		InputFocus: lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color("#1A1A1A")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00D7FF")),

		Option: lipgloss.NewStyle().
			Padding(0, 2).
			Foreground(lipgloss.Color("#CCCCCC")),

		OptionFocus: lipgloss.NewStyle().
			Padding(0, 2).
			Background(lipgloss.Color("#383838")).
			Foreground(lipgloss.Color("#FFFFFF")),

		Selected: lipgloss.NewStyle().
			Padding(0, 2).
			Background(lipgloss.Color("#00D7FF")).
			Foreground(lipgloss.Color("#000000")),

		Checkbox: lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(lipgloss.Color("#CCCCCC")),

		CheckboxFocus: lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color("#383838")).
			Foreground(lipgloss.Color("#FFFFFF")),

		Button: lipgloss.NewStyle().
			Padding(0, 2).
			Background(lipgloss.Color("#383838")).
			Foreground(lipgloss.Color("#CCCCCC")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#383838")),

		ButtonFocus: lipgloss.NewStyle().
			Padding(0, 2).
			Background(lipgloss.Color("#00D7FF")).
			Foreground(lipgloss.Color("#000000")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#00D7FF")),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Background(lipgloss.Color("#2D1B1B")).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF6B6B")),

		Success: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#98FB98")).
			Background(lipgloss.Color("#1B2D1B")).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#98FB98")),

		Output: lipgloss.NewStyle().
			Padding(1, 2).
			Background(lipgloss.Color("#1A1A1A")).
			Foreground(lipgloss.Color("#CCCCCC")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#383838")),
	}
}

// Render renders the blueprint page
func (p *Page) Render(width, height int) string {
	var content []string

	content = append(content, p.styles.Title.Render("🏗️  Go Blueprint Project Creator"))
	content = append(content, "")

	switch p.currentStep {
	case StepProjectName:
		content = append(content, p.renderProjectNameStep()...)
	case StepFramework:
		content = append(content, p.renderFrameworkStep()...)
	case StepDatabase:
		content = append(content, p.renderDatabaseStep()...)
	case StepFeatures:
		content = append(content, p.renderFeaturesStep()...)
	case StepGitOption:
		content = append(content, p.renderGitOptionStep()...)
	case StepConfirm:
		content = append(content, p.renderConfirmStep()...)
	case StepCreating:
		content = append(content, p.renderCreatingStep()...)
	case StepComplete:
		content = append(content, p.renderCompleteStep()...)
	case StepError:
		content = append(content, p.renderErrorStep()...)
	}

	// Add navigation help
	content = append(content, "")
	content = append(content, p.styles.Description.Render("Navigation: [enter] next/confirm • [esc] back • [ctrl+c] exit"))

	return lipgloss.JoinVertical(lipgloss.Left, content...)
}

func (p *Page) renderProjectNameStep() []string {
	var content []string
	
	content = append(content, p.styles.FormLabel.Render("📝 Project Name:"))
	content = append(content, p.styles.Description.Render("Enter a name for your new Go project"))
	content = append(content, "")
	
	inputStyle := p.styles.Input
	if p.currentStep == StepProjectName {
		inputStyle = p.styles.InputFocus
	}
	
	inputText := p.input
	if len(inputText) == 0 {
		inputText = "my-awesome-project"
	}
	
	content = append(content, inputStyle.Render(inputText))
	content = append(content, "")
	content = append(content, p.styles.Description.Render("Example: my-api, web-app, microservice"))
	
	return content
}

func (p *Page) renderFrameworkStep() []string {
	var content []string
	
	content = append(content, p.styles.FormLabel.Render("🚀 Web Framework:"))
	content = append(content, p.styles.Description.Render("Choose a web framework for your project"))
	content = append(content, "")
	
	for i, framework := range p.frameworks {
		style := p.styles.Option
		prefix := "  "
		
		if i == p.selectedIndex {
			style = p.styles.OptionFocus
			prefix = "▶ "
		}
		
		if framework == p.framework {
			style = p.styles.Selected
			prefix = "✓ "
		}
		
		content = append(content, style.Render(prefix+framework))
	}
	
	return content
}

func (p *Page) renderDatabaseStep() []string {
	var content []string
	
	content = append(content, p.styles.FormLabel.Render("🗄️  Database Driver:"))
	content = append(content, p.styles.Description.Render("Choose a database driver (optional)"))
	content = append(content, "")
	
	// Add "none" option
	options := append([]string{"none"}, p.databases...)
	
	for i, database := range options {
		style := p.styles.Option
		prefix := "  "
		
		if i == p.selectedIndex {
			style = p.styles.OptionFocus
			prefix = "▶ "
		}
		
		if database == p.database {
			style = p.styles.Selected
			prefix = "✓ "
		}
		
		content = append(content, style.Render(prefix+database))
	}
	
	return content
}

func (p *Page) renderFeaturesStep() []string {
	var content []string
	
	content = append(content, p.styles.FormLabel.Render("🎨 Additional Features:"))
	content = append(content, p.styles.Description.Render("Select features to include (space to toggle, enter to continue)"))
	content = append(content, "")
	
	for i, feature := range p.allFeatures {
		style := p.styles.Checkbox
		prefix := "  [ ] "
		
		if i == p.selectedIndex {
			style = p.styles.CheckboxFocus
		}
		
		if p.multiSelectStates[feature] {
			prefix = "  [✓] "
		}
		
		if i == p.selectedIndex {
			prefix = "▶" + prefix[1:]
		}
		
		content = append(content, style.Render(prefix+feature))
	}
	
	content = append(content, "")
	selectedFeatures := []string{}
	for feature, selected := range p.multiSelectStates {
		if selected {
			selectedFeatures = append(selectedFeatures, feature)
		}
	}
	
	if len(selectedFeatures) > 0 {
		content = append(content, p.styles.Description.Render("Selected: "+strings.Join(selectedFeatures, ", ")))
	}
	
	return content
}

func (p *Page) renderGitOptionStep() []string {
	var content []string
	
	content = append(content, p.styles.FormLabel.Render("📚 Git Initialization:"))
	content = append(content, p.styles.Description.Render("Choose how to handle Git repository"))
	content = append(content, "")
	
	for i, option := range p.gitOptions {
		style := p.styles.Option
		prefix := "  "
		description := ""
		
		switch option {
		case "init":
			description = " - Initialize git repository"
		case "commit":
			description = " - Initialize and create initial commit"
		case "skip":
			description = " - Skip git initialization"
		}
		
		if i == p.selectedIndex {
			style = p.styles.OptionFocus
			prefix = "▶ "
		}
		
		if option == p.gitOption {
			style = p.styles.Selected
			prefix = "✓ "
		}
		
		content = append(content, style.Render(prefix+option+description))
	}
	
	return content
}

func (p *Page) renderConfirmStep() []string {
	var content []string
	
	content = append(content, p.styles.FormLabel.Render("✅ Confirm Project Creation"))
	content = append(content, "")
	
	content = append(content, p.styles.Description.Render("Project Name: "+p.projectName))
	content = append(content, p.styles.Description.Render("Framework: "+p.framework))
	
	if p.database != "" && p.database != "none" {
		content = append(content, p.styles.Description.Render("Database: "+p.database))
	}
	
	if len(p.features) > 0 {
		content = append(content, p.styles.Description.Render("Features: "+strings.Join(p.features, ", ")))
	}
	
	content = append(content, p.styles.Description.Render("Git: "+p.gitOption))
	content = append(content, "")
	
	// Show the command that will be executed
	command := p.blueprint.GetCommandString(p.projectName, p.framework, p.database, p.gitOption, p.features, p.advanced)
	content = append(content, p.styles.FormLabel.Render("Command to execute:"))
	content = append(content, p.styles.Output.Render(command))
	content = append(content, "")
	
	// Buttons
	buttons := []string{"Create Project", "Cancel"}
	for i, button := range buttons {
		style := p.styles.Button
		if i == p.selectedIndex {
			style = p.styles.ButtonFocus
		}
		content = append(content, style.Render("  "+button+"  "))
	}
	
	return content
}

func (p *Page) renderCreatingStep() []string {
	var content []string
	
	content = append(content, p.styles.FormLabel.Render("🚧 Creating Project..."))
	content = append(content, "")
	content = append(content, p.styles.Description.Render("Please wait while your project is being created."))
	content = append(content, "")
	
	// Show creation output
	if len(p.creationOutput) > 0 {
		content = append(content, p.styles.Output.Render(strings.Join(p.creationOutput, "\n")))
	}
	
	return content
}

func (p *Page) renderCompleteStep() []string {
	var content []string
	
	content = append(content, p.styles.Success.Render("🎉 Project Created Successfully!"))
	content = append(content, "")
	content = append(content, p.styles.Description.Render("Your project '"+p.projectName+"' has been created."))
	content = append(content, "")
	content = append(content, p.styles.Description.Render("Next steps:"))
	content = append(content, p.styles.Description.Render("  1. cd "+p.projectName))
	content = append(content, p.styles.Description.Render("  2. go mod tidy"))
	content = append(content, p.styles.Description.Render("  3. go run main.go"))
	content = append(content, "")
	
	// Show button to create another project
	style := p.styles.ButtonFocus
	content = append(content, style.Render("  Create Another Project  "))
	
	return content
}

func (p *Page) renderErrorStep() []string {
	var content []string
	
	content = append(content, p.styles.Error.Render("❌ Error Creating Project"))
	content = append(content, "")
	content = append(content, p.styles.Description.Render("An error occurred while creating your project:"))
	content = append(content, "")
	content = append(content, p.styles.Error.Render(p.error))
	content = append(content, "")
	
	// Show button to try again
	style := p.styles.ButtonFocus
	content = append(content, style.Render("  Try Again  "))
	
	return content
}

// CreateProjectCmd represents a command to create a project
type CreateProjectCmd struct {
	projectName string
	framework   string
	database    string
	features    []string
	gitOption   string
	advanced    bool
	blueprint   *golang.Blueprint
}

// ProjectCreatedMsg represents a successful project creation
type ProjectCreatedMsg struct{}

// ProjectErrorMsg represents a project creation error
type ProjectErrorMsg struct {
	error string
}

func (c CreateProjectCmd) Execute() tea.Msg {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Check if go-blueprint is installed
	if !c.blueprint.IsInstalled() {
		if err := c.blueprint.InstallCLI(ctx); err != nil {
			return ProjectErrorMsg{error: fmt.Sprintf("Failed to install go-blueprint: %v", err)}
		}
	}

	// Create the project
	var err error
	if c.advanced && len(c.features) > 0 {
		err = c.blueprint.CreateAdvanced(ctx, c.projectName, c.framework, c.database, c.features)
	} else {
		err = c.blueprint.CreateWithGit(ctx, c.projectName, c.framework, c.database, c.gitOption)
	}

	if err != nil {
		return ProjectErrorMsg{error: fmt.Sprintf("Failed to create project: %v", err)}
	}

	return ProjectCreatedMsg{}
}

// HandleInput handles input for the blueprint page
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	// Handle escape key for going back to previous step or route
	if msg.String() == "esc" {
		switch p.currentStep {
		case StepProjectName:
			// At first step, let router handle escape to go back to parent route
			return false, nil
		case StepFramework:
			p.currentStep = StepProjectName
			p.selectedIndex = 0
			return true, nil
		case StepDatabase:
			p.currentStep = StepFramework
			p.selectedIndex = 0
			return true, nil
		case StepFeatures:
			p.currentStep = StepDatabase
			p.selectedIndex = 0
			return true, nil
		case StepGitOption:
			p.currentStep = StepFeatures
			p.selectedIndex = 0
			return true, nil
		case StepConfirm:
			p.currentStep = StepGitOption
			p.selectedIndex = 1 // Default to "commit"
			return true, nil
		case StepCreating:
			// Can't go back while creating
			return true, nil
		case StepComplete:
			// Reset and go back to first step
			p.reset()
			return true, nil
		case StepError:
			// Go back to confirm step to try again
			p.currentStep = StepConfirm
			p.selectedIndex = 0
			return true, nil
		}
		return true, nil
	}

	switch p.currentStep {
	case StepProjectName:
		return p.handleProjectNameInput(msg)
	case StepFramework:
		return p.handleFrameworkInput(msg)
	case StepDatabase:
		return p.handleDatabaseInput(msg)
	case StepFeatures:
		return p.handleFeaturesInput(msg)
	case StepGitOption:
		return p.handleGitOptionInput(msg)
	case StepConfirm:
		return p.handleConfirmInput(msg)
	case StepCreating:
		return p.handleCreatingInput(msg)
	case StepComplete:
		return p.handleCompleteInput(msg)
	case StepError:
		return p.handleErrorInput(msg)
	}
	return true, nil
}

func (p *Page) handleProjectNameInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "enter":
		if p.input != "" {
			p.projectName = p.input
			p.currentStep = StepFramework
			p.selectedIndex = 0
		}
		return true, nil
	case "backspace":
		if len(p.input) > 0 {
			p.input = p.input[:len(p.input)-1]
		}
		return true, nil
	default:
		if len(msg.String()) == 1 {
			p.input += msg.String()
		}
		return true, nil
	}
}

func (p *Page) handleFrameworkInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if p.selectedIndex > 0 {
			p.selectedIndex--
		}
		return true, nil
	case "down", "j":
		if p.selectedIndex < len(p.frameworks)-1 {
			p.selectedIndex++
		}
		return true, nil
	case "enter":
		p.framework = p.frameworks[p.selectedIndex]
		p.currentStep = StepDatabase
		p.selectedIndex = 0
		return true, nil
	}
	return true, nil
}

func (p *Page) handleDatabaseInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	options := append([]string{"none"}, p.databases...)
	
	switch msg.String() {
	case "up", "k":
		if p.selectedIndex > 0 {
			p.selectedIndex--
		}
		return true, nil
	case "down", "j":
		if p.selectedIndex < len(options)-1 {
			p.selectedIndex++
		}
		return true, nil
	case "enter":
		p.database = options[p.selectedIndex]
		p.currentStep = StepFeatures
		p.selectedIndex = 0
		return true, nil
	}
	return true, nil
}

func (p *Page) handleFeaturesInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if p.selectedIndex > 0 {
			p.selectedIndex--
		}
		return true, nil
	case "down", "j":
		if p.selectedIndex < len(p.allFeatures)-1 {
			p.selectedIndex++
		}
		return true, nil
	case " ":
		feature := p.allFeatures[p.selectedIndex]
		p.multiSelectStates[feature] = !p.multiSelectStates[feature]
		return true, nil
	case "enter":
		// Collect selected features
		p.features = []string{}
		for feature, selected := range p.multiSelectStates {
			if selected {
				p.features = append(p.features, feature)
			}
		}
		p.currentStep = StepGitOption
		p.selectedIndex = 1 // Default to "commit"
		return true, nil
	}
	return true, nil
}

func (p *Page) handleGitOptionInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if p.selectedIndex > 0 {
			p.selectedIndex--
		}
		return true, nil
	case "down", "j":
		if p.selectedIndex < len(p.gitOptions)-1 {
			p.selectedIndex++
		}
		return true, nil
	case "enter":
		p.gitOption = p.gitOptions[p.selectedIndex]
		p.currentStep = StepConfirm
		p.selectedIndex = 0
		return true, nil
	}
	return true, nil
}

func (p *Page) handleConfirmInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if p.selectedIndex > 0 {
			p.selectedIndex--
		}
		return true, nil
	case "down", "j":
		if p.selectedIndex < 1 {
			p.selectedIndex++
		}
		return true, nil
	case "enter":
		if p.selectedIndex == 0 { // Create Project
			p.currentStep = StepCreating
			p.isCreating = true
			return true, CreateProjectCmd{
				projectName: p.projectName,
				framework:   p.framework,
				database:    p.database,
				features:    p.features,
				gitOption:   p.gitOption,
				advanced:    p.advanced,
				blueprint:   p.blueprint,
			}.Execute
		} else { // Cancel
			return false, nil // Let router handle navigation back
		}
		return true, nil
	}
	return true, nil
}

func (p *Page) handleCreatingInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	// Don't handle input while creating
	return true, nil
}

func (p *Page) handleCompleteInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "enter":
		// Reset form for new project
		p.reset()
		return true, nil
	}
	return true, nil
}

func (p *Page) handleErrorInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "enter":
		// Go back to confirm step to try again
		p.currentStep = StepConfirm
		p.selectedIndex = 0
		return true, nil
	}
	return true, nil
}

// Update handles blueprint-specific messages
func (p *Page) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case ProjectCreatedMsg:
		p.currentStep = StepComplete
		p.isCreating = false
		return nil
	case ProjectErrorMsg:
		p.currentStep = StepError
		p.error = msg.error
		p.isCreating = false
		return nil
	}
	return nil
}

func (p *Page) reset() {
	p.currentStep = StepProjectName
	p.projectName = ""
	p.framework = ""
	p.database = ""
	p.features = []string{}
	p.gitOption = "commit"
	p.input = ""
	p.selectedIndex = 0
	p.multiSelectStates = make(map[string]bool)
	p.error = ""
	p.creationOutput = []string{}
	p.isCreating = false
}

// GetTitle returns the page title
func (p *Page) GetTitle() string {
	return "Go Blueprint Creator"
}

// GetKeyBindings returns the key bindings for this page
func (p *Page) GetKeyBindings() []types.KeyBinding {
	bindings := []types.KeyBinding{
		{Key: "enter", Description: "Next/Confirm", Action: "next"},
		{Key: "↑/↓", Description: "Navigate", Action: "navigate"},
	}
	
	if p.currentStep == StepFeatures {
		bindings = append(bindings, types.KeyBinding{Key: "space", Description: "Toggle", Action: "toggle"})
	}
	
	// Add escape key description based on current step
	if p.currentStep == StepProjectName {
		bindings = append(bindings, types.KeyBinding{Key: "esc", Description: "Back to Go Tools", Action: "back"})
	} else if p.currentStep != StepCreating {
		bindings = append(bindings, types.KeyBinding{Key: "esc", Description: "Previous Step", Action: "back"})
	}
	
	return bindings
}
