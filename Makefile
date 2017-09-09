DB=postgres://postgres:notes@localhost:5432/?sslmode=disable

.PHONY: start-dev
stop-dev:
	docker stop notes-postgres
	docker rm notes-postgres

.PHONY: start-dev
start-dev:
	docker run --name notes-postgres -p 5432:5432 -e POSTGRES_PASSWORD=notes -d postgres


.PHONY: psql
createdb:
	docker run -it --rm --link notes-postgres:postgres postgres createdb notes -h postgres -U postgres

.PHONY: migrate
migrate: export DATABASE = ${DB}
migrate:
	./notes-srv migrate

.PHONY: psql
psql:
	docker run -it --rm --link notes-postgres:postgres postgres psql -h postgres -U postgres

.PHONY: build
build:
	go build

.PHONY: run
run: export DATABASE = ${DB}
run:
	./notes-srv server
