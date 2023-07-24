postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	sudo docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path ./schema -database "postgresql://postgres:123456@localhost:5432/db_done?sslmode=disable" -verbose up

migratedown:
	migrate -path schema/migration -database "postgresql://postgres:123456@localhost:5432/db_done?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run cmd/main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ovasquezbrito/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock