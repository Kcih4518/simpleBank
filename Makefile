DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
AWS_DB_URL=postgresql://root:EyX0uCTu3qnb2bCclUzT@simple-bank.ci28ofkmintl.us-east-1.rds.amazonaws.com:5432/simple_bank

network:
	docker network create bank-network

postgres:
					docker stop postgres
					docker rm postgres
					docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret --network bank-network -d postgres:14-alpine

createdb:
					docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
					docker exec -it postgres dropdb simple_bank

migrateup:
					migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
					migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
					migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
					migrate -path db/migration -database "$(DB_URL)" -verbose down 1

migrateup_aws:
					migrate -path db/migration -database "$(AWS_DB_URL)" -verbose up

migrateup1_aws:
					migrate -path db/migration -database "$(AWS_DB_URL)" -verbose up 1

migratedown_aws:
					migrate -path db/migration -database "$(AWS_DB_URL)" -verbose down

migratedown1_aws:
					migrate -path db/migration -database "$(AWS_DB_URL)" -verbose down 1

sqlc:
					sqlc generate

test:
					go test -v ./...

server:
					go run main.go

mock:
					mockgen -build_flags=--mod=mod -destination db/mock/store.go -package mockdb github.com/Kcih4518/simpleBank/db/sqlc Store

run:
					docker run --name simplebank --network bank-network -p 8080:8080 -e DB_SOURCE=postgresql://root:secret@172.18.0.2:5432/simple_bank?sslmode=disable -e GIN_MODE=release simplebank:latest

build:
					docker build -t simplebank:latest .

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 run build migrateup_aws migrateup1_aws migratedown_aws migratedown1_aws
