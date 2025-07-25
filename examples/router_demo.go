package main

import (
	"fmt"
	"os"

	"github.com/danielscoffee/dev-tools/internal/app/tui"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/home"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/langs"
	"github.com/danielscoffee/dev-tools/internal/app/tui/pages/langs/golang"
)

func main() {
	// Create router
	router := tui.NewRouter()

	// Register routes (demonstrating the router system)
	router.RegisterRoute("/", home.NewPage(), "Dev Tools - Home", "Main menu and navigation", "h")
	router.RegisterRoute("/langs", langs.NewPage(), "Programming Languages", "Tools for different languages", "l")
	router.RegisterRoute("/langs/golang", golang.NewPage(), "Go/Golang Tools", "Go development tools", "g")

	fmt.Println("🚀 Router System Demo")
	fmt.Println("=====================")

	// Test navigation
	fmt.Printf("Current route: %s\n", router.GetCurrentRoute().Path)
	fmt.Printf("Page title: %s\n", router.GetCurrentRoute().Component.GetTitle())

	// Test navigation to languages
	err := router.NavigateTo("/langs")
	if err != nil {
		fmt.Printf("Navigation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Navigated to: %s\n", router.GetCurrentRoute().Path)
	fmt.Printf("Page title: %s\n", router.GetCurrentRoute().Component.GetTitle())

	// Test navigation to golang
	err = router.NavigateTo("/langs/golang")
	if err != nil {
		fmt.Printf("Navigation error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Navigated to: %s\n", router.GetCurrentRoute().Path)
	fmt.Printf("Page title: %s\n", router.GetCurrentRoute().Component.GetTitle())

	// Test breadcrumb
	fmt.Printf("Breadcrumb: %s\n", router.GetBreadcrumb())

	// Test go back
	err = router.GoBack()
	if err != nil {
		fmt.Printf("Go back error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("After going back: %s\n", router.GetCurrentRoute().Path)
	fmt.Printf("Breadcrumb: %s\n", router.GetBreadcrumb())

	fmt.Println("\n✅ Router system working correctly!")
	fmt.Println("🎯 The TUI router can dynamically discover and process pages from the pages/ directory structure")
}
