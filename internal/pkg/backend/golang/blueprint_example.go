package golang

import (
	"context"
	"fmt"
	"log"
	"time"
)

// ExampleUsage demonstrates how to use the Blueprint CLI integration
func ExampleUsage() {
	bp := &Blueprint{}
	ctx := context.Background()

	// Check if go-blueprint is installed
	if !bp.IsInstalled() {
		fmt.Println("Installing go-blueprint CLI...")
		if err := bp.InstallCLI(ctx); err != nil {
			log.Printf("Failed to install go-blueprint: %v", err)
			return
		}
		fmt.Println("✓ go-blueprint installed")
	}

	// Example 1: Simple API with Gin and PostgreSQL
	fmt.Println("\n=== Example 1: Simple REST API ===")
	fmt.Println("Command:", bp.GetCommandString("my-api", "gin", "postgres", "commit", nil, false))
	// err := bp.CreateSimple(ctx, "my-api", "gin", "postgres")

	// Example 2: Web app with Chi, SQLite, and HTMX
	fmt.Println("\n=== Example 2: Web App with HTMX ===")
	features := []string{"htmx", "tailwind"}
	fmt.Println("Command:", bp.GetCommandString("my-webapp", "chi", "sqlite", "commit", features, true))
	// err := bp.CreateAdvanced(ctx, "my-webapp", "chi", "sqlite", features)

	// Example 3: Full-stack app with React
	fmt.Println("\n=== Example 3: React Full-stack App ===")
	reactFeatures := []string{"react", "docker", "githubaction"}
	fmt.Println("Command:", bp.GetCommandString("fullstack-app", "fiber", "mongo", "commit", reactFeatures, true))
	// err := bp.CreateAdvanced(ctx, "fullstack-app", "fiber", "mongo", reactFeatures)

	// Example 4: Microservice with minimal setup
	fmt.Println("\n=== Example 4: Microservice ===")
	microFeatures := []string{"docker"}
	fmt.Println("Command:", bp.GetCommandString("user-service", "echo", "redis", "init", microFeatures, true))
	// err := bp.CreateAdvanced(ctx, "user-service", "echo", "redis", microFeatures)

	// Show supported options
	fmt.Println("\n=== Supported Options ===")
	fmt.Println("Frameworks:", bp.GetSupportedFrameworks())
	fmt.Println("Drivers:", bp.GetSupportedDrivers())
	fmt.Println("Features:", bp.GetSupportedFeatures())
}

// CreateSimpleAPI creates a basic REST API
func CreateSimpleAPI(projectName, framework, database string) error {
	bp := &Blueprint{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Install CLI if needed
	if !bp.IsInstalled() {
		fmt.Println("Installing go-blueprint...")
		if err := bp.InstallCLI(ctx); err != nil {
			return fmt.Errorf("failed to install go-blueprint: %w", err)
		}
	}

	fmt.Printf("Creating project: %s\n", projectName)
	fmt.Printf("Framework: %s, Database: %s\n", framework, database)

	return bp.CreateWithGit(ctx, projectName, framework, database, "commit")
}

// CreateAdvancedProject creates a project with advanced features
func CreateAdvancedProject(projectName, framework, database string, features []string) error {
	bp := &Blueprint{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	if !bp.IsInstalled() {
		fmt.Println("Installing go-blueprint...")
		if err := bp.InstallCLI(ctx); err != nil {
			return fmt.Errorf("failed to install go-blueprint: %w", err)
		}
	}

	fmt.Printf("Creating advanced project: %s\n", projectName)
	fmt.Printf("Framework: %s, Database: %s\n", framework, database)
	fmt.Printf("Features: %v\n", features)

	return bp.CreateAdvanced(ctx, projectName, framework, database, features)
}

// Interactive example showing different project types
func ShowProjectTypes() {
	bp := &Blueprint{}

	projects := []struct {
		name        string
		description string
		framework   string
		database    string
		features    []string
		command     string
	}{
		{
			name:        "rest-api",
			description: "Simple REST API",
			framework:   "gin",
			database:    "postgres",
			features:    []string{"docker"},
			command:     bp.GetCommandString("rest-api", "gin", "postgres", "commit", []string{"docker"}, true),
		},
		{
			name:        "web-app",
			description: "Web application with HTMX",
			framework:   "chi",
			database:    "sqlite",
			features:    []string{"htmx", "tailwind"},
			command:     bp.GetCommandString("web-app", "chi", "sqlite", "commit", []string{"htmx", "tailwind"}, true),
		},
		{
			name:        "realtime-app",
			description: "Real-time app with WebSockets",
			framework:   "fiber",
			database:    "redis",
			features:    []string{"websocket", "docker"},
			command:     bp.GetCommandString("realtime-app", "fiber", "redis", "commit", []string{"websocket", "docker"}, true),
		},
		{
			name:        "fullstack-react",
			description: "Full-stack with React frontend",
			framework:   "echo",
			database:    "mongo",
			features:    []string{"react", "docker", "githubaction"},
			command:     bp.GetCommandString("fullstack-react", "echo", "mongo", "commit", []string{"react", "docker", "githubaction"}, true),
		},
	}

	fmt.Println("=== Go Blueprint Project Templates ===")
	fmt.Println()

	for _, p := range projects {
		fmt.Printf("📁 %s - %s\n", p.name, p.description)
		fmt.Printf("   Framework: %s | Database: %s\n", p.framework, p.database)
		fmt.Printf("   Features: %v\n", p.features)
		fmt.Printf("   Command: %s\n", p.command)
		fmt.Println()
	}
}
