DB_URL="postgresql://root:secret@localhost:5432/workflows?sslmode=disable"
postgres:
	docker run --name postgres-latest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres 
createdb: 
	docker exec -it postgres-latest createdb --username=root --owner=root workflows
migrateup:
	migrate -path src/db/migration -database $(DB_URL) -verbose up  
migratedown:
	migrate -path src/db/migration -database $(DB_URL) -verbose down 
.PHONY: dev
sqlc:
	sqlc generate
dev:
	@MODE=dev go run main.go
pro:
	@MODE=pro go run main.go