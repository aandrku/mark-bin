include .env

.PHONY: build/web
build/web:
	go build -o web ./cmd/web

.PHONY: build/reload
build/reload:
	go build -o tmp/bin/web ./cmd/web

.PHONY: run/web
run/web:
	@go run ./cmd/web \
	-addr=${ADDR} \
	-mode=${MODE}

.PHONY: db/psql
db/psql:
	@psql ${MARKBIN_DB_DSN}


.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	@migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: db/migrations/up
db/migrations/up:
	@echo 'Running up migrations...'
	@migrate -path ./migrations -database ${MARKBIN_DB_DSN} up

.PHONY: db/migrations/down
db/migrations/down:
	@echo 'Running up migrations...'
	@migrate -path ./migrations -database ${MARKBIN_DB_DSN} down
		

	go run github.com/cosmtrek/air@v1.51.0 \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--build.cmd "make build/reload" \
	--build.bin "./tmp/bin/web -addr=${ADDR} -mode=${MODE}" \
	--build.delay "100" \
	--misc.clean_on_exit true 

.PHONY: live/templ
live/templ:
	templ generate --watch --proxy="http://localhost${ADDR}" 

.PHONY: live/templ
live/tailwind:
	npx @tailwindcss/cli -i ./input.css -o ./ui/static/css/styles.css --watch
