package main

import (
	"log"
	"os"
	"time"

	"database/sql"
	db "ganso-core/db/sqlc"
	"ganso-core/pkg/server"

	_ "github.com/lib/pq"
)


const webUrl = "0.0.0.0:80"
var counts int64

func main() {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	store := db.NewStore(conn)
	server := server.NewServer(store)
	
	err := server.Start(webUrl)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB{
	dsn := os.Getenv("DSN")

	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready.")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return conn
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Timeout for 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}