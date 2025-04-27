package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func connString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("PGDBUSER"), os.Getenv("PGDBPASS"), os.Getenv("PGDBHOST"), os.Getenv("PGDBPORT"), os.Getenv("PGDB"))
}

func NewConn() *sql.DB {
	connString := connString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Could not connect to database:\nconnString: '%s'\nDatabase: '+%v'", connString, db)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}
