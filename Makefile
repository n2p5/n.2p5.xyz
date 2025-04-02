
DIST_DIR := dst

.PHONY: generate build run dev

deploy: generate
	./scripts/deploy.sh

generate:
	@templ generate

build: generate
	@go run .

serve: build
	@cd $(DIST_DIR) && python -m http.server 8000

# dev watches a set of files and if anything changes
# it will re-run the application gen, build and serve processes
dev:
	@find . -type f \( \
	-name '*.go' \
	-o -name '*.toml' \
	-o -name '*.templ' \) \
	! -name '*_templ.go' \
	| entr -r sh -c 'make serve'