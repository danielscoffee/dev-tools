// Package executor provides utilities for executing system commands
package executor
import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)
type CommandExecutor struct {
	WorkingDir string
	Timeout    time.Duration
	Env        []string
}
func NewExecutor() *CommandExecutor {
	return &CommandExecutor{
		WorkingDir: ".",
		Timeout:    30 * time.Second,
		Env:        os.Environ(),
	}
}
func (e *CommandExecutor) WithWorkingDir(dir string) *CommandExecutor {
	e.WorkingDir = dir
	return e
}
func (e *CommandExecutor) WithTimeout(timeout time.Duration) *CommandExecutor {
	e.Timeout = timeout
	return e
}
func (e *CommandExecutor) WithEnv(env []string) *CommandExecutor {
	e.Env = env
	return e
}
type CommandResult struct {
	Command    string
	Args       []string
	Stdout     string
	Stderr     string
	ExitCode   int
	Error      error
	Duration   time.Duration
	WorkingDir string
}
func (e *CommandExecutor) Execute(ctx context.Context, command string, args ...string) *CommandResult {
	start := time.Now()
	if e.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, e.Timeout)
		defer cancel()
	}
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir = e.WorkingDir
	cmd.Env = e.Env
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	duration := time.Since(start)
	result := &CommandResult{
		Command:    command,
		Args:       args,
		Stdout:     stdout.String(),
		Stderr:     stderr.String(),
		Error:      err,
		Duration:   duration,
		WorkingDir: e.WorkingDir,
	}
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitError.ExitCode()
		} else {
			result.ExitCode = -1
		}
	}
	return result
}
func (e *CommandExecutor) ExecuteShell(ctx context.Context, command string) *CommandResult {
	return e.Execute(ctx, "sh", "-c", command)
}
func (e *CommandExecutor) ExecuteInteractive(ctx context.Context, command string, args ...string) error {
	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir = e.WorkingDir
	cmd.Env = e.Env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func (r *CommandResult) String() string {
	var output strings.Builder
	output.WriteString(fmt.Sprintf("Command: %s %s\n", r.Command, strings.Join(r.Args, " ")))
	output.WriteString(fmt.Sprintf("Working Dir: %s\n", r.WorkingDir))
	output.WriteString(fmt.Sprintf("Duration: %v\n", r.Duration))
	output.WriteString(fmt.Sprintf("Exit Code: %d\n", r.ExitCode))
	if r.Stdout != "" {
		output.WriteString("\n--- STDOUT ---\n")
		output.WriteString(r.Stdout)
	}
	if r.Stderr != "" {
		output.WriteString("\n--- STDERR ---\n")
		output.WriteString(r.Stderr)
	}
	if r.Error != nil {
		output.WriteString(fmt.Sprintf("\n--- ERROR ---\n%v\n", r.Error))
	}
	return output.String()
}
func (r *CommandResult) Success() bool {
	return r.ExitCode == 0 && r.Error == nil
}
func (r *CommandResult) Failed() bool {
	return !r.Success()
}
func (r *CommandResult) Output() string {
	var output strings.Builder
	if r.Stdout != "" {
		output.WriteString(r.Stdout)
	}
	if r.Stderr != "" {
		if output.Len() > 0 {
			output.WriteString("\n")
		}
		output.WriteString(r.Stderr)
	}
	return output.String()
}
