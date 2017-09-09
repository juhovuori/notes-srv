.PHONY: start-dev
stop-dev:
	docker stop --name minitwitter-postgres
	docker rm --name minitwitter-postgres

.PHONY: start-dev
start-dev:
	docker run --name minitwitter-postgres -p 5432:5432 -e POSTGRES_PASSWORD=minitwitter -d postgres


.PHONY: psql
createdb:
	docker run -it --rm --link minitwitter-postgres:postgres postgres createdb minitwitter -h postgres -U postgres

.PHONY: psql
psql:
	docker run -it --rm --link minitwitter-postgres:postgres postgres psql -h postgres -U postgres

.PHONY: build
build:
	go build

.PHONY: run
run: export DATABASE = postgres://postgres:minitwitter@localhost:5432/?sslmode=disable
run:
	./minitwitter-srv server
