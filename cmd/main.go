package main

import "github.com/danielscoffee/dev-tools/internal/app/cli"

func main() {
	c := new(cli.CLI)
	c.Execute()
}
