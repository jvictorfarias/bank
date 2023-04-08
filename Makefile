SHELL=cmd
postgres:
	docker run --name bank-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.1-alpine

createdb:
	docker exec -it bank-postgres createdb --username=root --owner=root bank

dropdb:
	docker exec -it bank-postgres dropdb bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "C:\Users\victo\Documents\Backend-Masterclass\bank:/src" -w /src kjconroy/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc

