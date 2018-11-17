PROJECTNAME=demo
GOBASE=$(pwd)
GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin

build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/main src/main/main.go

run:
	@echo "  >  Running binary..."
	`pwd`/bin/main `pwd`/input.txt

test:
	@echo "  >  Testing binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test src/main/main.go