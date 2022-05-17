package main

import (
	"database/sql"
	"log"

	"github.com/Malarkey-Jhu/simple-bank/api"
	db "github.com/Malarkey-Jhu/simple-bank/db/sqlc"
	"github.com/Malarkey-Jhu/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddr)

	if err != nil {
		log.Fatal("can not start server:", err)
	}

}
