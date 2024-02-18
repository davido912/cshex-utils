BINARY_NAME=cshex

any: help

build: 
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME} main.go

run: build ## builds the binary and runs the application
	./${BINARY_NAME}

install: build ## builds the binary and installs it on your system (requires sudo)
	mv ${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

uninstall: ## removes the application from system (requires sudo)
	rm /usr/local/bin/${BINARY_NAME}

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'