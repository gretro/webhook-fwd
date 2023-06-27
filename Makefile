API_BIN=dist/webhook-fwd-api
API_SRC=cmd/api/main.go

CLI_BIN=bin/webhook-fwd
CLI_SRC=cmd/cli/main.go

.DEFAULT_GOAL := webserver

clean:
	rm -rf bin/ dist/
	go clean -cache
	go clean -testcache

build_api:
	go build -o $(API_BIN) $(API_SRC)

debug_api:
	go build -o $(API_BIN) -gcflags="all=-N -l" $(API_SRC)

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
	go test -count=1 -timeout=5m ./...

race:
	go test -race -count=1 -timeout=5m ./...

vet:
	go test -count=1 -vet="" ./...

.PHONY: clean build_api build_cli webserver dist_cli test race vet