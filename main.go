package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ruslanislan/simple_bank/api"
	db "github.com/ruslanislan/simple_bank/db/sqlc"
	"github.com/ruslanislan/simple_bank/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server,err := api.NewServer(config,store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}
	
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
