include .env

build:
	go build -o web ./cmd/web

build/reload:
	go build -o tmp/bin/web ./cmd/web


run:
	@go run ./cmd/web \
	-addr=${ADDR} \
	-mode=${MODE}

db/psql:
	@psql ${MARKBIN_DB_DSN}

db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	@migrate create -seq -ext=.sql -dir=./migrations ${name}

db/migrations/up:
	@echo 'Running up migrations...'
	@migrate -path ./migrations -database ${MARKBIN_DB_DSN} up

db/migrations/down:
	@echo 'Running up migrations...'
	@migrate -path ./migrations -database ${MARKBIN_DB_DSN} down
		

live/go:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--build.cmd "make build/reload" \
	--build.bin "./tmp/bin/web -addr=${ADDR} -mode=${MODE}" \
	--build.delay "100" \
	--misc.clean_on_exit true 

live/templ:
	templ generate --watch --proxy="http://localhost${ADDR}" 

live/tailwind:
	npx @tailwindcss/cli -i ./input.css -o ./ui/static/css/styles.css --watch
