.PHONY: all
all:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/release-version-tool *.go
