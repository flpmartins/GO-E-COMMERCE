# Makefile para projetos Go

.PHONY: lint format

# Comando para rodar linters
lint:
	@echo "🔍 Running golangci-lint..."
	golangci-lint run ./...

# Comando para formatar código
format:
	@echo "🎨 Formatting code with goimports and gofmt..."
	goimports -w .
	gofmt -w .