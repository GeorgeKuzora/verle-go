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
build:
	go build -o verle-go cmd/cli/main.go
.PHONY:build
build-win:
	GOOS=windows GOARCH=amd64 go build -o verle-go.exe cmd/cli/main.go
.PHONY:build-win
run: vet
	go run cmd/cli/main.go
.PHONY:run
