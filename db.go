package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type DB struct {
	conn *sql.DB
}

const (
	Items string = "public.\"Items\""
)

func connString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("PGDBUSER"), os.Getenv("PGDBPASS"), os.Getenv("PGDBHOST"), os.Getenv("PGDBPORT"), os.Getenv("PGDB"))
}

func NewConn() DB {
	connString := connString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Could not connect to database:\nconnString: '%s'\nDatabase: '+%v'", connString, db)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}

	return DB{conn: db}
}

func (db DB) Insert(table string, cols []string, values []string) error {
	if len(cols) != len(values) {
		return errors.New("Number of columns and values do not match")
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, format(cols, "\""), format(values, "'"))
	_, err := db.conn.Exec(query)

	return err
}

func format(collection []string, surroundings string) string {
	var sb strings.Builder

	const SEP = ", "
	for _, c := range collection {
		_, _ = sb.WriteString(surroundings)
		_, _ = sb.WriteString(c)
		_, _ = sb.WriteString(surroundings)
		_, _ = sb.WriteString(SEP)
	}

	return sb.String()[:sb.Len()-len(SEP)]
}

// rows, err := conn.Query("select * from public.\"Items\"")
// if err != nil {
// 	log.Printf("Error querying: %s\n", err.Error())
// }
