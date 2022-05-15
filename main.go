package main

import (
	"database/sql"
	"log"

	"github.com/Malarkey-Jhu/simple-bank/api"
	db "github.com/Malarkey-Jhu/simple-bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	addr     = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(addr)

	if err != nil {
		log.Fatal("can not start server:", err)
	}

}
