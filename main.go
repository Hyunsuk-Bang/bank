package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gitlab.com/hbang3/simple_bank/api"
	db "gitlab.com/hbang3/simple_bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:root@localhost:5432/bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
}
