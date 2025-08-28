// Package blueprint
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
type Page struct {
	styles      *PageStyles
	currentStep FormStep
	blueprint   *golang.Blueprint
	projectName string
	framework   string
	database    string
	features    []string
	gitOption   string
	input             string
	cursor            int
	selectedIndex     int
	multiSelectStates map[string]bool
	error             string
	creationOutput    []string
	isCreating        bool
	frameworks  []string
	databases   []string
	allFeatures []string
	gitOptions  []string
}
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
func NewPage() *Page {
	bp := golang.NewBlueprint()
	return &Page{
		styles:            NewPageStyles(),
		currentStep:       StepProjectName,
		blueprint:         bp,
		multiSelectStates: make(map[string]bool),
		frameworks:  bp.GetSupportedFrameworks(),
		databases:   bp.GetSupportedDrivers(),
		allFeatures: bp.GetSupportedFeatures(),
		gitOptions:  []string{"init", "commit", "skip"},
		gitOption: "commit",
	}
}
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
	command := p.blueprint.GetCommandString(p.projectName, p.framework, p.database, p.gitOption, p.features)
	content = append(content, p.styles.FormLabel.Render("Command to execute:"))
	content = append(content, p.styles.Output.Render(command))
	content = append(content, "")
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
	style := p.styles.ButtonFocus
	content = append(content, style.Render("  Try Again  "))
	return content
}
type CreateProjectCmd struct {
	projectName string
	framework   string
	database    string
	features    []string
	gitOption   string
	blueprint   *golang.Blueprint
}
type ProjectCreatedMsg struct {
	output string
}
type ProjectErrorMsg struct {
	error  string
	output string
}
type ProjectDebugMsg struct {
	message string
}
func (c CreateProjectCmd) Execute() tea.Msg {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	var fullOutput strings.Builder
	if !c.blueprint.IsInstalled() {
		fullOutput.WriteString("📦 Installing go-blueprint CLI...\n")
		if err := c.blueprint.InstallCLI(ctx); err != nil {
			return ProjectErrorMsg{
				error:  fmt.Sprintf("Failed to install go-blueprint: %v", err),
				output: fullOutput.String(),
			}
		}
		fullOutput.WriteString("✅ go-blueprint CLI installed successfully\n\n")
	}
	database := c.database
	if database == "" {
		database = "none"
	}
	fullOutput.WriteString("🚀 Creating project...\n")
	if len(c.features) > 0 {
		fullOutput.WriteString("🎨 Features: " + strings.Join(c.features, ", ") + "\n")
	}
	if c.gitOption != "" && c.gitOption != "skip" {
		fullOutput.WriteString("📚 Git: " + c.gitOption + "\n")
	}
	gitOpt := c.gitOption
	if gitOpt == "skip" {
		gitOpt = ""
	}
	command := c.blueprint.GetCommandString(c.projectName, c.framework, database, gitOpt, c.features)
	fullOutput.WriteString("📋 Executing: " + command + "\n\n")
	fullOutput.WriteString("⚡ Running go-blueprint...\n")
	output, err := c.blueprint.CreateProjectWithOutput(ctx, c.projectName, c.framework, database, gitOpt, c.features)
	if output != "" {
		fullOutput.WriteString("📤 Command Output:\n")
		fullOutput.WriteString(output)
		fullOutput.WriteString("\n")
	}
	if err != nil {
		fullOutput.WriteString("❌ Error occurred: " + err.Error() + "\n")
		return ProjectErrorMsg{
			error:  fmt.Sprintf("Failed to create project: %v", err),
			output: fullOutput.String(),
		}
	}
	fullOutput.WriteString("🎉 Project created successfully!")
	return ProjectCreatedMsg{
		output: fullOutput.String(),
	}
}
func (p *Page) HandleInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	if msg.String() == "esc" {
		switch p.currentStep {
		case StepProjectName:
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
			p.selectedIndex = 1
			return true, nil
		case StepCreating:
			return true, nil
		case StepComplete:
			p.reset()
			return true, nil
		case StepError:
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
		p.features = []string{}
		for feature, selected := range p.multiSelectStates {
			if selected {
				p.features = append(p.features, feature)
			}
		}
		p.currentStep = StepGitOption
		p.selectedIndex = 1
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
		if p.selectedIndex == 0 {
			p.currentStep = StepCreating
			p.isCreating = true
			return true, CreateProjectCmd{
				projectName: p.projectName,
				framework:   p.framework,
				database:    p.database,
				features:    p.features,
				gitOption:   p.gitOption,
				blueprint:   p.blueprint,
			}.Execute
		} else {
			return false, nil
		}
		return true, nil
	}
	return true, nil
}
func (p *Page) handleCreatingInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	return true, nil
}
func (p *Page) handleCompleteInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "enter":
		p.reset()
		return true, nil
	}
	return true, nil
}
func (p *Page) handleErrorInput(msg tea.KeyMsg) (bool, tea.Cmd) {
	switch msg.String() {
	case "enter":
		p.currentStep = StepConfirm
		p.selectedIndex = 0
		return true, nil
	}
	return true, nil
}
func (p *Page) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case ProjectCreatedMsg:
		p.currentStep = StepComplete
		p.isCreating = false
		p.creationOutput = strings.Split(msg.output, "\n")
		return nil
	case ProjectErrorMsg:
		p.currentStep = StepError
		p.error = msg.error
		p.isCreating = false
		p.creationOutput = strings.Split(msg.output, "\n")
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
func (p *Page) GetTitle() string {
	return "Go Blueprint Creator"
}
func (p *Page) GetKeyBindings() []types.KeyBinding {
	bindings := []types.KeyBinding{
		{Key: "enter", Description: "Next/Confirm", Action: "next"},
		{Key: "↑/↓", Description: "Navigate", Action: "navigate"},
	}
	if p.currentStep == StepFeatures {
		bindings = append(bindings, types.KeyBinding{Key: "space", Description: "Toggle", Action: "toggle"})
	}
	if p.currentStep == StepProjectName {
		bindings = append(bindings, types.KeyBinding{Key: "esc", Description: "Back to Go Tools", Action: "back"})
	} else if p.currentStep != StepCreating {
		bindings = append(bindings, types.KeyBinding{Key: "esc", Description: "Previous Step", Action: "back"})
	}
	return bindings
}
