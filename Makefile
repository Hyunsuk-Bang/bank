postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine

create_db:
	docker exec -it postgres16 createdb --username=root --owner=root bank

drop_db:
	docker exec -it postgres16 dropdb bank

migrate_up:
	migrate -path db/migration/ -database "postgres://root:root@localhost:5432/bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration/ -database "postgres://root:root@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres create_db drop_db migrate_up migrate_down
