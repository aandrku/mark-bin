include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## build/web: build main web application
.PHONY: build/web
build/web:
	go build -o web ./cmd/web

## build/reload: build main web application, but for hot-reloading
.PHONY: build/reload
build/reload:
	go build -o tmp/bin/web ./cmd/web

## run/web: run main web application
.PHONY: run/web
run/web:
	@go run ./cmd/web \
	-addr=${ADDR} \
	-mode=${MODE} \
	-dsn=${MARKBIN_DB_DSN}

## db/psql: connect to local database using psql
.PHONY: db/psql
db/psql:
	@psql ${MARKBIN_DB_DSN}


## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	@migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	@migrate -path ./migrations -database ${MARKBIN_DB_DSN} up

## db/migrations/down: remove all database migrations
.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running up migrations...'
	@migrate -path ./migrations -database ${MARKBIN_DB_DSN} down
		

## live/go: start hot reloading
.PHONY: live/go
live/go:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--build.cmd "make build/reload" \
	--build.bin "./tmp/bin/web -addr=${ADDR} -mode=${MODE} -dsn=${MARKBIN_DB_DSN}" \
	--build.delay "100" \
	--misc.clean_on_exit true 

## live/templ: start hot reloading templ files
.PHONY: live/templ
live/templ:
	templ generate --watch --proxy="http://localhost${ADDR}" 

## live/tailwind: start hot reloading tailwind
.PHONY: live/tailwind
live/tailwind:
	npx @tailwindcss/cli -i ./input.css -o ./ui/static/css/styles.css --watch
