DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
					podman stop postgres
					podman rm postgres
					podman run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
					podman exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
					podman exec -it postgres dropdb simple_bank

migrateup:
					migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
					migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
					sqlc generate

test:
					go test -v ./...

server:
					go run main.go

mock:
					mockgen -build_flags=--mod=mod -destination db/mock/store.go -package mockdb github.com/Kcih4518/simpleBank/db/sqlc Store

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc test server mock
