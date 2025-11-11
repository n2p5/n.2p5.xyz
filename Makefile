DIST_DIR := dst

.PHONY: build serve dev deploy

deploy: build
	./scripts/deploy.sh

build:
	@go run .

serve: build
	@cd $(DIST_DIR) && python -m http.server 8000

# dev watches files and rebuilds/serves on changes
dev:
	@find . -type f \( \
	-name '*.go' \
	-o -name '*.json' \
	-o -name '*.gohtml' \
	-o -name '*.css' \) \
	| entr -r sh -c 'make serve'