ifneq (,$(wildcard ./backend/.env))
    include ./backend/.env
    export
endif

.PHONY: migrate-up
migrate-up:
	migrate -database ${DB_URL_LOCALHOST} -path backend/db/migrations up

.PHONY: migrate-down
migrate-down:
	migrate -database ${DB_URL_LOCALHOST} -path backend/db/migrations down

# make migrate-create NAME=<name>
.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir backend/db/migrations -seq $(NAME)

# make migrate-force-version VERSION=<version>
.PHONY: migrate-force-version
migrate-force-version:
	migrate -database ${DB_URL_LOCALHOST} -path backend/db/migrations force $(VERSION)