package golang

import (
	"context"
	"os"
	"os/exec"
	"strings"
)

// Blueprint integrates with the go-blueprint CLI tool from Melkeydev
// Repository: https://github.com/Melkeydev/go-blueprint
type Blueprint struct{}

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

func (b *Blueprint) CreateSimple(ctx context.Context, projectName, framework, driver string) error {
	args := []string{"--framework", framework}
	if driver != "" {
		args = append(args, "--driver", driver)
	}
	return b.Create(ctx, projectName, args...)
}

func (b *Blueprint) CreateAdvanced(ctx context.Context, projectName, framework, driver string, features []string) error {
	args := []string{"--framework", framework, "--advanced"}

	if driver != "" && driver != "none" {
		args = append(args, "--driver", driver)
	}

	for _, feature := range features {
		args = append(args, "--feature", feature)
	}

	return b.Create(ctx, projectName, args...)
}

func (b *Blueprint) CreateWithGit(ctx context.Context, projectName, framework, driver, gitOption string) error {
	args := []string{"--framework", framework, "--git", gitOption}
	if driver != "" && driver != "none" {
		args = append(args, "--driver", driver)
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

func (b *Blueprint) BuildCommand(projectName, framework, driver, gitOption string, features []string, advanced bool) []string {
	args := []string{"create", "--name", projectName}

	if framework != "" {
		args = append(args, "--framework", framework)
	}

	if driver != "" && driver != "none" {
		args = append(args, "--driver", driver)
	}

	if gitOption != "" && gitOption != "skip" {
		args = append(args, "--git", gitOption)
	}

	if advanced {
		args = append(args, "--advanced")
	}

	for _, feature := range features {
		args = append(args, "--feature", feature)
	}

	return args
}

func (b *Blueprint) GetCommandString(projectName, framework, driver, gitOption string, features []string, advanced bool) string {
	args := b.BuildCommand(projectName, framework, driver, gitOption, features, advanced)
	return "go-blueprint " + strings.Join(args, " ")
}
