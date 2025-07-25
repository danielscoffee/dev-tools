build:
	go build -o ./bin/dev-tools.exe ./cmd/main.go

run:
	build
	./bin/dev-tools.exe

add-cobra:
	@echo "Adding a new Cobra command..."
	@read -p "Enter the command name: " cmd_name && \
		cobra-cli add $$cmd_name --parent ./internal/app/cli/ \
		--author ""

# TUI development commands
test-pages:
	@echo "Testing TUI pages system..."
	@go run -tags test ./internal/app/tui/pages/examples.go

run-tui:
	@echo "Running TUI interface..."
	@go run ./cmd/main.go tui

validate-pages:
	@echo "Validating page system integrity..."
	@go test ./internal/app/tui/pages/...

# CLI generation with output to ../cli
generate-cli:
	@echo "Generating CLI structure in ../cli..."
	@mkdir -p ./../cli
	@cobra-cli init --pkg-name github.com/danielscoffee/dev-tools ./../cli

.PHONY: tui tui-test tui-run
tui: build
	@echo "🚀 Starting Dev Tools TUI..."
	@./bin/dev-tools tui

tui-test:
	@echo "🧪 Running TUI tests..."
	@go test ./internal/app/tui/

tui-run: build tui

.PHONY: demo
demo: build
	@echo "📱 Running TUI demo (press Ctrl+C to exit)..."
	@timeout 30 ./bin/dev-tools tui || echo "Demo finished"

router-demo:
	@echo "🔀 Testing Router System..."
	@go run examples/router_demo.go

router-test: router-demo
	@echo "✅ Router system test complete"
