.PHONY: all build run clean test

BOOSTER=booster
NAME=test
LANG=c
DOMAIN=github.com/mrtuuro/test

build:
	@go build -o ./bin/$(BOOSTER) ./*.go


run: build
	@clear
	@./bin/$(BOOSTER) -name=$(NAME) -lang=$(LANG) -domain=$(DOMAIN)

clean:
	@echo "Cleaning up..."
	@go clean
	@rm -rf ./bin

test:
	@go  test ./...
