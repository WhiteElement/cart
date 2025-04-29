package db

import (
	"cartv2/cart/item/item"
	"cartv2/cart/shoppinglist/shoppinglist"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

type DB struct {
	Conn *sql.DB
}

const (
	Items string = "public.\"Items\""
	Lists string = "public.\"Lists\""
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

	return DB{Conn: db}
}

// TODO: Muss umgeschrieben werden:
// statt selbst den String zu formatieren, die parametrisierte Version benutzen
// 	if len(cols) != len(values) {
// 		return errors.New("Number of columns and values do not match")
// 	}
//
// 	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, format(cols, "\""), format(values, "'"))
// 	_, err := db.conn.Exec(query)
//
// 	return err
// }

func (db DB) QueryAllItems() []item.Item {
	rows, err := db.Conn.Query(fmt.Sprintf("SELECT * FROM %s", Items))
	var items []item.Item
	if err != nil {
		return items
	}

	for rows.Next() {
		var item item.Item
		rows.Scan(&item.Id, &item.Name)
		items = append(items, item)
	}

	return items
}

func (db DB) QueryAllLists() []shoppinglist.List {
	rows, err := db.Conn.Query(fmt.Sprintf("SELECT * FROM %s", Lists))
	var lists []shoppinglist.List
	if err != nil {
		return lists
	}

	for rows.Next() {
		var list shoppinglist.List
		rows.Scan(&list.Id, &list.Name, &list.Created, &list.Updated)
		lists = append(lists, list)
	}

	return lists
}

func (db DB) QueryList(id int) shoppinglist.List {
	row := db.Conn.QueryRow(fmt.Sprintf("SELECT * FROM %s WHERE \"Lists\".\"Id\" = %d", Lists, id))

	var list shoppinglist.List
	row.Scan(&list.Id, &list.Name, &list.Created, &list.Updated)

	return list
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
