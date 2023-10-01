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

migrate_up1:
	migrate -path db/migration/ -database "postgres://root:root@localhost:5432/bank?sslmode=disable" -verbose up 1

migrate_down1:
	migrate -path db/migration/ -database "postgres://root:root@localhost:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go gitlab.com/hbang3/simple_bank/db/sqlc Store 

.PHONY: postgres create_db drop_db migrate_up migrate_down migrate_up1 migrate_down1 server sqlc mock
