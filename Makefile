build:
	go build -o web ./cmd/web

run:
	go run ./cmd/web

live/go:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "go build  -o tmp/bin/main ./cmd/web" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

live/templ:
	templ generate --watch --proxy="http://localhost:4000" --cmd="go run ./cmd/web"

