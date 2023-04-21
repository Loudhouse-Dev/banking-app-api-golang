postgres:
	docker run --name banking-app -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Autumn0721 -d postgres:15.2-alpine

createdb:
	docker exec -it banking-app createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it banking-app dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:Autumn0721@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:Autumn0721@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown