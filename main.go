package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"gitlab.com/hbang3/simple_bank/api"
	db "gitlab.com/hbang3/simple_bank/db/sqlc"
	"gitlab.com/hbang3/simple_bank/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
}
