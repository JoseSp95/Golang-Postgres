package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetConnection() *sql.DB  {
	dsn := "postgres://postgres:root@localhost:5432/gocrud?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil{
		log.Fatal(err)
		return nil
	}
	return db
}


