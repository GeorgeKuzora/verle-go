.DEFAULT_GOAL := run
fmt:
	go fmt ./...
.PHONY:fmt
lint: fmt
	staticcheck ./...
.PHONY:lint
vet: lint
	go vet ./...
.PHONY:vet
build: vet
	go build -o verle-go.exe cmd/cli/main.go
.PHONY:build
run: vet
	go run cmd/cli/main.go
.PHONY:run