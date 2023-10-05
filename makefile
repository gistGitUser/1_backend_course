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

migrateup1:
	migrate -path project/migration -database "postgresql://back:back@localhost:9876/back_db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path project/migration -database "postgresql://back:back@localhost:9876/back_db?sslmode=disable" -verbose down

migratedown1:
	migrate -path project/migration -database "postgresql://back:back@localhost:9876/back_db?sslmode=disable" -verbose down 1

mock:
	mockgen -package mockdb -destination project/mock/store.go github.com/gistGitUser/course/project/sqlc Store

test-coverage:
	go clean -testcache
	go test -coverprofile cover.out ./...

test-coverage-view:
	go clean -testcache
	go test -coverprofile coverage.out ./...
	go tool cover -html=coverage.out

sqlc:
	sqlc generate

test:
	go clean -testcache
	go test -v -cover ./...

server:
	cd project; go run main.go

.PHONY: postgres postgresstart createdb dropdb migrateup migratedown test