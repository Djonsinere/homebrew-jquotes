.PHONY: build install

build:
	go build -o japanese_quotes main.go

install: build
	@echo "Add this to your shell config (.zshrc or .bashrc):"
	@echo "japanese_quotes"