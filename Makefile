.PHONY: tailwind-watch
tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify

.PHONY: templ-watch
templ-watch:
	templ generate --watch --proxy="http://localhost:4000" --cmd="go run ./cmd/main.go"
	
.PHONY: dev
dev:
	go build -o ./tmp/ ./cmd/main.go && air

.PHONY: build
build:
	make tailwind-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/ ./cmd/main.go

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	go test -race -v -timeout 30s ./...

.PHONY: killall
killall:
	@echo "Stopping all development processes..."
	-killall templ
	-killall air
	-killall tailwindcss
	@echo "All processes stopped."

.PHONY: install-templ
install-templ:
	@which templ > /dev/null || go install github.com/a-h/templ/cmd/templ@latest

.PHONY: templ-generate
templ-generate: install-templ
	templ generate

.PHONY: build-netlify
build-netlify:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.min.css --minify
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go