

postgres:
	docker run --name backend_corse_postgres -e POSTGRES_PASSWORD=back -e POSTGRES_USER=back -p 9876:5432 -d postgres:14.8-alpine

postgresstart:
	docker start backend_corse_postgres

createdb:
	docker exec -it backend_corse_postgres createdb --username=back --owner=back back_db

dropdb:
	docker exec -it backend_corse_postgres dropdb -U back back_db

migrateup:
	migrate -path project/migration -database "postgresql://back:back@localhost:9876/back_db?sslmode=disable" -verbose up

migratedown:
	migrate -path project/migration -database "postgresql://back:back@localhost:9876/back_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown test