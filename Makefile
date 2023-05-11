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

migrateup1:
					migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
					migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
					migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
					sqlc generate

test:
					go test -v ./...

server:
					go run main.go

mock:
					mockgen -build_flags=--mod=mod -destination db/mock/store.go -package mockdb github.com/Kcih4518/simpleBank/db/sqlc Store

run:
					docker run --name simplebank -p 8080:8080 simplebank:latest

build:
					docker build -t simplebank:latest .

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 run build
