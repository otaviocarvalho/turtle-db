APPNAME=turtle-cli

.PHONY = build run

build:
	go build -o $(APPNAME) ./cmd/turtle-cli

run: build
	./$(APPNAME)
