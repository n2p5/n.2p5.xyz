
DIST_DIR := dst

.PHONY: generate build run dev

deploy:
	./scripts/deploy.sh

generate:
	@templ generate

build: generate
	@go run .

run: build
	@cd $(DIST_DIR) && python -m http.server 8000 

dev:
	@find . -type f \( \
	-name '*.go' ! -name '*_templ.go' -o -name '*.templ' \) \
	| entr -r sh -c 'make run'