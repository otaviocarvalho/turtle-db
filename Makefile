APPNAME=turtle-cli

.PHONY = build run fmt vet

fmt: 
	go fmt ./...

vet: 
	go vet ./...

build:
	go build -o $(APPNAME) ./cmd/turtle-cli

run: build
	./$(APPNAME)
