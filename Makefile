API_BIN=dist/webhook-fwd-api
API_SRC=src/cmd/api/main.go

CLI_BIN=bin/webhook-fwd
CLI_SRC=src/cmd/cli/main.go

.DEFAULT_GOAL := webserver

build_api:
	go build -o $(API_BIN) $(API_SRC)

build_cli:
	go build -o $(CLI_BIN) $(CLI_SRC)

webserver:
	go run $(API_SRC)

dist_cli:
	GOOS=windows GOARCH=amd64 go build -o "$(CLI_BIN)-win64.exe" $(CLI_SRC)
	GOOS=darwin GOARCH=amd64 go build -o "$(CLI_BIN)-mac-amd64" $(CLI_SRC)
	GOOS=darwin GOARCH=arm64 go build -o "$(CLI_BIN)-mac-arm64" $(CLI_SRC)
	GOOS=linux GOARCH=amd64 go build -o "$(CLI_BIN)-linux-amd64" $(CLI_SRC)

test:
	go test ./src/...

.PHONY: build_api build_cli dist_cli test