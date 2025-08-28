// Package golang
package golang
import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"strings"
	"github.com/danielscoffee/dev-tools/internal/pkg/executor"
)
type Blueprint struct {
	executor *executor.CommandExecutor
}
func NewBlueprint() *Blueprint {
	return &Blueprint{
		executor: executor.NewExecutor(),
	}
}
func (b *Blueprint) WithWorkingDir(dir string) *Blueprint {
	b.executor = b.executor.WithWorkingDir(dir)
	return b
}
func (b *Blueprint) ExecuteCommand(ctx context.Context, args ...string) *executor.CommandResult {
	return b.executor.Execute(ctx, "go-blueprint", args...)
}
func (b *Blueprint) InstallCLI(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "go", "install", "github.com/melkeydev/go-blueprint@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func (b *Blueprint) Create(ctx context.Context, projectName string, args ...string) error {
	cmdArgs := []string{"create", "--name", projectName}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.CommandContext(ctx, "go-blueprint", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
func (b *Blueprint) CreateWithOutput(ctx context.Context, projectName string, args ...string) (string, error) {
	cmdArgs := []string{"create", "--name", projectName}
	cmdArgs = append(cmdArgs, args...)
	fullCmd := "go-blueprint " + strings.Join(cmdArgs, " ")
	cmd := exec.CommandContext(ctx, "go-blueprint", cmdArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	output := "DEBUG: Executed command: " + fullCmd + "\n\n"
	if stdout.String() != "" {
		output += "STDOUT:\n" + stdout.String()
	}
	if stderr.String() != "" {
		output += "\nSTDERR:\n" + stderr.String()
	}
	if stdout.String() == "" && stderr.String() == "" {
		output += "No output from command (this might be normal for go-blueprint)"
	}
	return output, err
}
func (b *Blueprint) CreateProjectWithOutputV2(ctx context.Context, projectName, framework, driver, gitOption string, features []string) *executor.CommandResult {
	args := []string{"create", "--name", projectName, "--framework", framework, "--driver", driver}
	if gitOption != "" {
		args = append(args, "--git", gitOption)
	}
	for _, feature := range features {
		args = append(args, "--feature", feature)
	}
	return b.ExecuteCommand(ctx, args...)
}
func (b *Blueprint) CreateProjectWithOutput(ctx context.Context, projectName, framework, driver, gitOption string, features []string) (string, error) {
	args := []string{"--framework", framework, "--driver", driver}
	if gitOption != "" {
		args = append(args, "--git", gitOption)
	}
	for _, feature := range features {
		args = append(args, "--feature", feature)
	}
	return b.CreateWithOutput(ctx, projectName, args...)
}
func (b *Blueprint) CreateSimpleWithOutput(ctx context.Context, projectName, framework, driver string) (string, error) {
	args := []string{"--framework", framework, "--driver", driver}
	return b.CreateWithOutput(ctx, projectName, args...)
}
func (b *Blueprint) CreateAdvancedWithOutput(ctx context.Context, projectName, framework, driver string, features []string) (string, error) {
	args := []string{"--framework", framework, "--driver", driver}
	for _, feature := range features {
		args = append(args, "--feature", feature)
	}
	return b.CreateWithOutput(ctx, projectName, args...)
}
func (b *Blueprint) CreateWithGitAndOutput(ctx context.Context, projectName, framework, driver, gitOption string) (string, error) {
	args := []string{"--framework", framework, "--driver", driver, "--git", gitOption}
	return b.CreateWithOutput(ctx, projectName, args...)
}
func (b *Blueprint) InstallCLIWithOutput(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "go", "install", "github.com/melkeydev/go-blueprint@latest")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	output := stdout.String()
	if stderr.String() != "" {
		output += "\nErrors:\n" + stderr.String()
	}
	return output, err
}
func (b *Blueprint) CreateAdvanced(ctx context.Context, projectName, framework, driver string, features []string) error {
	args := []string{"--framework", framework, "--driver", driver}
	for _, feature := range features {
		args = append(args, "--feature", feature)
	}
	return b.Create(ctx, projectName, args...)
}
func (b *Blueprint) CreateWithGit(ctx context.Context, projectName, framework, driver, gitOption string) error {
	args := []string{"--framework", framework, "--driver", driver, "--git", gitOption}
	return b.Create(ctx, projectName, args...)
}
func (b *Blueprint) CreateProject(ctx context.Context, projectName, framework, driver, gitOption string, features []string) error {
	args := []string{"--framework", framework, "--driver", driver}
	if gitOption != "" {
		args = append(args, "--git", gitOption)
	}
	for _, feature := range features {
		args = append(args, "--feature", feature)
	}
	return b.Create(ctx, projectName, args...)
}
func (b *Blueprint) IsInstalled() bool {
	_, err := exec.LookPath("go-blueprint")
	return err == nil
}
func (b *Blueprint) RunCommand(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "go-blueprint", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
func (b *Blueprint) GetSupportedFrameworks() []string {
	return []string{
		"chi",
		"gin",
		"fiber",
		"echo",
		"gorillamux",
		"httprouter",
		"standardlibrary",
	}
}
func (b *Blueprint) GetSupportedDrivers() []string {
	return []string{
		"none",
		"mysql",
		"postgres",
		"sqlite",
		"mongo",
		"redis",
		"scylla",
	}
}
func (b *Blueprint) GetSupportedFeatures() []string {
	return []string{
		"htmx",
		"githubaction",
		"websocket",
		"tailwind",
		"docker",
		"react",
	}
}
func (b *Blueprint) BuildCommand(projectName, framework, driver, gitOption string, features []string) []string {
	args := []string{"create", "--name", projectName}
	if framework != "" {
		args = append(args, "--framework", framework)
	}
	args = append(args, "--driver", driver)
	if gitOption != "" && gitOption != "skip" {
		args = append(args, "--git", gitOption)
	}
	for _, feature := range features {
		args = append(args, "--feature", feature)
	}
	return args
}
func (b *Blueprint) GetCommandString(projectName, framework, driver, gitOption string, features []string) string {
	args := b.BuildCommand(projectName, framework, driver, gitOption, features)
	return "go-blueprint " + strings.Join(args, " ")
}
