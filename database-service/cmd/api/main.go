package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	db "database-svc/db/sqlc"
	"database-svc/pkg/server"
	"database/sql"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)


const webPort = "80"
var counts int64

func main() {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	store := db.NewStore(conn)
	
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: server.Router(store),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
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