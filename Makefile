.PHONY: stop-postgres start-postgres migration-up migration-down

start-postgres:
	docker compose up -d

stop-postgres:
	docker compose down

DB_URL=postgres://postgres:postgres@localhost:15432/pocket_link?sslmode=disable
migrate-up:
	dbmate --url ${DB_URL} wait
	dbmate --url ${DB_URL} --migrations-dir ./db/migration up

migrate-down:
	dbmate --url ${DB_URL} --migrations-dir ./db/migration down

test:
	go test ./... -v

sqlc:
	sqlc generate